// -*-go-*-
package cron

import (
    "strconv"
    "fmt"
    "time"
    "errors"
)

%% machine parse;
%% variable data s;
%% write data;

func parse(s string, tz *time.Location)(nextTime, error){
    nt:=nextTime{loc:tz}
    cs, p, pe, eof:= 0, 0,len(s), len(s)


    // init scanner vars
    act, ts, te := 0, 0, 0
    _, _, _ = act, ts, te // we have to do this.

    // for fcall
    top := 0
    _ = top
    stack := [8]int{}
    mark, backtrack := 0,0
    _ = mark
    _ = backtrack
    m, d, start, end:=uint64(0), uint64(0), uint64(0), uint64(0)
    _ = d
    // TODO(docmerlin): handle ranges
    var err error
    %% write init;
    //m,h := 1<<0,1<<0
    %%{
        action mark {
            mark = p;
        }
        action testLen { backtrack >= 5 }
        action appendsecond {
            fmt.Println("m", m)
            if m>60 {
                return nt, fmt.Errorf("invalid minute %d", m)
            }
            nt.second |= 1<<m
        }

        action appendmonth {
            if m>12 || m<1{
                return nt, fmt.Errorf("invalid month %d", m)
            }
            nt.month |= 1<<(m-1) // we zero index months
        }

        action appendminute {
            if m>60 {
                return nt, fmt.Errorf("invalid minute %d", m)
            }
            nt.minute |= 1<<m
        }

        action appenddow {
            {
                const sundaysAtFirst = uint64(1 | 1<<7 | 1<<14 | 1<<21 | 1<<28 | 1<< 35 | 1<<42)
                if m==7 { // early return for sunday for compatibility with some crons that allow sunday to be 0 or 7
                    nt.dow |= sundaysAtFirst
                    return nt, nil
                }
                if m>6 || m<0 {
                    return nt, fmt.Errorf("invalid day of week %d", m)
                }
                nt.dow |= sundaysAtFirst<<m
            }
        }

        action appendhour {
            if m>24 || m <0 {
                return nt, fmt.Errorf("invalid hour of day %d", m)
            }
            nt.hour |= 1<<m
        }

        action appenddom {
            if m>30 || m <0 {
                return nt,  fmt.Errorf("invalid day of month %d", m)
            }
            nt.dom |= 1<<(m-1)// again we zero index day
        }

        action appendyear {
            if m<1970 || m>2099{
                return nt, errors.New("only years on interval [1970, 2099] are supported")
            }
            nt.year.set(int(m))
        }

        action appendStarSlashSeconds {
            {
                x:= [...]uint64{
                    1|1<<1|1<<2|1<<3|1<<4|1<<5|1<<6|1<<7|1<<8|1<<9|1<<10|1<<11|1<<12|1<<13|1<<14|1<<15|1<<16|1<<17|1<<18|1<<19|1<<20|1<<21|1<<22|1<<23|1<<24|1<<25|1<<26|1<<27|1<<28|1<<29|1<<30|1<<31|1<<32|1<<33|1<<34|1<<35|1<<36|1<<37|1<<38|1<<39|1<<40|1<<41|1<<42|1<<43|1<<44|1<<45|1<<46|1<<47|1<<48|1<<49|1<<50|1<<51|1<<52|1<<53|1<<54|1<<55|1<<56|1<<57|1<<58|1<<59,
                    1|1<<2|1<<4|1<<6|1<<8|1<<10|1<<12|1<<14|1<<16|1<<18|1<<20|1<<22|1<<24|1<<26|1<<28|1<<30|1<<32|1<<34|1<<36|1<<38|1<<40|1<<42|1<<44|1<<46|1<<48|1<<50|1<<52|1<<54|1<<56|1<<58,
                    1|1<<3|1<<6|1<<9|1<<12|1<<15|1<<18|1<<21|1<<24|1<<27|1<<30|1<<33|1<<36|1<<39|1<<42|1<<45|1<<48|1<<51|1<<54|1<<57,
                    1|1<<4|1<<8|1<<12|1<<16|1<<20|1<<24|1<<28|1<<32|1<<36|1<<40|1<<44|1<<48|1<<52|1<<56,
                    1|1<<5|1<<10|1<<15|1<<20|1<<25|1<<30|1<<35|1<<40|1<<45|1<<50|1<<55,
                    1|1<<6|1<<12|1<<18|1<<24|1<<30|1<<36|1<<42|1<<48|1<<54,
                    1|1<<7|1<<14|1<<21|1<<28|1<<35|1<<42|1<<49|1<<56,
                    1|1<<8|1<<16|1<<24|1<<32|1<<40|1<<48|1<<56,
                    1|1<<9|1<<18|1<<27|1<<36|1<<45|1<<54,
                    1|1<<10|1<<20|1<<30|1<<40|1<<50,
                    1|1<<11|1<<22|1<<33|1<<44|1<<55,
                    1|1<<12|1<<24|1<<36|1<<48,
                    1|1<<13|1<<26|1<<39|1<<52,
                    1|1<<14|1<<28|1<<42|1<<56,
                    1|1<<15|1<<30|1<<45,
                    1|1<<16|1<<32|1<<48,
                    1|1<<17|1<<34|1<<51,
                    1|1<<18|1<<36|1<<54,
                    1|1<<19|1<<38|1<<57,
                    1|1<<20|1<<40,
                    1|1<<21|1<<42,
                    1|1<<22|1<<44,
                    1|1<<23|1<<46,
                    1|1<<24|1<<48,
                    1|1<<25|1<<50,
                    1|1<<26|1<<52,
                    1|1<<27|1<<54,
                    1|1<<28|1<<56,
                    1|1<<29|1<<58,
                    1|1<<30,
                    1|1<<31,
                    1|1<<32,
                    1|1<<33,
                    1|1<<34,
                    1|1<<35,
                    1|1<<36,
                    1|1<<37,
                    1|1<<38,
                    1|1<<39,
                    1|1<<40,
                    1|1<<41,
                    1|1<<42,
                    1|1<<43,
                    1|1<<44,
                    1|1<<45,
                    1|1<<46,
                    1|1<<47,
                    1|1<<48,
                    1|1<<49,
                    1|1<<50,
                    1|1<<51,
                    1|1<<52,
                    1|1<<53,
                    1|1<<54,
                    1|1<<55,
                    1|1<<56,
                    1|1<<57,
                    1|1<<58,
                    1|1<<59,
                }
                if d>=60 {
                    return nt, fmt.Errorf("invalid second */%d", d)
                }
                if start<0 || start>=60 {
                    return nt, fmt.Errorf("invalid start second %d", start)
                }
                if end>=60 {
                    return nt, fmt.Errorf("invalid end second %d", start)
                }
                // handle the case that isn't a 
                endOp := 64-end
                if end==0{
                    endOp = 0
                }
                endMask := (^uint64(0))<<endOp>>endOp
                if d==0{
                    nt.second |= 1<<m
                }else{
                    nt.second |= (x[d-1]<<start)&endMask
                }
                fmt.Println("appendStarSlashSeconds", endOp,endMask, (x[d]<<start))
            }
        }

        action appendStarSlashMinutes {
            {
                // I really wish that golang had a binary bit representation, this is pretty aweful, but easier to understand than using hex or octal I guess
                x:= [...]uint64{
                    1|1<<1|1<<2|1<<3|1<<4|1<<5|1<<6|1<<7|1<<8|1<<9|1<<10|1<<11|1<<12|1<<13|1<<14|1<<15|1<<16|1<<17|1<<18|1<<19|1<<20|1<<21|1<<22|1<<23|1<<24|1<<25|1<<26|1<<27|1<<28|1<<29|1<<30|1<<31|1<<32|1<<33|1<<34|1<<35|1<<36|1<<37|1<<38|1<<39|1<<40|1<<41|1<<42|1<<43|1<<44|1<<45|1<<46|1<<47|1<<48|1<<49|1<<50|1<<51|1<<52|1<<53|1<<54|1<<55|1<<56|1<<57|1<<58|1<<59,
                    1|1<<2|1<<4|1<<6|1<<8|1<<10|1<<12|1<<14|1<<16|1<<18|1<<20|1<<22|1<<24|1<<26|1<<28|1<<30|1<<32|1<<34|1<<36|1<<38|1<<40|1<<42|1<<44|1<<46|1<<48|1<<50|1<<52|1<<54|1<<56|1<<58,
                    1|1<<3|1<<6|1<<9|1<<12|1<<15|1<<18|1<<21|1<<24|1<<27|1<<30|1<<33|1<<36|1<<39|1<<42|1<<45|1<<48|1<<51|1<<54|1<<57,
                    1|1<<4|1<<8|1<<12|1<<16|1<<20|1<<24|1<<28|1<<32|1<<36|1<<40|1<<44|1<<48|1<<52|1<<56,
                    1|1<<5|1<<10|1<<15|1<<20|1<<25|1<<30|1<<35|1<<40|1<<45|1<<50|1<<55,
                    1|1<<6|1<<12|1<<18|1<<24|1<<30|1<<36|1<<42|1<<48|1<<54,
                    1|1<<7|1<<14|1<<21|1<<28|1<<35|1<<42|1<<49|1<<56,
                    1|1<<8|1<<16|1<<24|1<<32|1<<40|1<<48|1<<56,
                    1|1<<9|1<<18|1<<27|1<<36|1<<45|1<<54,
                    1|1<<10|1<<20|1<<30|1<<40|1<<50,
                    1|1<<11|1<<22|1<<33|1<<44|1<<55,
                    1|1<<12|1<<24|1<<36|1<<48,
                    1|1<<13|1<<26|1<<39|1<<52,
                    1|1<<14|1<<28|1<<42|1<<56,
                    1|1<<15|1<<30|1<<45,
                    1|1<<16|1<<32|1<<48,
                    1|1<<17|1<<34|1<<51,
                    1|1<<18|1<<36|1<<54,
                    1|1<<19|1<<38|1<<57,
                    1|1<<20|1<<40,
                    1|1<<21|1<<42,
                    1|1<<22|1<<44,
                    1|1<<23|1<<46,
                    1|1<<24|1<<48,
                    1|1<<25|1<<50,
                    1|1<<26|1<<52,
                    1|1<<27|1<<54,
                    1|1<<28|1<<56,
                    1|1<<29|1<<58,
                    1|1<<30,
                    1|1<<31,
                    1|1<<32,
                    1|1<<33,
                    1|1<<34,
                    1|1<<35,
                    1|1<<36,
                    1|1<<37,
                    1|1<<38,
                    1|1<<39,
                    1|1<<40,
                    1|1<<41,
                    1|1<<42,
                    1|1<<43,
                    1|1<<44,
                    1|1<<45,
                    1|1<<46,
                    1|1<<47,
                    1|1<<48,
                    1|1<<49,
                    1|1<<50,
                    1|1<<51,
                    1|1<<52,
                    1|1<<53,
                    1|1<<54,
                    1|1<<55,
                    1|1<<56,
                    1|1<<57,
                    1|1<<58,
                    1|1<<59,
                }
                if d>=60 {
                    return nt, fmt.Errorf("invalid second */%d", d)
                }
                if start<0 || start>=60 {
                    return nt, fmt.Errorf("invalid start second %d", start)
                }
                if end>=60 {
                    return nt, fmt.Errorf("invalid end second %d", start)
                }
                // handle the case that isn't a 
                endOp := 64-end
                if end==0{
                    endOp = 0
                }
                endMask := (^uint64(0))<<endOp>>endOp
                if d==0{
                    nt.minute |= 1<<m
                }else{
                    nt.minute |= (x[d-1]<<start)&endMask
                }
            }
        }

        action appendStarSlashHours {
            {
                x:= [...]uint64{
                    1|1<<1|1<<2|1<<3|1<<4|1<<5|1<<6|1<<7|1<<8|1<<9|1<<10|1<<11|1<<12|1<<13|1<<14|1<<15|1<<16|1<<17|1<<18|1<<19|1<<20|1<<21|1<<22|1<<23,
                    1|1<<2|1<<4|1<<6|1<<8|1<<10|1<<12|1<<14|1<<16|1<<18|1<<20|1<<22,
                    1|1<<3|1<<6|1<<9|1<<12|1<<15|1<<18|1<<21,
                    1|1<<4|1<<8|1<<12|1<<16|1<<20,
                    1|1<<5|1<<10|1<<15|1<<20,
                    1|1<<6|1<<12|1<<18,
                    1|1<<7|1<<14|1<<21,
                    1|1<<8|1<<16,
                    1|1<<9|1<<18,
                    1|1<<10|1<<20,
                    1|1<<11|1<<22,
                    1|1<<12,
                    1|1<<13,
                    1|1<<14,
                    1|1<<15,
                    1|1<<16,
                    1|1<<17,
                    1|1<<18,
                    1|1<<19,
                    1|1<<20,
                    1|1<<21,
                    1|1<<22,
                    1|1<<23,
                }
                if m<1 || m>=24 {
                    return nt,  fmt.Errorf("invalid month */%d", m)
                }
                nt.hour |= x[m-1]
            }
        }

        action appendStarSlashMonths {
            {
                x:= [...]uint64{
                    1|1<<1|1<<2|1<<3|1<<4|1<<5|1<<6|1<<7|1<<8|1<<9|1<<10|1<<11,
                    1<<1|1<<3|1<<5|1<<7|1<<9|1<<11,
                    1<<2|1<<5|1<<8|1<<11,
                    1<<3|1<<7|1<<11,
                    1<<4|1<<9,
                    1<<5|1<<11,
                    1<<6,
                    1<<7,
                    1<<8,
                    1<<9,
                    1<<10,
                    1<<11,
                }
                if m<1 || m>=12 {
                    return nt,  fmt.Errorf("invalid month */%d", m)
                }
                nt.month |= x[m-1]
            }
        }

        action appendStarSlDoW {
            fmt.Println("here")
            {
                const sundaysAtFirst = uint64(1 | 1<<7 | 1<<14 | 1<<21 | 1<<28 | 1<<35 | 1<<42)
                x:= [...]uint64{
                    sundaysAtFirst|sundaysAtFirst<<1|sundaysAtFirst<<2|sundaysAtFirst<<3|sundaysAtFirst<<4|sundaysAtFirst<<5|sundaysAtFirst<<6, //every day
                    sundaysAtFirst<<1|sundaysAtFirst<<3|sundaysAtFirst<<5|sundaysAtFirst<<7, //every 2nd day
                    sundaysAtFirst<<2|sundaysAtFirst<<5|sundaysAtFirst<<8, //every 3rd day
                    sundaysAtFirst<<3|sundaysAtFirst<<6|sundaysAtFirst<<9, //every 4th day
                    sundaysAtFirst<<4|sundaysAtFirst<<8, //every 5th day
                    sundaysAtFirst<<5, //every 6th day
                    sundaysAtFirst<<6, //every 7th day
                }
                if m==7{
                    m=1
                }
                if m<1 || m>6 {
                    return nt,  fmt.Errorf("invalid day of week */%d", m)
                }
                nt.dow |= x[m-1]
            }
        }
        # this isn't optimized, because I really doubt people will use it very often.  If someone wants to optimize it, go ahead
        action appendStarSlYears {
            // short circuit for the most common cases.
            if m==1{
                nt.year.low=^uint64(0)
                nt.year.high=^uint64(0)
            }else{
                for i:=uint64(0);i<2099;i++{
                    if i%m==0{
                        nt.year.set(int(i))
                    }
                }
            }
        }
        action printf {
            fmt.Println("vars", cs, p, pe)
        }

        action appendStarSlDoM {
            {
                x:= [...]uint64{
                    1|1<<1|1<<2|1<<3|1<<4|1<<5|1<<6|1<<7|1<<8|1<<9|1<<10|1<<11|1<<12|1<<13|1<<14|1<<15|1<<16|1<<17|1<<18|1<<19|1<<20|1<<21|1<<22|1<<23|1<<24|1<<25|1<<26|1<<27|1<<28|1<<29|1<<30,
                    1<<1|1<<3|1<<5|1<<7|1<<9|1<<11|1<<13|1<<15|1<<17|1<<19|1<<21|1<<23|1<<25|1<<27|1<<29,
                    1<<2|1<<5|1<<8|1<<11|1<<14|1<<17|1<<20|1<<23|1<<26|1<<29,
                    1<<3|1<<7|1<<11|1<<15|1<<19|1<<23|1<<27,
                    1<<4|1<<9|1<<14|1<<19|1<<24|1<<29,
                    1<<5|1<<11|1<<17|1<<23|1<<29,
                    1<<6|1<<13|1<<20|1<<27,
                    1<<7|1<<15|1<<23,
                    1<<8|1<<17|1<<26,
                    1<<9|1<<19|1<<29,
                    1<<10|1<<21,
                    1<<11|1<<23,
                    1<<12|1<<25,
                    1<<13|1<<27,
                    1<<14|1<<29,
                    1<<15,
                    1<<16,
                    1<<17,
                    1<<18,
                    1<<19,
                    1<<20,
                    1<<21,
                    1<<22,
                    1<<23,
                    1<<24,
                    1<<25,
                    1<<26,
                    1<<27,
                    1<<28,
                    1<<29,
                    1<<30,
                }
                if m<0 || m>30 {
                    return nt,  fmt.Errorf("invalid day of month */%d", m)
                }
                nt.dom |= x[m-1]
            }
        }

        action incBacktrack {
            backtrack++
        }

        action len_err {
            return nt, fmt.Errorf("too many positions in cron")
        }

        action parse_err {
            fhold;
            return nt, fmt.Errorf("error in parsing at char %d, '%s'", p, s[p:p+1])
        }

        digits = (digit+) >mark %{
            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        };
        allowedNonSpace = alnum|"/"|"*"|","|"-";
        slash = "/";
        comma = ",";
        hypen = "-";

        dow = ( /SUN/i @{m=0} | /MON/i @{m=1} | /TUE/i @{m=2} | /WED/i @{m=3} | /THU/i @{m=4} | /FRI/i @{m=5} | /SAT/i @{m=6} );
        monthName = ( /JAN/i @{m=1} | /FEB/i @{m=2} | /MAR/i @{m=3} | /APR/i @{m=4} | /MAY/i @{m=5} | /JUN/i @{m=6} | /JUL/i @{m=7} | /AUG/i @{m=8} | /SEP/i @{m=9} | /OCT/i @{m=10} | /NOV/i @{m=11} | /DEC/i @{m=12} ) ;
        last = ( digit+ "L" ) >mark %{
            m, err = strconv.ParseUint(s[mark:p-1], 10, 64)
            if err!=nil {
                return nt,  err
            }
        };
        digitlist = digits ("," space* digits)*;
        starSlashDigits = ( ("*" @{m=1})( slash %mark digits )? );
        digitsAndSlashList = ( starSlashDigits | digits ) ( ',' ( starSlashDigits | digits ) )*;

        #range = (digits @{s=m;} hyphen digits @{e=m;})
        secminrange = ( ( "*" %{ start=0;end=59;m=1;d=1; } ) |  (digits %{ start=m; end=0;d=0;} ( "-" digits %{ end=m; d=1;} )? )) >{d=0;};
        second =  ( secminrange ("/" digits %{d=m})? ) %appendStarSlashSeconds; 
        seconds = second ("," second) *;

        
        #( "/" digits %appendStarSlashSeconds ) ) ( "," ( starSlashDigits %appendStarSlashSeconds | digits %appendsecond ) )*;
        minute = ( secminrange ("/" digits %{d=m})? ) %appendStarSlashMinutes; 
        minutes = minute ("," minute) *;
        hours = ( starSlashDigits %appendStarSlashHours | digits %appendhour ) ( "," ( starSlashDigits %appendStarSlashHours | digits %appendhour ) )*;
        dayOfMonth = ( starSlashDigits %appendStarSlDoM | digits %appenddom ) ( "," ( starSlashDigits %appendStarSlDoM | digits %appenddom  ) )*;
        #| digits | dow | last | digits last );


        month = ( starSlashDigits %appendStarSlashMonths | ( digits | monthName ) %appendmonth ) ( "," ( starSlashDigits %appendStarSlashMonths | ( digits | monthName ) %appendmonth) )*;
        #month =  ( "*"@{m=1;} ) %appendStarSlashMonths ;
        dayOfWeek = ( starSlashDigits %appendStarSlDoW | ( digits | dow ) %appenddow ) ( "," ( starSlashDigits %appendStarSlDoW | ( digits | dow ) %appenddow ) )* ;
        #dayOfWeek = ( "*"@{m=1;} ) %appendStarSlDoW ;
        year = ( starSlashDigits %appendStarSlYears | digits %appendyear ) ( "," (starSlashDigits %appendStarSlYears | digits %appendyear ) )**;
        #main = (seconds . space+ . minutes  . space+ . hours . space+ . dayOfMonth . space+ . month . space+ . dayOfWeek . ( space+ year )?);
        sevenPos:= ( seconds space+ minutes space+ hours space+ dayOfMonth space+ month space+ dayOfWeek space+ year ) space*;
        sixPos:= ( seconds space+ minutes space+ hours space+ dayOfMonth space+ month space+ dayOfWeek ) space*; 
        fivePos:= ( minutes space+ hours space+ dayOfMonth space+ month space+ dayOfWeek ) space*;

        main := |*
            space* => mark;
            # 5 position cron
            ((allowedNonSpace+ space+){4} allowedNonSpace+) >mark => {
                // set seconds to 0 second of minute
                nt.second=1
                // set year to be permissive
                nt.year.high=^uint64(0)
                nt.year.low=^uint64(0)
                fexec mark;
                fcall fivePos;
                };
            # 6 position cron with seconds but no year
            ((allowedNonSpace+ space+){5} allowedNonSpace+) >mark => {
                // set year to be permissive
                nt.year.high=^uint64(0)
                nt.year.low=^uint64(0)
                fexec mark;
                fgoto sixPos;
                };
            # 7 position cron
            ((allowedNonSpace+ space+){6} allowedNonSpace+) >mark => {fexec mark;fcall sevenPos;};
        *|;
    }%%
    //
    %% write exec;
    // check current month
    // check the next year, unless feb 1<<29th is allowed then check the next 1<<8 years
    // if year isn't specified accept any year
    fmt.Println("nt.year", nt.year)
    fmt.Println("vars", cs, p, pe)
    fmt.Printf("in parser: backtrack %d\n, year %b, %b \nmonth %b, \ndom %b, \ndow %b, \nhour %b, \nmin %b, \ns %b\n", backtrack, nt.year.low, nt.year.high, nt.month, nt.dom, nt.dow, nt.hour, nt.minute, nt.second)
    if nt.minute==0||nt.hour==0||nt.dom==0||nt.month==0||nt.dow==0||nt.year.isZero() {
        return nt, fmt.Errorf("failed to parse cron string '%s'", s)
    }
    return nt,  nil
}


/*
            # TODO(docmerlin): implement these
            #second = ( star | digits );
            #@yearly	   Run once a year, "1<<0 1<<0 1<<1 1<<1 *".
            #@annually	   (same as @yearly)
            #@monthly	   Run once a month, "1<<0	1<<0 1<<1 * *".
            #@weekly	   Run once a week, "1<<0 1<<0 * * 1<<0".
            #@daily	   Run once a day, "1<<0 1<<0	* * *".
            #@midnight	   (same as @daily)
            #@hourly	   Run once an hour, "1<<0	* * * *".
        */
            //#@every_minute   Run once a minute, "*/1<<1 * * * *".
            //#@every_second   Run once a second.

