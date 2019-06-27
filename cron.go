package cron

//go:generate ragel -Z -G1 parse.rl

import (
	"math/bits"
	"time"
)

type nextTime struct {
	second uint64
	minute uint64
	hour   uint64
	dow    uint64
	month  uint64
	dom    uint64
}

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
	if nt.month == 2 && bits.TrailingZeros64(nt.dom) >= 28 {
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
	if nt.month < 1<<m {
		return -1
	}

	if maxMonthLengths[m-1] < (1<<d) && m < 12 {
		m++
	}

	return time.Month(m)
}

func (nt *nextTime) nextDayOfMonth(m, d uint64) int {
	// no need to increment day, because we zero index and the result of Date is 1 indexed
	// find next avaliable dom
	d = uint64(bits.TrailingZeros64(nt.dom>>d)) + d
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

func (nt *nextTime) next(from time.Time) time.Time {
	y, M, d := from.Date()
	h := from.Hour()
	m := from.Minute()
	s := from.Second()

	if s2 := nt.nextSecond(uint64(s)); s2 < 0 {
		s = bits.TrailingZeros64(nt.second)
		if m2 := nt.nextMinute(uint64(m)); m2 < 0 {
			m = bits.TrailingZeros64(nt.minute)
			if h2 := nt.nextHour(uint64(h)); h2 < 0 {
				h = bits.TrailingZeros64(nt.hour)
				if d2 := nt.nextDayOfMonth(uint64(y), uint64(M)); d2 < 0 {
					d = bits.TrailingZeros64(nt.month)
					if M2 := nt.nextMonth(uint64(M), uint64(d)); M2 < 0 {
						M = 0
						y2 := nt.nextYear(uint64(y), uint64(M), uint64(d))
						if y2 < 0 {
							return time.Time{}
						}
						y = y2
					} else {
						M = M2
					}
				} else {
					d = d2
				}
			} else {
				h = h2
			}
		} else {
			m = m2
		}
	} else {
		s = s2
	}
	return time.Date(y, M, d, h, m, s, 0, from.Location())
}

func (nt *nextTime) intersectionExists(ct *nextTime) bool {
	return ct.month&nt.month != 0 &&
		ct.dom&nt.dom != 0 &&
		ct.hour&nt.hour != 0 &&
		ct.minute&nt.minute != 0 &&
		ct.second&nt.second != 0
}

func timebm(ts time.Time) nextTime {
	_, M, D := ts.Date()
	hr, m, s := ts.Clock()
	dow := ts.Weekday()
	return nextTime{
		second: 1 << uint(s),
		minute: 1 << uint(m),
		hour:   1 << uint(hr),
		dow:    1 << uint(dow),
		month:  1 << uint(M-1),
		dom:    1 << uint(D-1),
	}
}
