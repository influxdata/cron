package cron

//go:generate ragel -Z -G1 parse.rl

import (
	"errors"
	"fmt"
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

func (ys years) set(year int) {
	y := uint64(year - 1970)
	if y > 64 {
		ys.high |= (1 << (y - 64))
	}
	ys.low |= (1 << y)
}

func (ys years) unset(year int) {
	y := uint64(year - 1970)
	if y > 64 {
		ys.high &= ^(1 << (y - 64))
	}
	ys.low &= ^(1 << y)
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
	if y > 64 {
		return ys.high&(1<<(y-64)) > 0
	}
	return ys.low&(1<<y) > 0
}

type nextTime struct {
	second uint64
	minute uint64
	hour   uint64
	dow    uint64
	month  uint64
	dom    uint64
	year   years
	loc    *time.Location
}

const sundaysAtFirst = uint64(1 | 1<<7 | 1<<14 | 1<<21 | 1<<28)

var maxMonthLengths = [12]int{1 << 31, 1 << 28, 1 << 31, 1 << 30, 1 << 31, 1 << 30, 1 << 31, 1 << 31, 1 << 30, 1 << 31, 1 << 30, 1 << 31}

func monthLen(year, month int) int {
	l := maxMonthLengths[month-1]
	// check if feb and not leap year
	if month == 2 && (year&3 != 0 || year%100 == 0) {
		l--
	}
	return l
}

func (nt *nextTime) nextYear(y, m, d uint64) int {
	y++
	//Feb 29th special casing
	// if we only allow feb 29th, or  then the next year must be a leap year
	// remember we zero index months here
	if m == 1 && bits.TrailingZeros64(nt.dom) > 28 {
		if y&3 == 0 && y%100 != 0 {
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
	return -1
}

func (nt *nextTime) nextMonth(m, d uint64) time.Month {
	// no need to increment month, because we zero index and the result of Date is 1 indexed
	m = uint64(bits.TrailingZeros64(nt.month>>m)) + m
	d = uint64(bits.TrailingZeros64(nt.dom>>d)) + d
	if nt.month < 1<<m { // if there is no next avaliable months
		return -1
	}

	if maxMonthLengths[m-1] < (1<<d) && m < 12 {
		m++
	}
	return time.Month(m)
}

func (nt *nextTime) nextDay(y int, m time.Month, d uint64) int {
	firstOfMonth := time.Date(y, m, 1, 0, 0, 0, 0, nt.loc)
	days := ((sundaysAtFirst << uint64(firstOfMonth.Weekday())) * nt.dow) & nt.dom
	d = uint64(bits.TrailingZeros64(days>>d)) + d
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

func (nt *nextTime) next(from time.Time) (time.Time, error) {
	y, M, d := from.Date()
	d-- //day is 1 indexed but our day is zero indexed
	M-- //month is 1 indexed in time but ours is 0 indexed
	h := from.Hour()
	m := from.Minute()
	s := from.Second()

	updateMin := nt.minute&(1<<uint64(m)) == 0
	updateHour := nt.hour&(1<<uint(h)) == 0
	updateMonth := nt.month&(1<<uint(M)) == 0
	updateYear := nt.year.in(y)

	firstOfMonth := time.Date(y, M+1, 1, 0, 0, 0, 0, nt.loc)
	days := ((sundaysAtFirst << uint64(firstOfMonth.Weekday())) * nt.dow) & nt.dom

	updateDay := days&(1<<uint(d)) == 0

	updateSec := nt.second&(1<<uint64(s)) == 0 || !(updateMin && updateHour && updateMonth && updateYear)

	if updateSec {
		fmt.Println("UPDATING SEC")
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
	}
	if updateHour {
		if h2 := nt.nextHour(uint64(h)); h < 0 {
			h = bits.TrailingZeros64(nt.hour) // if not found set to first second and advance minute
			updateDay = true
		} else {
			h = h2
		}
	}
processDay:
	if updateDay {
		if d2 := nt.nextDay(y, time.Month(M), uint64(d)); d2 < 0 {
			d = 1
			updateDay = true
			updateMonth = true
		} else {
			d = d2
			updateDay = false
		}
	}
processMonth:
	if y > 2099 {
		return time.Time{}, errors.New("could not fulfil schedule")
	}
	if updateMonth {
		if M2 := nt.nextMonth(uint64(M), uint64(d)); M2 < 0 {
			M = time.Month(bits.TrailingZeros64(nt.month))
			updateYear = true
			goto processYear
		} else {
			M = M2
		}
		d = 0
		goto processDay
	}
processYear:
	if updateYear {
		if y2 := nt.nextYear(uint64(y), uint64(M), uint64(d)); y2 < 1 {
			return time.Time{}, errors.New("could not fulfil schedule")
		}
		updateMonth = true
		m = 0
		d = 0
		goto processMonth
	}
	return time.Date(y, M, d, h, m, s, 0, from.Location()), nil
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

func timebm(ts time.Time) *nextTime {
	_, M, D := ts.Date()
	hr, m, s := ts.Clock()
	dow := ts.Weekday()
	return &nextTime{
		second: 1 << uint(s),
		minute: 1 << uint(m),
		hour:   1 << uint(hr),
		dow:    1 << uint(dow),
		month:  1 << uint(M-1),
		dom:    1 << uint(D-1),
	}
}
