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
	parsit := func(s string, zone *time.Location) func(t *testing.T) (nextTime, error) {
		return func(t *testing.T) (nextTime, error) {
			return parse(s, zone)
		}
	}
	tests := []struct {
		name         string
		nt           func(t *testing.T) (nextTime, error)
		from         time.Time
		want         time.Time
		wanterr      bool
		wantParseErr bool
	}{
		{
			name: "2023",
			nt:   parsit(" *   * * * * * 2023 ", time.UTC),
			from: ts,
			want: mustParseTime("2023-01-01T00:00:00Z"),
		},
		{
			name: "all stars",
			nt:   parsit("* * * * * 4-6 *", time.UTC),
			from: ts,
			want: ceilTime(ts, time.Second),
		},
		{
			name: "all seconds starting from ten by 7",
			nt:   parsit("10/4 * * * * * *", time.UTC),
			from: ts,
			want: mustParseTime("2019-06-27T06:05:14Z"),
		},
		{
			name: "all seconds between 3 and 12 by 3",
			nt:   parsit("3-12/2 * * * * * *", time.UTC),
			from: ts,
			want: mustParseTime("2019-06-27T06:06:03Z"),
		},
		{
			name: "all minutes between 4 and 22 by 3",
			nt:   parsit("* 3-22/3 * * * * *", time.UTC),
			from: ts,
			want: mustParseTime("2019-06-27T06:06:00Z"),
		},
		{
			name: "at every 5nd minute starting at minute 3.",
			nt:   parsit("* 3/5 * * * * *", time.UTC),
			from: ts,
			want: mustParseTime("2019-06-27T06:08:00Z"),
		},
		{
			name: "at every 2nd minute starting at minute 3 for even seconds.",
			nt:   parsit("*/2 3/2 * * * * *", time.UTC),
			from: ts,
			want: mustParseTime("2019-06-27T06:05:14Z"),
		},
		{
			name: "once a minute at second 1",
			nt:   parsit("1 * * * * * *", time.UTC),
			from: ts,
			want: mustParseTime("2019-06-27T06:06:01Z"),
		},
		{
			name: "all stars extra space at end",
			nt:   parsit("* * * * * * * ", time.UTC),
			from: ts,
			want: ceilTime(ts, time.Second),
		},
		{
			name: "all stars extra space at beginning",
			nt:   parsit(" * * * * * * *", time.UTC),
			from: ts,
			want: ceilTime(ts, time.Second),
		},
		{
			name: "all stars extra spaces randomly ",
			nt:   parsit(" * * *  * \t*\t* * ", time.UTC),
			from: ts,
			want: ceilTime(ts, time.Second),
		},
		{
			name: "fourth second",
			nt:   parsit("4 * * * * * *", time.UTC),
			from: ts,
			want: ceilTime(ts, time.Minute).Add(4 * time.Second),
		},
		{
			name: "every fourth second",
			nt:   parsit("*/4 * * * * * *", time.UTC),
			from: ts,
			want: mustParseTime("2019-06-27T06:05:16Z"),
		},
		{
			name: "every fourth second and every 2nd hour",
			nt:   parsit("*/4 * */2 * * * *", time.UTC),
			from: ts,
			want: mustParseTime("2019-06-27T06:05:16Z"),
		},
		{
			name: "every fourth second and every 4th hour between 4 and 16",
			nt:   parsit("*/4 * 4-16/4 * * * *", time.UTC),
			from: ts,
			want: mustParseTime("2019-06-27T08:00:00Z"),
		},
		{
			name: "every second except on thursday",
			nt:   parsit("* * * * * FRI,SAT,SUN,MON,TUE,WED *", time.UTC),
			from: ts,
			want: mustParseTime("2019-06-28T00:00:00Z"),
		},
		{
			name: "every second on saturday",
			nt:   parsit("* * * * * SAT *", time.UTC),
			from: ts,
			want: mustParseTime("2019-06-29T00:00:00Z"),
		},
		{
			name: "fourth second, 7th minute",
			nt:   parsit("4 7 * * * * *", time.UTC),
			from: ts,
			want: mustParseTime("2019-06-27T06:07:04Z"),
		},
		{
			name: "fourth second, 7th minute on saturday",
			nt:   parsit("4 7 * * * 6 *", time.UTC),
			from: ts,
			want: mustParseTime("2019-06-29T00:07:04Z"),
		},
		{
			name: "fourth second, 5-10th minute on saturday",
			nt:   parsit("4 5-10 * * * 6 *", time.UTC),
			from: ts,
			want: mustParseTime("2019-06-29T00:05:04Z"),
		},
		{
			name: "noon on dec 3rd",
			nt:   parsit("* * 12 3 12 * *", time.UTC),
			from: ts,
			want: mustParseTime("2019-12-03T12:00:00Z"),
		},
		{
			name: "noon on leap day",
			nt:   parsit("* * 12 29 FEB * *", time.UTC),
			from: ts,
			want: mustParseTime("2020-02-29T12:00:00Z"),
		},
		{
			name: "noon on leap day on day on wed",
			nt:   parsit("* * 12 29 FEB WED *", time.UTC),
			from: ts,
			want: mustParseTime("2040-02-29T12:00:00Z"),
		},
		{
			name: "noon on leap day on day on tuesday,wednesday, or thursday",
			nt:   parsit("* * 12 29 FEB tue-thu *", time.UTC),
			from: ts,
			want: mustParseTime("2024-02-29T12:00:00Z"),
		},
		{
			name: "noon on leap day on day on Monday or sunday",
			nt:   parsit("* * 12 29 2 MON,0 *", time.UTC),
			from: ts,
			want: mustParseTime("2032-02-29T12:00:00Z"),
		},
		{
			name: "when july 4 is a weekday",
			nt:   parsit("* * * 4 7 MON,TUE,WED,4,5 *", time.UTC),
			from: ts,
			want: mustParseTime("2019-07-04T00:00:00Z"),
		},
		{
			name:    "when leap day is a monday and the year is divisible by 3",
			nt:      parsit("* * * 29 2 MON */3", time.UTC),
			from:    ts,
			want:    time.Time{},
			wanterr: true,
		},
		{
			name: "jan 1 falls on a monday and year is divisible by 3",
			nt:   parsit("* * * 1 1 MON */3", time.UTC),
			from: ts,
			want: mustParseTime("2046-01-01T00:00:00Z"),
		},
		{
			name: "jan 1 on year is divisible by 2",
			nt:   parsit("* * * * * * */2", time.UTC),
			from: ts,
			want: mustParseTime("2020-01-01T00:00:00Z"),
		},
		{
			name: "6 position (no year)",
			nt:   parsit("* * * * * *", time.UTC),
			from: ts,
			want: mustParseTime("2019-06-27T06:05:13Z"),
		},
		{
			name: "5 position (no year, no sec) extra spaces",
			nt:   parsit("  * *  * * *  ", time.UTC),
			from: ts,
			want: mustParseTime("2019-06-27T06:06:00Z"),
		},
		{
			name: "1970",
			nt:   parsit("*  * *  * * *  1970", time.UTC),
			from: time.Time{},
			want: mustParseTime("1970-01-01T00:00:00Z"),
		},
		{
			name: "2097",
			nt:   parsit("*  * *  * * *  2097", time.UTC),
			from: time.Time{},
			want: mustParseTime("2097-01-01T00:00:00Z"),
		},
		{
			name: "yearly",
			nt:   parsit("@yearly", time.UTC),
			from: ts,
			want: mustParseTime("2020-01-01T00:00:00Z"),
		},
		{
			name: "annually",
			nt:   parsit("@annually", time.UTC),
			from: ts,
			want: mustParseTime("2020-01-01T00:00:00Z"),
		},
		{
			name:         "fred likes snickerdoodles",
			nt:           parsit("fred likes snickerdoodles", time.UTC),
			from:         ts,
			wantParseErr: true,
		},
		{
			name:         "@notanemailaddress",
			nt:           parsit("@notanemailaddress", time.UTC),
			from:         ts,
			wantParseErr: true,
		},
		{
			name:         "@yearlyplusextraletters",
			nt:           parsit("@yearlyplusextraletters", time.UTC),
			from:         ts,
			wantParseErr: true,
		},
		{
			name:         "too many fields",
			nt:           parsit("* * 8 8 1 1 0", time.UTC),
			from:         ts,
			wantParseErr: true,
		},
		{
			name:         "too many fields",
			nt:           parsit("* * 8 8 1 1 0", time.UTC),
			from:         ts,
			wantParseErr: true,
		},
		{
			name:         "error instead of summoning the old ones",
			nt:           parsit("c̴̡͈̪̪̬͙͈̼̞̥̩̜̪̳̜̰̬̱͇̮̼̳̟̩̳̫̾̃͐̽̒̎̈́̽͠͝͝͠t̴͍̫͓͉̥̤̺͇͇̦̮̞͈̬̮͕͖̼̺̭̥͓̬̣͉̻̭̣̓̐̿̇͊̈́̉̏͐̆͑͂͛̎͒͌̈́͋̃̾̔̕͝h̸̜̭͕͎̰̮̭͙̙͇͎̺̤̙͖̒͌̂̌̍̿̒̀̊̿̾̚̕̕͝͠ͅū̷̲̝̂̑̈́͊͋̓̆̈́́́̄̎̄̃͐̈́̇̚͝ͅl̷̢̡̛͉̠̞̮̝͙̻̜͓̹̻̩̯̙̹̯͓̝͇̫̫̠̿̑̇̓̀̂̀͗̒͋̀͌͒̕̕͜͠ͅh̵̡̛̖͉̟̞̣̖̬͓̜̄̀̂̽̄͒̅̀͋̓̀͌̆̂̀͌͌̉̀̐̈́͋̓͂̚̕͠͠ų̶̛͎̖͖̜̹̬͙̖̜̲̺͐̀̎̎̽͛͌͆̂̌͒̋̆̿̎̎̽̑̏̏̈̾͘̕̕̕̕̕̚͠ͅ ̸̢̨̡̛̙̺̹͖͙͔̩̮͙̝̳̘̰̣̪̦̥̰̠̝̼͕͔͌̑̃́̐̈͗̍̅̽̌̈́̑̉̋̃͘͜͝͠͠f̶̹̹͉̉ä̸̲͈̺͉̮̹̘̦̦̟̘̼̪͚̯͓̻͉̣̬͓̰̟͇͕̳̟̗̉̈̔̉̿̔̇̑̕͜g̴̢̧̡̱̞̭͉̼̤̟̗͚͖̩̖̤͉͉̤̪͖̣̠̜̟̟̜̜̺̯̭̋̐̌̆͒͌̓̒͛́̌͒̀̓̆͋͐̈́̐̚͝ň̴̨̢̨̡͓̥̪̬͍̙̳̫̰̲̗̰̥̖̫̹̱̟̰̒̿̓̀̒̊́͛̊̈́̒̌̂͐́̕͝", time.UTC),
			from:         ts,
			wantParseErr: true,
		},
		{
			name:         "empty string",
			nt:           parsit("", time.UTC),
			from:         ts,
			wantParseErr: true,
		},
		{
			name:         "white space",
			nt:           parsit("       ", time.UTC),
			from:         ts,
			wantParseErr: true,
		},
		{
			name:         "white space",
			nt:           parsit("       ", time.UTC),
			from:         ts,
			wantParseErr: true,
		},
		{
			name:         "incorrectly formatted range",
			nt:           parsit("a-7/7 1 1 1 1       ", time.UTC),
			from:         ts,
			wantParseErr: true,
		},
		{
			name: "correctly formatted range",
			nt:   parsit("2-7/7 1 1 1 0-7       ", time.UTC),
			from: ts,
			want: mustParseTime("2020-01-01T01:02:00Z"),
		},
		{
			name: "month ranges",
			nt:   parsit("* * * * 8-10/3 * *", time.UTC),
			from: ts,
			want: mustParseTime("2019-10-27T00:00:00Z"),
		},
		{
			name: "more ranges",
			nt:   parsit("2-20/2 * * * 8-10/3 * *", time.UTC),
			from: ts,
			want: mustParseTime("2019-10-27T00:00:02Z"),
		},
		{
			name:         "zeros in months",
			nt:           parsit("2-20/2 * * * 0-2 * *", time.UTC),
			from:         ts,
			wantParseErr: true,
		},
		{
			name:         "zeros in month ranges",
			nt:           parsit("2-20/2 * * * 0-2 * *", time.UTC),
			from:         ts,
			wantParseErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nt, err := tt.nt(t)
			if (!tt.wantParseErr) && (err != nil) {
				t.Errorf("expected no parse error but got %v", err)
				return // if parse errors we can't schedule.
			}
			if tt.wantParseErr && (err == nil) {
				t.Errorf("expected parse error but got nil")
				return
			}
			if tt.wantParseErr && err != nil {
				return
			}
			if got, err := nt.next(tt.from); !reflect.DeepEqual(got, tt.want) || tt.wanterr == (err == nil) {
				if (!tt.wanterr) && (err != nil) {
					t.Errorf("expected no error but got %v", err)
				}
				if tt.wanterr && (err == nil) {
					t.Errorf("expected error but got nil")
				}
				t.Errorf("from = %v, nextTime.next() = %v, want %v", tt.from, got, tt.want)
			}
		})
	}
}

func BenchmarkParse(b *testing.B) {
	nt := nextTime{}
	var err error
	b.Run("0 * * * * * *", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			nt, err = parse("0 * * * * * *", time.UTC)
		}
	})
	_, _ = nt, err
	b.Run("1-6/2 * * * Feb-Oct/3 * *", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			nt, err = parse("1-6/2 * * * Feb-Oct/3 * *", time.UTC)
		}
	})
	_, _ = nt, err
	b.Run("1-6/2 * * * Feb-Oct/3 * 2020/4", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			nt, err = parse("1-6/2 * * * Feb-Oct/3 * 2020/4", time.UTC)
		}
	})
	_, _ = nt, err
}

func BenchmarkNext(b *testing.B) {
	nt := nextTime{}
	var err error
	ts := time.Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		nt, err = parse("* * * * * * *", time.UTC)
		nt.next(ts)
	}
	_, _, _ = nt, err, ts
}
