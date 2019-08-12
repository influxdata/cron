package cron

//go:generate ragel -G2 -Z parse.rl

import (
	"errors"
	"math/bits"
	"time"
)

// ParseUTC parses a cron string in UTC timezone.
func ParseUTC(s string) (Parsed, error) {
	return parse(s, time.UTC)
}

type years struct {
	low  uint64
	high uint64
	end  uint64 // this is here so we can support 2098 and 2099
}

func (ys years) isZero() bool {
	return ys.low|ys.high|ys.end == 0
}

func (ys *years) set(year int) {
	y := uint64(year - 1970)
	switch {
	case y > 128:
		ys.end |= (1 << (y - 128))
	case y > 64:
		ys.high |= (1 << (y - 64))
	default:
		ys.low |= (1 << y)
	}
}

func (ys years) TrailingZeros64() int {
	zeros := bits.TrailingZeros64(ys.low)
	if zeros == 64 {
		zeros += bits.TrailingZeros64(ys.high)
		if zeros == 128 {
			zeros += bits.TrailingZeros64(ys.end)
		}
	}
	return zeros
}

func (ys years) Intersection(y2 years) years {
	return years{
		low:  ys.low & y2.low,
		high: ys.high & y2.high,
		end:  ys.end & y2.end,
	}
}

// this is undefined if the year is above 2070 or below 1970
func (ys years) in(year int) bool {
	y := uint64(year - 1970)
	if y >= 128 {
		return ys.end&(1<<(y-128)) > 0
	}
	if y >= 64 {
		return ys.high&(1<<(y-64)) > 0
	}
	return ys.low&(1<<y) > 0
}

func isLeap(y int) bool {
	return (y&3 == 0 && y%100 != 0) || y%400 == 0
}

// Parsed is a parsed cron string.  It can be used to find the next instance of a cron task.
type Parsed struct {
	second uint64
	minute uint64
	hour   uint64
	dom    uint64
	month  uint64
	dow    uint64
	year   years
	loc    *time.Location
}

var maxMonthLengths = [13]uint64{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31, 31} // we wrap back around from jan to jan to make stuff easy

func (nt *Parsed) nextYear(y, m, d uint64) int {
	if y < 1969 {
		y = 1969
	}
	y++
	yBits := y - 1970
	var addToY uint64
	if yBits >= 128 {
		y += uint64(bits.TrailingZeros64(nt.year.end >> (yBits - 128)))
	} else if yBits >= 64 {
		addToY += uint64(bits.TrailingZeros64(nt.year.high >> (yBits - 64)))
		if addToY == 64 {
			addToY += uint64(bits.TrailingZeros64(nt.year.high))
		}
	} else {
		addToY = uint64(bits.TrailingZeros64(nt.year.low >> yBits))
		if addToY == 64 {
			addToY -= yBits
			addToY += uint64(bits.TrailingZeros64(nt.year.high))
		}
		if addToY == 128 {
			addToY += uint64(bits.TrailingZeros64(nt.year.end))
		}
	}
	y += addToY

	if y > 2099 {
		return -1
	}
	//Feb 29th special casing
	// if we only allow feb 29th, or  then the next year must be a leap year
	// remember we zero index months here
	if m == 1 && d == 28 {
		if isLeap(int(y)) {
			return int(y)
		}
		if y%100 == 0 {
			return int(y + 4)
		}
		if y%3 != 0 { // is multiple of 4?
			y = ((y + 4) >> 2) << 2 // next multiple of 4
			if y%100 != 0 {
				return int(y)
			}
			return int(y + 4)
		}
	}

	return int(y)
}

// returns the index for the month
func (nt *Parsed) nextMonth(m, d uint64) int {
	m = uint64(bits.TrailingZeros64(nt.month>>m)) + m
	d = uint64(bits.TrailingZeros64(nt.dom>>d)) + d
	if m > 11 { // if there is no next avaliable months
		return -1
	}
	if maxMonthLengths[m] <= d {
		m++
	}
	return int(m)
}

func (nt *Parsed) nextDay(y int, m int, d uint64) int {
	firstOfMonth := time.Date(y, time.Month(m+1), 1, 0, 0, 0, 0, nt.loc)
	days := nt.prepDays(firstOfMonth.Weekday(), m, y)
	d++
	d = uint64(bits.TrailingZeros64(days>>d)) + d
	if m >= 12 {
		return -1
	}
	if d >= uint64(maxMonthLengths[m]) {
		return -1
	}
	return int(d)
}

func (nt *Parsed) nextHour(h uint64) int {
	h++
	h = uint64(bits.TrailingZeros64(nt.hour>>h)) + h
	if h >= 24 {
		return -1
	}
	return int(h)
}

func (nt *Parsed) nextMinute(m uint64) int {
	m++
	m = uint64(bits.TrailingZeros64(nt.minute>>m)) + m
	if m >= 60 {
		return -1
	}
	return int(m)
}

func (nt *Parsed) nextSecond(s uint64) int {
	s++
	s = uint64(bits.TrailingZeros64(nt.second>>s)) + s
	if s >= 60 {
		return -1
	}
	return int(s)
}

// undefined for shifts larger than 7
// firstDayOfWeek is the 0 indexed day of first day of the month
func (nt *Parsed) prepDays(firstDayOfWeek time.Weekday, month int, year int) uint64 {
	doms := uint64(1<<maxMonthLengths[month]) - 1
	if month == 1 && isLeap(year) { // 0 indexed feb
		doms = doms | (1 << 28)
	}
	return (nt.dow >> uint64(firstDayOfWeek)) & (doms & nt.dom)
}

// Next returns the next time a cron task should run given a Parsed cron string.
// It will error if the Parsed is not from a zero value.
func (nt Parsed) Next(from time.Time) (time.Time, error) {
	y, MTime, dTime := from.Date()
	d := int(dTime - 1) //time's day is 1 indexed but our day is zero indexed
	M := int(MTime - 1) //time's month is 1 indexed in time but ours is 0 indexed
	h := from.Hour()
	m := from.Minute()
	s := from.Second()

	updateHour := nt.hour&(1<<uint(h)) == 0
	updateMin := nt.minute&(1<<uint64(m)) == 0
	updateSec := nt.second&(1<<uint64(s)) == 0 || !(updateMin && updateHour)

	if updateSec {
		if s2 := nt.nextSecond(uint64(s)); s2 < 0 {
			s = bits.TrailingZeros64(nt.second) // if not found set to first second and advance minute
			updateMin = true
		} else {
			s = s2
		}
	}
	if updateMin {
		if m2 := nt.nextMinute(uint64(m)); m2 < 0 {
			m = bits.TrailingZeros64(nt.minute) // if not found set to first second and advance minute
			updateHour = true
		} else {
			m = m2
		}
		s = bits.TrailingZeros64(nt.second) // if not found set to first second and advance minute
	}
	if updateHour {
		if h2 := nt.nextHour(uint64(h)); h2 < 0 {
			h = bits.TrailingZeros64(nt.hour) // if not found set to first hour and advance the day
			d++
		} else {
			h = h2
		}
		m = bits.TrailingZeros64(nt.minute)
		s = bits.TrailingZeros64(nt.second)
	}
	weekDayOfFirst := time.Date(y, time.Month(M+1), 1, 0, 0, 0, 0, from.Location()).Weekday()

	if (d == 28 && M == 1 && !isLeap(y)) || !nt.year.in(y) || (nt.month&(1<<uint(M)) == 0) || nt.prepDays(weekDayOfFirst, M, y)&(1<<uint(d)) == 0 {
		h = bits.TrailingZeros64(nt.hour) // if not found set to first hour and advance the day
		m = bits.TrailingZeros64(nt.minute)
		s = bits.TrailingZeros64(nt.second)
	}
	oldYear, oldMonth, oldDay := y, M, d

	for y <= 2099 {
		if nt.prepDays(weekDayOfFirst, M, y)&(1<<uint(d)) == 0 {
			if d2 := nt.nextDay(y, M, uint64(d)); d2 < 0 {
				M++
				weekDayOfFirst = time.Date(y, time.Month(M+1), 1, 0, 0, 0, 0, from.Location()).Weekday()
				d = bits.TrailingZeros64(nt.prepDays(weekDayOfFirst, M, y))
			} else {
				d = d2
			}
		}

		if nt.month&(1<<uint(M)) == 0 {
			if M2 := nt.nextMonth(uint64(M), uint64(d)); M2 < 0 {
				M = bits.TrailingZeros64(nt.month)
				y++
				weekDayOfFirst = time.Date(y, time.Month(M+1), 1, 0, 0, 0, 0, from.Location()).Weekday()
				d = bits.TrailingZeros64(nt.prepDays(weekDayOfFirst, M, y))
			} else {
				M = M2
			}
		}
		if !nt.year.in(y) || (d == 28 && M == 1 && !isLeap(y)) {
			y2 := nt.nextYear(uint64(y), uint64(M), uint64(d))
			if y2 < 0 {
				return time.Time{}, errors.New("could not fulfil schedule due to year")
			}

			y = y2
			M = bits.TrailingZeros64(nt.month)
			weekDayOfFirst = time.Date(y, time.Month(M+1), 1, 0, 0, 0, 0, from.Location()).Weekday()
			d = bits.TrailingZeros64(nt.prepDays(weekDayOfFirst, M, y))
		}
		// if no changes, return
		if oldYear == y && oldMonth == M && oldDay == d {
			if nt.prepDays(weekDayOfFirst, M, y)&(1<<uint(d)) != 0 && (nt.year.in(y) && (d != 28 || M != 1 || isLeap(y))) && nt.month&(1<<uint(M)) != 0 {
				return time.Date(y, time.Month(M+1), int(d+1), h, m, s, 0, from.Location()), nil // again we 0 index day of month, and the month but the actual date doesn't
			}
		}
		oldYear, oldMonth, oldDay = y, M, d
	}
	return time.Time{}, errors.New("could not fulfil schedule before 2100")
}
