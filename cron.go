package cron

//go:generate ragel -G2 -Z parse.rl
//go:generate ragel -G2 -Z parse.rl -Smain_start -o parse.dot
//go:generate dot -Tpdf parse.dot -o parse.pdf

import (
	"errors"
	"math/bits"
	"time"
)

type years struct {
	low  uint64
	high uint64
}

func (y years) isZero() bool {
	return y.low|y.high == 0
}

func (ys *years) set(year int) {
	y := uint64(year - 1970)
	if y > 64 {
		ys.high |= (1 << (y - 64))
	}
	ys.low |= (1 << y)
}

func (y years) TrailingZeros64() int {
	zeros := bits.TrailingZeros64(y.low)
	if zeros == 64 {
		zeros += bits.TrailingZeros64(y.high)
	}
	return zeros
}

func (y years) Intersection(y2 years) years {
	return years{
		low:  y.low & y2.low,
		high: y.high & y2.high,
	}
}

// this is undefined if the year is above 2070 or below 1970
func (ys years) in(year int) bool {
	y := uint64(year - 1970)
	if y >= 64 {
		return ys.high&(1<<(y-64)) > 0
	}
	return ys.low&(1<<y) > 0
}

func isLeap(y int) bool {
	return (y&3 == 0 && y%100 != 0) || y%400 == 0
}

type nextTime struct {
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

// func monthLen(year, month int) uint64 {
// 	l := maxMonthLengths[month-1]
// 	// check if feb and not leap year
// 	if month == 1 && (year&3 != 0 || year%100 == 0) {
// 		l--
// 	}
// 	return l
// }

func (nt *nextTime) nextYear(y, m, d uint64) int {
	yBits := uint64(0)
	y++
	if y >= 1970 {
		yBits = y - 1970
	}

	if yBits >= 64 {
		yBits -= 64
		y += uint64(bits.TrailingZeros64(nt.year.high >> yBits))
	} else {
		addToY := uint64(bits.TrailingZeros64(nt.year.low >> yBits))
		if addToY == 64 {
			addToY = uint64(bits.TrailingZeros64(nt.year.high))
		}
		y += addToY
	}

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
func (nt *nextTime) nextMonth(m, d uint64) int {
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

func (nt *nextTime) nextDay(y int, m time.Month, d uint64) int {
	firstOfMonth := time.Date(y, m+1, 1, 0, 0, 0, 0, nt.loc)
	days := nt.prepDays(firstOfMonth.Weekday())
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

func (nt *nextTime) nextHour(h uint64) int {
	h++
	h = uint64(bits.TrailingZeros64(nt.hour>>h)) + h
	if h >= 24 {
		return -1
	}
	return int(h)
}

func (nt *nextTime) nextMinute(m uint64) int {
	m++
	m = uint64(bits.TrailingZeros64(nt.minute>>m)) + m
	if m >= 60 {
		return -1
	}
	return int(m)
}

func (nt *nextTime) nextSecond(s uint64) int {
	s++
	s = uint64(bits.TrailingZeros64(nt.second>>s)) + s
	if s >= 60 {
		return -1
	}
	return int(s)
}

// undefined for shifts larger than 7
// firstDayOfWeek is the 0 indexed day of first day of the month
func (nt *nextTime) prepDays(firstDayOfWeek time.Weekday) uint64 {
	return (nt.dow >> uint64(firstDayOfWeek)) & nt.dom
}

func (nt *nextTime) next(from time.Time) (time.Time, error) {
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
		if m2 := nt.nextMinute(uint64(m)); m < 0 {
			m = bits.TrailingZeros64(nt.minute) // if not found set to first second and advance minute
			updateHour = true
		} else {
			m = m2
		}
		s = bits.TrailingZeros64(nt.second) // if not found set to first second and advance minute
	}
	if updateHour {
		if h2 := nt.nextHour(uint64(h)); h < 0 {
			h = bits.TrailingZeros64(nt.hour) // if not found set to first hour and advance the day
			d++
		} else {
			h = h2
		}
		m = bits.TrailingZeros64(nt.minute)
		s = bits.TrailingZeros64(nt.second)
	}
	weekDayOfFirst := time.Date(y, time.Month(M+1), 1, 0, 0, 0, 0, from.Location()).Weekday()

	if (d == 28 && M == 1 && !isLeap(y)) || !nt.year.in(y) || (nt.month&(1<<uint(M)) == 0) || nt.prepDays(weekDayOfFirst)&(1<<uint(d)) == 0 {
		h = bits.TrailingZeros64(nt.hour) // if not found set to first hour and advance the day
		m = bits.TrailingZeros64(nt.minute)
		s = bits.TrailingZeros64(nt.second)
	}
	oldYear, oldMonth, oldDay := y, M, d

	for y < 2099 {
		if nt.prepDays(weekDayOfFirst)&(1<<uint(d)) == 0 {
			if d2 := nt.nextDay(y, time.Month(M), uint64(d)); d2 < 0 {
				M++
				weekDayOfFirst = time.Date(y, time.Month(M+1), 1, 0, 0, 0, 0, from.Location()).Weekday()
				d = bits.TrailingZeros64(nt.prepDays(weekDayOfFirst))
			} else {
				d = d2
			}
		}

		if nt.month&(1<<uint(M)) == 0 {
			if M2 := nt.nextMonth(uint64(M), uint64(d)); M2 < 0 {
				M = bits.TrailingZeros64(nt.month)
				y++
				weekDayOfFirst = time.Date(y, time.Month(M+1), 1, 0, 0, 0, 0, from.Location()).Weekday()
				d = bits.TrailingZeros64(nt.prepDays(weekDayOfFirst))
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
			d = bits.TrailingZeros64(nt.prepDays(weekDayOfFirst))
		}
		// if no changes, return
		if oldYear == y && oldMonth == M && oldDay == d {
			if nt.prepDays(weekDayOfFirst)&(1<<uint(d)) != 0 && (nt.year.in(y) && (d != 28 || M != 1 || isLeap(y))) && nt.month&(1<<uint(M)) != 0 {
				return time.Date(y, time.Month(M+1), int(d+1), h, m, s, 0, from.Location()), nil // again we 0 index day of month, and the month but the actual date doesn't
			}
		}
		oldYear, oldMonth, oldDay = y, M, d
	}
	return time.Time{}, errors.New("could not fulfil schedule before 2099")
}

func (nt *nextTime) intersection(ct *nextTime) *nextTime {
	return &nextTime{
		second: nt.second & ct.second,
		minute: nt.minute & ct.minute,
		hour:   nt.hour & ct.hour,
		dow:    nt.dow & ct.dow,
		month:  nt.month & ct.month,
		dom:    nt.dom & ct.dom,
		year: years{
			high: nt.year.high & ct.year.high,
			low:  nt.year.low & ct.year.low,
		},
	}
}

func (nt *nextTime) updateDateIntersectDate(y int, M, d uint64) {
	nt.dom = nt.dom & 1 << d
	nt.year.high = 0
	nt.year.low = 0
	nt.year.set(y)
	nt.month = nt.month & 1 << M
}

// func timebm(ts time.Time) *nextTime {
// 	_, M, D := ts.Date()
// 	hr, m, s := ts.Clock()
// 	dow := ts.Weekday()
// 	return &nextTime{
// 		second: 1 << uint(s),
// 		minute: 1 << uint(m),
// 		hour:   1 << uint(hr),
// 		dow:    1 << uint(dow),
// 		month:  1 << uint(M-1),
// 		dom:    1 << uint(D-1),
// 	}
// }
