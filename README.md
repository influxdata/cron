# cron
A fast non-allocating ron parser in ragel and golang.


# features
- [x] standard five-position cron
- [x] six-position cron with seconds
- [x] seven-position cron with seconds and years
- [x] case-insensitive days of the week
- [x] case-insensitive month names
- [x] [Quartz](http://www.quartz-scheduler.org) compatible ranges e.g.: `4/10`, `SUN-THURS/2`
- [ ] timezone handling (other than UTC)
- [ ] Quartz compatible # handling, i.e. `5#3` meaning the third friday of the month
- [ ] Quartz compatible L handling, i.e. `5L` in the day of week field meaning the last friday of a month, or `3L` in the day of month field meaning third to last day of the month

# performance
On a MacBook Pro (Retina, 15-inch, Mid 2014):

| | | | |
|-|-|-|-|
| BenchmarkParse/0_*_*_*_*_*_*-8 | 20000000 |  97.9 ns/op |  0 B/op | 0 allocs/op |
| BenchmarkParse/1-6/2_*_*_*_Feb-Oct/3_*_*-8 | 10000000 |  184  ns/op |  0 B/op | 0 allocs/op |
| BenchmarkParse/1-6/2_*_*_*_Feb-Oct/3_*_2020/4-8 |  1000000 |  2262 ns/op |  0 B/op | 0 allocs/op |

As you can see the parsing is pretty fast.  That being said, ranges over years are much slower to parse than basically anything else, as they were likely to be the least used so I didn't bother optimizing their bitmap generation.

# TODO
- I'd like to make year ranges more performant.
- Once we are done adding all of the Quartz cron features, to copy and pass all of Quartz's parsing tests.