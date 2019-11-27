package cron_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/influxdata/cron"
)

func BenchmarkCount(b *testing.B) {
	cases := []struct {
		cron string
		from time.Time
		to   time.Time
	}{
		{
			cron: "20,44 7 5 * * * *",
			from: mustParseTime("2019-12-02T00:00:00Z"),
			to:   mustParseTime("2020-03-30T00:00:00Z"),
		},
		{
			cron: "20,44 7 * * * * 2018",
			from: mustParseTime("2019-12-02T00:00:00Z"),
			to:   mustParseTime("2019-12-03T00:00:00Z"),
		},
		{
			cron: "* */2 * * * * 2023",
			from: mustParseTime("2013-02-03T01:02:04Z"),
			to:   mustParseTime("2013-02-03T01:20:55Z"),
		},
		{
			cron: "20,44 7 * * * *",
			from: mustParseTime("2019-12-02T00:00:00Z"),
			to:   mustParseTime("2019-12-03T00:00:00Z"),
		},
		{
			cron: "20,44 * * * * *",
			from: mustParseTime("2019-12-02T00:00:00Z"),
			to:   mustParseTime("2019-12-03T00:00:00Z"),
		},
		{
			cron: "*/7 * * * * Thu",
			from: mustParseTime("2019-12-02T00:00:00Z"),
			to:   mustParseTime("2019-12-09T05:01:03Z"),
		},
		{
			cron: "* * * * Thu",
			from: mustParseTime("2019-12-02T00:00:00Z"),
			to:   mustParseTime("2019-12-09T05:01:03Z"),
		},
		{
			cron: "* * * * * *",
			from: mustParseTime("2013-02-03T01:02:04Z"),
			to:   mustParseTime("2013-03-15T05:01:03Z"),
		},

		{
			cron: "* */2 * * * * 2023",
			from: mustParseTime("2013-02-03T01:02:04Z"),
			to:   mustParseTime("2013-02-03T01:20:55Z"),
		},
		{
			cron: "@every 1s",
			from: mustParseTime("2013-02-03T01:02:04Z"),
			to:   mustParseTime("2013-02-04T01:02:04Z"),
		},
		{
			cron: "@every 1d1h",
			from: mustParseTime("2013-02-03T01:02:04Z"),
			to:   mustParseTime("2013-02-04T01:02:04Z"),
		},
		{
			cron: "@every 1d100ms",
			from: mustParseTime("2013-02-03T00:00:00Z"),
			to:   mustParseTime("2013-02-04T01:02:04Z"),
		},
		{
			cron: "@every 1mo100d",
			from: mustParseTime("2013-01-03T00:00:00Z"),
			to:   mustParseTime("2013-11-20T01:02:04Z"),
		},
		{
			cron: "@every 2y",
			from: mustParseTime("2013-01-03T00:00:00Z"),
			to:   mustParseTime("2019-11-20T01:02:04Z"),
		},
	}

	for _, tt := range cases {
		b.Run(fmt.Sprintf("%s, [%s, %s)", tt.cron, tt.from, tt.to), func(b *testing.B) {
			nt, err := cron.ParseUTC(tt.cron)

			if err != nil {
				b.Fatal(err)
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				nt.Count(tt.from, tt.to)
			}
		})
	}
}

func TestParsed_Count(t *testing.T) {
	tests := []struct {
		cron      string
		from      time.Time
		to        time.Time
		wantCount int
	}{
		{
			cron:      "20,44 7 5 * * * *",
			from:      mustParseTime("2020-01-01T00:00:00Z"),
			to:        mustParseTime("2020-01-03T00:00:00Z"),
			wantCount: int(mustParseTime("2020-01-03T00:00:00Z").Sub(mustParseTime("2020-01-01T00:00:00Z")) / time.Hour / 12),
		},
		{
			cron:      "20,44 7 5 * * * *",
			from:      mustParseTime("2019-12-20T00:00:00Z"),
			to:        mustParseTime("2020-01-03T00:00:00Z"),
			wantCount: int(mustParseTime("2020-01-03T00:00:00Z").Sub(mustParseTime("2019-12-20T00:00:00Z")) / time.Hour / 12),
		},
		{
			cron:      "20,44 7 * * * * 2018",
			from:      mustParseTime("2019-12-02T00:00:00Z"),
			to:        mustParseTime("2019-12-03T00:00:00Z"),
			wantCount: 0,
		},
		{
			cron:      "20,44 7 * * * *",
			from:      mustParseTime("2019-12-02T00:00:00Z"),
			to:        mustParseTime("2019-12-03T00:00:00Z"),
			wantCount: 24 * 2,
		},
		{
			cron:      "20,44 * * * * *",
			from:      mustParseTime("2019-12-02T00:00:00Z"),
			to:        mustParseTime("2019-12-03T00:00:00Z"),
			wantCount: 24 * 60 * 2,
		},
		{
			cron:      "*/7 * * * * Thu",
			from:      mustParseTime("2019-12-02T00:00:00Z"),
			to:        mustParseTime("2019-12-09T05:01:03Z"),
			wantCount: 24 * 60 * 9,
		},
		{
			cron:      "* * * * Thu",
			from:      mustParseTime("2019-12-02T00:00:00Z"),
			to:        mustParseTime("2019-12-09T05:01:03Z"),
			wantCount: 24 * 60,
		},

		{
			cron:      "* * * * * Thu",
			from:      mustParseTime("2019-12-02T00:00:00Z"),
			to:        mustParseTime("2019-12-09T05:01:03Z"),
			wantCount: 24 * 60 * 60,
		},

		{
			cron:      "* * * * * *",
			from:      mustParseTime("2013-02-03T01:02:04Z"),
			to:        mustParseTime("2013-03-15T05:01:03Z"),
			wantCount: int(mustParseTime("2013-03-15T05:01:03Z").Sub(mustParseTime("2013-02-03T01:02:04Z").Add(-1)).Truncate(time.Second) / time.Second),
		},
		{
			cron:      "* * * * * *",
			from:      mustParseTime("2013-02-03T01:02:04Z"),
			to:        mustParseTime("2013-02-15T05:01:03Z"),
			wantCount: int(mustParseTime("2013-02-15T05:01:03Z").Sub(mustParseTime("2013-02-03T01:02:04Z").Add(-1)).Truncate(time.Second) / time.Second),
		},

		{
			cron:      "* * * * * *",
			from:      mustParseTime("2013-02-03T01:02:04Z"),
			to:        mustParseTime("2013-02-03T05:01:03Z"),
			wantCount: int(mustParseTime("2013-02-03T05:01:03Z").Sub(mustParseTime("2013-02-03T01:02:04Z").Add(-1)).Truncate(time.Second) / time.Second),
		},

		{
			cron:      "* * * * * *",
			from:      mustParseTime("2013-02-03T01:02:04Z"),
			to:        mustParseTime("2013-02-03T05:01:03Z"),
			wantCount: int(mustParseTime("2013-02-03T05:01:03Z").Sub(mustParseTime("2013-02-03T01:02:04Z").Add(-1)).Truncate(time.Second) / time.Second),
		},

		{
			cron:      "* * * * * *",
			from:      mustParseTime("2013-02-03T01:02:04Z"),
			to:        mustParseTime("2013-02-03T05:20:55Z"),
			wantCount: int(mustParseTime("2013-02-03T05:20:55Z").Sub(mustParseTime("2013-02-03T01:02:04Z").Add(-1)).Truncate(time.Second) / time.Second),
		},

		{
			cron:      "* * * * * *",
			from:      mustParseTime("2013-02-03T01:02:04Z"),
			to:        mustParseTime("2013-02-03T02:20:55Z"),
			wantCount: int(mustParseTime("2013-02-03T02:20:55Z").Sub(mustParseTime("2013-02-03T01:02:04Z").Add(-1)).Truncate(time.Second) / time.Second),
		},
		{
			cron:      "* */2 * * * * 2023",
			from:      mustParseTime("2013-02-03T01:02:04Z"),
			to:        mustParseTime("2013-02-03T01:20:55Z"),
			wantCount: 0,
		},
		{
			cron:      "*/2 * * * *",
			from:      mustParseTime("2013-02-03T01:02:04Z"),
			to:        mustParseTime("2013-02-03T01:20:55Z"),
			wantCount: 9,
		},

		{
			cron:      "* * * * *",
			from:      mustParseTime("2013-02-03T01:02:04Z"),
			to:        mustParseTime("2013-02-03T01:10:55Z"),
			wantCount: 8,
		},
		{
			cron:      "* * * * *",
			from:      mustParseTime("2013-02-03T01:02:04Z"),
			to:        mustParseTime("2013-02-03T01:03:55Z"),
			wantCount: 1,
		},
		{
			cron:      "* * * * * *",
			from:      mustParseTime("2013-02-03T01:02:04Z"),
			to:        mustParseTime("2013-02-03T01:03:55Z"),
			wantCount: int(mustParseTime("2013-02-03T01:03:55Z").Sub(mustParseTime("2013-02-03T01:02:04Z").Add(-1)).Truncate(time.Second) / time.Second),
		},
		{
			cron:      "* * * * * *",
			from:      mustParseTime("2013-02-03T01:02:04Z"),
			to:        mustParseTime("2013-02-03T01:02:55Z"),
			wantCount: int(mustParseTime("2013-02-03T01:02:55Z").Sub(mustParseTime("2013-02-03T01:02:04Z").Add(-1)).Truncate(time.Second) / time.Second),
		},
		{
			cron:      "@every 1s",
			from:      mustParseTime("2013-02-03T01:02:04Z"),
			to:        mustParseTime("2013-02-04T01:02:04Z"),
			wantCount: int(mustParseTime("2013-02-04T01:02:04Z").Sub(mustParseTime("2013-02-03T01:02:04Z").Add(-1)).Truncate(time.Second) / time.Second),
		},
		{
			cron:      "@every 1d",
			from:      mustParseTime("2013-02-03T01:02:04Z"),
			to:        mustParseTime("2013-02-04T01:02:04Z"),
			wantCount: 0,
		},
		{
			cron:      "@every 1d1h",
			from:      mustParseTime("2013-02-03T01:02:04Z"),
			to:        mustParseTime("2013-02-04T01:02:04Z"),
			wantCount: 0,
		},
		{
			cron:      "@every 1d",
			from:      mustParseTime("2013-02-03T01:02:04Z"),
			to:        mustParseTime("2013-02-04T01:02:04Z"),
			wantCount: 0,
		},
		{
			cron:      "@every 2d4s",
			from:      mustParseTime("2013-02-03T01:02:04Z"),
			to:        mustParseTime("2013-02-04T01:02:04Z"),
			wantCount: 0,
		},
		{
			cron:      "@every 1d100ms",
			from:      mustParseTime("2013-02-03T00:00:00Z"),
			to:        mustParseTime("2013-02-04T01:02:04Z"),
			wantCount: 1,
		},
		{
			cron:      "@every 3d",
			from:      mustParseTime("2013-02-03T00:00:00Z"),
			to:        mustParseTime("2013-05-20T01:02:04Z"),
			wantCount: int(mustParseTime("2013-05-20T01:02:04Z").Sub(mustParseTime("2013-02-03T00:00:00Z")).Truncate(time.Second) / (72 * time.Hour)),
		},
		{
			cron:      "@every 4d",
			from:      mustParseTime("2013-02-03T00:00:00Z"),
			to:        mustParseTime("2013-05-20T01:02:04Z"),
			wantCount: int(mustParseTime("2013-05-20T01:02:04Z").Sub(mustParseTime("2013-02-03T00:00:00Z")).Truncate(time.Second) / (96 * time.Hour)),
		},
		{
			cron:      "@every 4d12h",
			from:      mustParseTime("2013-02-03T00:00:00Z"),
			to:        mustParseTime("2013-05-20T01:02:04Z"),
			wantCount: int(mustParseTime("2013-05-20T01:02:04Z").Sub(mustParseTime("2013-02-03T00:00:00Z")).Truncate(time.Second) / (108 * time.Hour)),
		},
		{
			cron:      "@every 1mo",
			from:      mustParseTime("2013-02-03T00:00:00Z"),
			to:        mustParseTime("2013-05-20T01:02:04Z"),
			wantCount: 3,
		},
		{
			cron:      "@every 1mo100d",
			from:      mustParseTime("2013-01-03T00:00:00Z"),
			to:        mustParseTime("2013-11-20T01:02:04Z"),
			wantCount: 2,
		},
		{
			cron:      "@every 1y",
			from:      mustParseTime("2013-01-03T00:00:00Z"),
			to:        mustParseTime("2019-11-20T01:02:04Z"),
			wantCount: 6,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s, [%s, %s)", tt.cron, tt.from, tt.to), func(t *testing.T) {
			nt, err := cron.ParseUTC(tt.cron)
			if err != nil {
				t.Fatal(err)
			}
			if gotCount := nt.Count(tt.from, tt.to); gotCount != tt.wantCount {
				t.Fatalf("Parsed.Count() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
