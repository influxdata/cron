package cron

import (
	"reflect"
	"testing"
	"time"
)

// func Test_timeBM(t *testing.T) {
// 	tests := struct {
// 		name string
// 		year time.Month
// 		month time.Month
// 		dom  int
// 		hour int
// 		min  int
// 		sec  int
// 		nsec int
// 		loc  *time.Location
// 	}{
// 		{name: "basic", year: 2019, month: time.February, dom: 19, hour: 17, min: 37, sec: 29, nsec: 45678, loc: time.UTC},
// 	}
// 	for _, x := range tests {
// 		nt:=timebm(time.Date(x.year, x.month, x.dom, x.hour, x.min, x.sec, x.nsec, x.loc))
// 		switch {
// 		case nt.month && x.month != x.month:
// 			t.Fatal("month incorrect")
// 		// TODO(docmerlin): check year
// 		case nt.dom && x.dom != x.dom:

// 		}
// 	}
// }

func mustParseTime(s string) time.Time {
	ts, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		panic(err)
	}
	return ts

}

func Test_nextTime_next(t *testing.T) {
	ts, err := time.Parse(time.RFC3339Nano, "2019-06-27T06:05:12.622652Z")
	if err != nil {
		t.Fatal(err)
	}
	parsit := func(s string, zone *time.Location) nextTime {
		nt, err := parse(s, zone)
		if err != nil {
			t.Fatal(err)
		}
		return nt
	}
	tests := []struct {
		name    string
		nt      nextTime
		from    time.Time
		want    time.Time
		wanterr bool
	}{
		// TODO: Add test cases.
		{
			name: "all stars",
			nt:   parsit("* * * * * * *", time.UTC),
			from: ts,
			want: ceilTime(ts, time.Second),
		},
		{
			name: "forth second",
			nt:   parsit("4 * * * * * *", time.UTC),
			from: ts,
			want: ceilTime(ts, time.Minute).Add(4 * time.Second),
		},
		{
			name: "forth second, 7th minute",
			nt:   parsit("4 7 * * * * *", time.UTC),
			from: ts,
			want: mustParseTime("2019-06-27T06:07:04Z"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := tt.nt.next(tt.from); !reflect.DeepEqual(got, tt.want) || tt.wanterr != (err != nil) {
				t.Errorf("from = %v, nextTime.next() = %v, want %v", tt.from, got, tt.want)
			}
		})
	}
}

// func Test_next(t *testing.T) {
// 	type args struct {
// 		nt   nextTime
// 		from time.Time
// 		tz   *time.Location
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    time.Time
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := next(tt.args.nt, tt.args.from, tt.args.tz)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("next() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("next() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
