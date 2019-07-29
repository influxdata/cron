package cron_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/influxdata/cron"
)

func ceilTime(t time.Time, d time.Duration) time.Time {
	if t.Truncate(d).Equal(t) {
		return t
	}
	return t.Add(d).Truncate(d)
}

func mustParseTime(s string) time.Time {
	ts, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		panic(err)
	}
	return ts

}

func Test_Parsed_next(t *testing.T) {
	ts, err := time.Parse(time.RFC3339Nano, "2019-06-27T06:05:12.622652Z")
	_ = ts
	if err != nil {
		t.Fatal(err)
	}
	parsit := func(s string) func(t *testing.T) (cron.Parsed, error) {
		return func(t *testing.T) (cron.Parsed, error) {
			return cron.ParseUTC(s)
		}
	}
	tests := []struct {
		name         string
		nt           func(t *testing.T) (cron.Parsed, error)
		from         time.Time
		want         time.Time
		wanterr      bool
		wantParseErr bool
	}{
		{
			name: "2023",
			nt:   parsit(" *   * * * * * 2023 "),
			from: ts,
			want: mustParseTime("2023-01-01T00:00:00Z"),
		},
		{
			name: "all stars",
			nt:   parsit("* * * * * 4-6 *"),
			from: ts,
			want: ceilTime(ts, time.Second),
		},
		{
			name: "all seconds starting from ten by 7",
			nt:   parsit("10/4 * * * * * *"),
			from: ts,
			want: mustParseTime("2019-06-27T06:05:14Z"),
		},
		{
			name: "all seconds between 3 and 12 by 3",
			nt:   parsit("3-12/2 * * * * * *"),
			from: ts,
			want: mustParseTime("2019-06-27T06:06:03Z"),
		},
		{
			name: "all minutes between 4 and 22 by 3",
			nt:   parsit("* 3-22/3 * * * * *"),
			from: ts,
			want: mustParseTime("2019-06-27T06:06:00Z"),
		},
		{
			name: "at every 5nd minute starting at minute 3.",
			nt:   parsit("* 3/5 * * * * *"),
			from: ts,
			want: mustParseTime("2019-06-27T06:08:00Z"),
		},
		{
			name: "at every 2nd minute starting at minute 3 for even seconds.",
			nt:   parsit("*/2 3/2 * * * * *"),
			from: ts,
			want: mustParseTime("2019-06-27T06:05:14Z"),
		},
		{
			name: "once a minute at second 1",
			nt:   parsit("1 * * * * * *"),
			from: ts,
			want: mustParseTime("2019-06-27T06:06:01Z"),
		},
		{
			name: "all stars extra space at end",
			nt:   parsit("* * * * * * * "),
			from: ts,
			want: ceilTime(ts, time.Second),
		},
		{
			name: "all stars extra space at beginning",
			nt:   parsit(" * * * * * * *"),
			from: ts,
			want: ceilTime(ts, time.Second),
		},
		{
			name: "all stars extra spaces randomly ",
			nt:   parsit(" * * *  * \t*\t* * "),
			from: ts,
			want: ceilTime(ts, time.Second),
		},
		{
			name: "fourth second",
			nt:   parsit("4 * * * * * *"),
			from: ts,
			want: ceilTime(ts, time.Minute).Add(4 * time.Second),
		},
		{
			name: "every fourth second",
			nt:   parsit("*/4 * * * * * *"),
			from: ts,
			want: mustParseTime("2019-06-27T06:05:16Z"),
		},
		{
			name: "every fourth second and every 2nd hour",
			nt:   parsit("*/4 * */2 * * * *"),
			from: ts,
			want: mustParseTime("2019-06-27T06:05:16Z"),
		},
		{
			name: "every fourth second and every 4th hour between 4 and 16",
			nt:   parsit("*/4 * 4-16/4 * * * *"),
			from: ts,
			want: mustParseTime("2019-06-27T08:00:00Z"),
		},
		{
			name: "every second except on thursday",
			nt:   parsit("* * * * * FRI,SAT,SUN,MON,TUE,WED *"),
			from: ts,
			want: mustParseTime("2019-06-28T00:00:00Z"),
		},
		{
			name: "every second on saturday",
			nt:   parsit("* * * * * SAT *"),
			from: ts,
			want: mustParseTime("2019-06-29T00:00:00Z"),
		},
		{
			name: "fourth second, 7th minute",
			nt:   parsit("4 7 * * * * *"),
			from: ts,
			want: mustParseTime("2019-06-27T06:07:04Z"),
		},
		{
			name: "fourth second, 7th minute on saturday",
			nt:   parsit("4 7 * * * 6 *"),
			from: ts,
			want: mustParseTime("2019-06-29T00:07:04Z"),
		},
		{
			name: "fourth second, 5-10th minute on saturday",
			nt:   parsit("4 5-10 * * * 6 *"),
			from: ts,
			want: mustParseTime("2019-06-29T00:05:04Z"),
		},
		{
			name: "noon on dec 3rd",
			nt:   parsit("* * 12 3 12 * *"),
			from: ts,
			want: mustParseTime("2019-12-03T12:00:00Z"),
		},
		{
			name: "noon on leap day",
			nt:   parsit("* * 12 29 FEB * *"),
			from: ts,
			want: mustParseTime("2020-02-29T12:00:00Z"),
		},
		{
			name: "noon on leap day on day on wed",
			nt:   parsit("* * 12 29 FEB WED *"),
			from: ts,
			want: mustParseTime("2040-02-29T12:00:00Z"),
		},
		{
			name: "noon on leap day on day on tuesday,wednesday, or thursday",
			nt:   parsit("* * 12 29 FEB tue-thu *"),
			from: ts,
			want: mustParseTime("2024-02-29T12:00:00Z"),
		},
		{
			name: "noon on leap day on day on Monday or sunday",
			nt:   parsit("* * 12 29 2 MON,0 *"),
			from: ts,
			want: mustParseTime("2032-02-29T12:00:00Z"),
		},
		{
			name: "when july 4 is a weekday",
			nt:   parsit("* * * 4 7 MON,TUE,WED,4,5 *"),
			from: ts,
			want: mustParseTime("2019-07-04T00:00:00Z"),
		},
		{
			name:    "when leap day is a monday and the year is a multiple of 3 after 1970",
			nt:      parsit("* * * 29 2 MON */3"),
			from:    ts,
			want:    time.Time{},
			wanterr: true,
		},
		{
			name: "jan 1 falls on a monday and year is a multiple of 3 after 1970",
			nt:   parsit("* * * 1 1 MON */3"),
			from: ts,
			want: mustParseTime("2046-01-01T00:00:00Z"),
		},
		{
			name: "jan 1 on year is divisible by 2",
			nt:   parsit("* * * * * * */2"),
			from: ts,
			want: mustParseTime("2020-01-01T00:00:00Z"),
		},
		{
			name: "6 position (no year)",
			nt:   parsit("* * * * * *"),
			from: ts,
			want: mustParseTime("2019-06-27T06:05:13Z"),
		},
		{
			name: "5 position (no year, no sec) extra spaces",
			nt:   parsit("  * *  * * *  "),
			from: ts,
			want: mustParseTime("2019-06-27T06:06:00Z"),
		},
		{
			name: "1970",
			nt:   parsit("*  * *  * * *  1970"),
			from: time.Time{},
			want: mustParseTime("1970-01-01T00:00:00Z"),
		},
		{
			name: "2097",
			nt:   parsit("*  * *  * * *  2097"),
			from: time.Time{},
			want: mustParseTime("2097-01-01T00:00:00Z"),
		},
		{
			name: "yearly",
			nt:   parsit("@yearly"),
			from: ts,
			want: mustParseTime("2020-01-01T00:00:00Z"),
		},
		{
			name: "annually",
			nt:   parsit("@annually"),
			from: ts,
			want: mustParseTime("2020-01-01T00:00:00Z"),
		},
		{
			name:         "fred likes snickerdoodles",
			nt:           parsit("fred likes snickerdoodles"),
			from:         ts,
			wantParseErr: true,
		},
		{
			name:         "@notanemailaddress",
			nt:           parsit("@notanemailaddress"),
			from:         ts,
			wantParseErr: true,
		},
		{
			name:         "@yearlyplusextraletters",
			nt:           parsit("@yearlyplusextraletters"),
			from:         ts,
			wantParseErr: true,
		},
		{
			name:         "too many fields",
			nt:           parsit("* * 8 8 1 1 0"),
			from:         ts,
			wantParseErr: true,
		},
		{
			name:         "too many fields",
			nt:           parsit("* * 8 8 1 1 0"),
			from:         ts,
			wantParseErr: true,
		},
		{
			name:         "error instead of summoning the old ones",
			nt:           parsit("c̴̡͈̪̪̬͙͈̼̞̥̩̜̪̳̜̰̬̱͇̮̼̳̟̩̳̫̾̃͐̽̒̎̈́̽͠͝͝͠t̴͍̫͓͉̥̤̺͇͇̦̮̞͈̬̮͕͖̼̺̭̥͓̬̣͉̻̭̣̓̐̿̇͊̈́̉̏͐̆͑͂͛̎͒͌̈́͋̃̾̔̕͝h̸̜̭͕͎̰̮̭͙̙͇͎̺̤̙͖̒͌̂̌̍̿̒̀̊̿̾̚̕̕͝͠ͅū̷̲̝̂̑̈́͊͋̓̆̈́́́̄̎̄̃͐̈́̇̚͝ͅl̷̢̡̛͉̠̞̮̝͙̻̜͓̹̻̩̯̙̹̯͓̝͇̫̫̠̿̑̇̓̀̂̀͗̒͋̀͌͒̕̕͜͠ͅh̵̡̛̖͉̟̞̣̖̬͓̜̄̀̂̽̄͒̅̀͋̓̀͌̆̂̀͌͌̉̀̐̈́͋̓͂̚̕͠͠ų̶̛͎̖͖̜̹̬͙̖̜̲̺͐̀̎̎̽͛͌͆̂̌͒̋̆̿̎̎̽̑̏̏̈̾͘̕̕̕̕̕̚͠ͅ ̸̢̨̡̛̙̺̹͖͙͔̩̮͙̝̳̘̰̣̪̦̥̰̠̝̼͕͔͌̑̃́̐̈͗̍̅̽̌̈́̑̉̋̃͘͜͝͠͠f̶̹̹͉̉ä̸̲͈̺͉̮̹̘̦̦̟̘̼̪͚̯͓̻͉̣̬͓̰̟͇͕̳̟̗̉̈̔̉̿̔̇̑̕͜g̴̢̧̡̱̞̭͉̼̤̟̗͚͖̩̖̤͉͉̤̪͖̣̠̜̟̟̜̜̺̯̭̋̐̌̆͒͌̓̒͛́̌͒̀̓̆͋͐̈́̐̚͝ň̴̨̢̨̡͓̥̪̬͍̙̳̫̰̲̗̰̥̖̫̹̱̟̰̒̿̓̀̒̊́͛̊̈́̒̌̂͐́̕͝"),
			from:         ts,
			wantParseErr: true,
		},
		{
			name:         "empty string",
			nt:           parsit(""),
			from:         ts,
			wantParseErr: true,
		},
		{
			name:         "white space",
			nt:           parsit("       "),
			from:         ts,
			wantParseErr: true,
		},
		{
			name:         "white space",
			nt:           parsit("       "),
			from:         ts,
			wantParseErr: true,
		},
		{
			name:         "incorrectly formatted range",
			nt:           parsit("a-7/7 1 1 1 1       "),
			from:         ts,
			wantParseErr: true,
		},
		{
			name: "correctly formatted range",
			nt:   parsit("2-7/7 1 1 1 0-7       "),
			from: ts,
			want: mustParseTime("2020-01-01T01:02:00Z"),
		},
		{
			name: "month ranges",
			nt:   parsit("* * * * 8-10/3 * *"),
			from: ts,
			want: mustParseTime("2019-10-27T00:00:00Z"),
		},
		{
			name: "more ranges",
			nt:   parsit("2-20/2 * * * 8-10/3 * *"),
			from: ts,
			want: mustParseTime("2019-10-27T00:00:02Z"),
		},
		{
			name:         "zeros in months",
			nt:           parsit("2-20/2 * * * 0-2 * *"),
			from:         ts,
			wantParseErr: true,
		},
		{
			name:         "zeros in month ranges",
			nt:           parsit("2-20/2 * * * 0-2 * *"),
			from:         ts,
			wantParseErr: true,
		},
		{
			name: "1970",
			nt:   parsit("* * * * * * 1970"),
			from: mustParseTime("1960-01-01T00:00:02Z"),
			want: mustParseTime("1970-01-01T00:00:00Z"),
		},
		{
			name: "2099",
			nt:   parsit("* * * * * * 2099"),
			from: mustParseTime("1960-01-01T00:00:02Z"),
			want: mustParseTime("2099-01-01T00:00:00Z"),
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
			if got, err := nt.Next(tt.from); !reflect.DeepEqual(got, tt.want) || tt.wanterr == (err == nil) {
				if (!tt.wanterr) && (err != nil) {
					t.Errorf("expected no error but got %v", err)
				}
				if tt.wanterr && (err == nil) {
					t.Errorf("expected error but got nil")
				}
				t.Errorf("from = %v,cron.Parsed.Next() = %v, want %v", tt.from, got, tt.want)
			}
		})
	}
}

func BenchmarkParse(b *testing.B) {
	nt := cron.Parsed{}
	var err error
	b.Run("0 * * * * * *", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			nt, err = cron.ParseUTC("0 * * * * * *")
		}
	})
	_, _ = nt, err
	b.Run("1-6/2 * * * Feb-Oct/3 * *", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			nt, err = cron.ParseUTC("1-6/2 * * * Feb-Oct/3 * *")
		}
	})
	_, _ = nt, err
	b.Run("1-6/2 * * * Feb-Oct/3 * 2020/4", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			nt, err = cron.ParseUTC("1-6/2 * * * Feb-Oct/3 * 2020/4")
		}
	})
	_, _ = nt, err
}

func BenchmarkNext(b *testing.B) {
	nt := cron.Parsed{}
	var err error
	ts := time.Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		nt, err = cron.ParseUTC("* * * * * * *")
		nt.Next(ts)
	}
	_, _, _ = nt, err, ts
}

func ExampleParseUTC() {
	p, err := cron.ParseUTC("10 * * * * *")
	if err != nil {
		fmt.Println(err)
	}
	ts, err := p.Next(time.Date(1999, 12, 31, 23, 59, 59, 0, time.UTC))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ts)
	// Output:
	// 2000-01-01 00:00:10 +0000 UTC
}
