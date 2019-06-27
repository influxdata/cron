// -*-go-*-
package cron

import (
    "strconv"
    "fmt"
    "time"
)



%% machine parse;
%% variable data s;
%% write data;


func parse(s string, tz *time.Location)(nextTime, error){
    nt:=nextTime{}
    cs, p, pe, eof:= 0, 0, len(s), len(s)
    mark := 0
    _ = mark
    m:=uint64(0)
    var err error
    %% write init;
    //m,h := 1<<0,1<<0
    %%{
            action mark {
            mark = p;
        }
        action printf { 
            fmt.Println(nt)
            //  fmt.Printf("\nhere%d cs: %d, p:%d, pe: %d, eof:%d, mark:%d sl:%s\n",  1<<0, cs, p, pe, eof, mark, data[mark:p])
        }

        action appendsecond {
            if m>1<<60 {
                return nt, fmt.Errorf("invalid minute %d", m)
            }
            nt.second |= 1<<m
        }

        action appendmonth {
            if m>12 || m<1{
                return nt, fmt.Errorf("invalid month %d", m)
            }
            nt.month |= 1<<(m-1)
        }
        action appendminute {
            if m>1<<60 {
                return nt, fmt.Errorf("invalid minute %d", m)
            }
            nt.minute |= 1<<m
        }

        action appenddow {
            if m>1<<6 {
                return nt, fmt.Errorf("invalid day of week %d", m)
            }
            nt.dow |= 1<<m
        }

        action appendhour {
            if m>24 || m <0 {
                return nt, fmt.Errorf("invalid day of hour %d", m)
            }
            nt.hour |= 1<<m
        }

        action appenddom {
            if m>24 || m <1 {
                return nt,  fmt.Errorf("invalid day of hour %d", m)
            }
            nt.dom |= 1<<(m-1)
        }

        action appendStarSlashSeconds {
            {
                x:= [...]uint64{
                    1|1<<1|1<<2|1<<3|1<<4|1<<5|1<<6|1<<7|1<<8|1<<9|1<<10|1<<11|1<<12|1<<13|1<<14|1<<15|1<<16|1<<17|1<<18|1<<19|1<<20|1<<21|1<<22|1<<23|1<<24|1<<25|1<<26|1<<27|1<<28|1<<29|1<<30|1<<31|1<<32|1<<33|1<<34|1<<35|1<<36|1<<37|1<<38|1<<39|1<<40|1<<41|1<<42|1<<43|1<<44|1<<45|1<<46|1<<47|1<<48|1<<49|1<<50|1<<51|1<<52|1<<53|1<<54|1<<55|1<<56|1<<57|1<<58|1<<59,
                    1<<1|1<<3|1<<5|1<<7|1<<9|1<<11|1<<13|1<<15|1<<17|1<<19|1<<21|1<<23|1<<25|1<<27|1<<29|1<<31|1<<33|1<<35|1<<37|1<<39|1<<41|1<<43|1<<45|1<<47|1<<49|1<<51|1<<53|1<<55|1<<57|1<<59,
                    1<<2|1<<5|1<<8|1<<11|1<<14|1<<17|1<<20|1<<23|1<<26|1<<29|1<<32|1<<35|1<<38|1<<41|1<<44|1<<47|1<<50|1<<53|1<<56|1<<59,
                    1<<3|1<<7|1<<11|1<<15|1<<19|1<<23|1<<27|1<<31|1<<35|1<<39|1<<43|1<<47|1<<51|1<<55|1<<59,
                    1<<4|1<<9|1<<14|1<<19|1<<24|1<<29|1<<34|1<<39|1<<44|1<<49|1<<54|1<<59,
                    1<<5|1<<11|1<<17|1<<23|1<<29|1<<35|1<<41|1<<47|1<<53|1<<59,
                    1<<6|1<<13|1<<20|1<<27|1<<34|1<<41|1<<48|1<<55,
                    1<<7|1<<15|1<<23|1<<31|1<<39|1<<47|1<<55,
                    1<<8|1<<17|1<<26|1<<35|1<<44|1<<53,
                    1<<9|1<<19|1<<29|1<<39|1<<49|1<<59,
                    1<<10|1<<21|1<<32|1<<43|1<<54,
                    1<<11|1<<23|1<<35|1<<47|1<<59,
                    1<<12|1<<25|1<<38|1<<51,
                    1<<13|1<<27|1<<41|1<<55,
                    1<<14|1<<29|1<<44|1<<59,
                    1<<15|1<<31|1<<47,
                    1<<16|1<<33|1<<50,
                    1<<17|1<<35|1<<53,
                    1<<18|1<<37|1<<56,
                    1<<19|1<<39|1<<59,
                    1<<20|1<<41,
                    1<<21|1<<43,
                    1<<22|1<<45,
                    1<<23|1<<47,
                    1<<24|1<<49,
                    1<<25|1<<51,
                    1<<26|1<<53,
                    1<<27|1<<55,
                    1<<28|1<<57,
                    1<<29|1<<59,
                    1<<30,
                    1<<31,
                    1<<32,
                    1<<33,
                    1<<34,
                    1<<35,
                    1<<36,
                    1<<37,
                    1<<38,
                    1<<39,
                    1<<40,
                    1<<41,
                    1<<42,
                    1<<43,
                    1<<44,
                    1<<45,
                    1<<46,
                    1<<47,
                    1<<48,
                    1<<49,
                    1<<50,
                    1<<51,
                    1<<52,
                    1<<53,
                    1<<54,
                    1<<55,
                    1<<56,
                    1<<57,
                    1<<58,
                    1<<59,
                }
                if m<1 || m>=60 {
                    return nt, fmt.Errorf("invalid month */%d", m)
                }
                nt.second |= x[m-1]
            }
        }

        action appendStarSlashMinutes {
            {
                // I really wish that golang had a binary bit representation, this is pretty aweful, but easier to understand than using hex or octal I guess
                x:= [...]uint64{
                    1|1<<1|1<<2|1<<3|1<<4|1<<5|1<<6|1<<7|1<<8|1<<9|1<<10|1<<11|1<<12|1<<13|1<<14|1<<15|1<<16|1<<17|1<<18|1<<19|1<<20|1<<21|1<<22|1<<23|1<<24|1<<25|1<<26|1<<27|1<<28|1<<29|1<<30|1<<31|1<<32|1<<33|1<<34|1<<35|1<<36|1<<37|1<<38|1<<39|1<<40|1<<41|1<<42|1<<43|1<<44|1<<45|1<<46|1<<47|1<<48|1<<49|1<<50|1<<51|1<<52|1<<53|1<<54|1<<55|1<<56|1<<57|1<<58|1<<59,
                    1<<1|1<<3|1<<5|1<<7|1<<9|1<<11|1<<13|1<<15|1<<17|1<<19|1<<21|1<<23|1<<25|1<<27|1<<29|1<<31|1<<33|1<<35|1<<37|1<<39|1<<41|1<<43|1<<45|1<<47|1<<49|1<<51|1<<53|1<<55|1<<57|1<<59,
                    1<<2|1<<5|1<<8|1<<11|1<<14|1<<17|1<<20|1<<23|1<<26|1<<29|1<<32|1<<35|1<<38|1<<41|1<<44|1<<47|1<<50|1<<53|1<<56|1<<59,
                    1<<3|1<<7|1<<11|1<<15|1<<19|1<<23|1<<27|1<<31|1<<35|1<<39|1<<43|1<<47|1<<51|1<<55|1<<59,
                    1<<4|1<<9|1<<14|1<<19|1<<24|1<<29|1<<34|1<<39|1<<44|1<<49|1<<54|1<<59,
                    1<<5|1<<11|1<<17|1<<23|1<<29|1<<35|1<<41|1<<47|1<<53|1<<59,
                    1<<6|1<<13|1<<20|1<<27|1<<34|1<<41|1<<48|1<<55,
                    1<<7|1<<15|1<<23|1<<31|1<<39|1<<47|1<<55,
                    1<<8|1<<17|1<<26|1<<35|1<<44|1<<53,
                    1<<9|1<<19|1<<29|1<<39|1<<49|1<<59,
                    1<<10|1<<21|1<<32|1<<43|1<<54,
                    1<<11|1<<23|1<<35|1<<47|1<<59,
                    1<<12|1<<25|1<<38|1<<51,
                    1<<13|1<<27|1<<41|1<<55,
                    1<<14|1<<29|1<<44|1<<59,
                    1<<15|1<<31|1<<47,
                    1<<16|1<<33|1<<50,
                    1<<17|1<<35|1<<53,
                    1<<18|1<<37|1<<56,
                    1<<19|1<<39|1<<59,
                    1<<20|1<<41,
                    1<<21|1<<43,
                    1<<22|1<<45,
                    1<<23|1<<47,
                    1<<24|1<<49,
                    1<<25|1<<51,
                    1<<26|1<<53,
                    1<<27|1<<55,
                    1<<28|1<<57,
                    1<<29|1<<59,
                    1<<30,
                    1<<31,
                    1<<32,
                    1<<33,
                    1<<34,
                    1<<35,
                    1<<36,
                    1<<37,
                    1<<38,
                    1<<39,
                    1<<40,
                    1<<41,
                    1<<42,
                    1<<43,
                    1<<44,
                    1<<45,
                    1<<46,
                    1<<47,
                    1<<48,
                    1<<49,
                    1<<50,
                    1<<51,
                    1<<52,
                    1<<53,
                    1<<54,
                    1<<55,
                    1<<56,
                    1<<57,
                    1<<58,
                    1<<59,
                }
                if m<1 || m>=60 {
                    return nt, fmt.Errorf("invalid month */%d", m)
                }
                nt.minute |= x[m-1]
            }
        }

        action appendStarSlashHours {
            {
                x:= [...]uint64{
                    1|1<<1|1<<2|1<<3|1<<4|1<<5|1<<6|1<<7|1<<8|1<<9|1<<10|1<<11|1<<12|1<<13|1<<14|1<<15|1<<16|1<<17|1<<18|1<<19|1<<20|1<<21|1<<22|1<<23,
                    1<<1|1<<3|1<<5|1<<7|1<<9|1<<11|1<<13|1<<15|1<<17|1<<19|1<<21|1<<23,
                    1<<2|1<<5|1<<8|1<<11|1<<14|1<<17|1<<20|1<<23,
                    1<<3|1<<7|1<<11|1<<15|1<<19|1<<23,
                    1<<4|1<<9|1<<14|1<<19,
                    1<<5|1<<11|1<<17|1<<23,
                    1<<6|1<<13|1<<20,
                    1<<7|1<<15|1<<23,
                    1<<8|1<<17,
                    1<<9|1<<19,
                    1<<10|1<<21,
                    1<<11|1<<23,
                    1<<12,
                    1<<13,
                    1<<14,
                    1<<15,
                    1<<16,
                    1<<17,
                    1<<18,
                    1<<19,
                    1<<20,
                    1<<21,
                    1<<22,
                    1<<23,
                }
                if m<0 || m>=24 {
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

        action appendStarSlashDoW {
            {
                x:= [...]uint64{
                    1|1<<1|1<<2|1<<3|1<<4|1<<5|1<<6,
                    1<<1|1<<3|1<<5,
                    1<<2|1<<5,
                    1<<3,
                    1<<4,
                    1<<5,
                    1<<6,
                }
                if m<1 || m>=7 {
                    return nt,  fmt.Errorf("invalid day of week */%d", m)
                }
                nt.dow |= x[m-1]
            }
        }

        action appendStarSlashDoM {
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
                if m<1 || m>31 {
                    return nt,  fmt.Errorf("invalid day of month */%d", m)
                }
                nt.dom |= x[m-1]
            }
        }

        #space = ( " " | "\t" );
        spaces = space+;
        #digit = [1<<0-1<<9];
        digits = digit+ >mark %{
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil{
                return nt,  err
            }
        };
        star = "*";
        slash = "/";
        comma = ",";
        hypen = "-";
        hash = "#";

        dow = ( "SUN" @{m=1} | "MON" @{m=1<<1} | "TUE" @{m=1<<2} | "WED" @{m=1<<3} | "THU" @{m=1<<4} | "FRI" @{m=1<<5} | "SAT" @{m=1<<6} );
        monthName = ( "JAN" @{m=1<<1} | "FEB" @{m=1<<2} | "MAR" @{m=1<<3} | "APR" @{m=1<<4} | "MAY" @{m=1<<5} | "JUN" @{m=1<<6} | "JUL" @{m=1<<7} | "AUG" @{m=1<<8} | "SEP" @{m=1<<9} | "OCT" @{m=1<<10} | "NOV" @{m=1<<11} | "DEC" @{m=1<<12} ) ;
        last = ( digit+ "L" ) >mark %{
            m, err = strconv.ParseUint(s[mark:p-1], 10, 64)
            if err!=nil{
                return nt,  err
            }
        };
        year = digits;

        digitlist = digits (comma space* digits)*;
        starSlashDigits = ( star @{m=1;}( slash %mark digits )? );
        digitsAndSlashList = ( starSlashDigits | digits ) ( comma ( starSlashDigits | digits ) )*;

        seconds =  ( starSlashDigits %appendStarSlashSeconds | digits %appendsecond ) ( comma ( starSlashDigits %appendStarSlashSeconds | digits %appendsecond ) )*; 
        minutes =  ( starSlashDigits %appendStarSlashMinutes | digits %appendminute ) ( comma ( starSlashDigits %appendStarSlashMinutes | digits %appendminute ) )*; 
        hours = ( starSlashDigits %appendStarSlashHours | digits %appendhour ) ( comma ( starSlashDigits %appendStarSlashHours | digits %appendhour ) )*;
        dayOfMonth = ( starSlashDigits %appendStarSlashDoM | digits %appenddom ) ( comma ( starSlashDigits %appendStarSlashDoM | digits %appenddom  ) )*;
        #| digits | dow | last | digits last );

        month = ( starSlashDigits %appendStarSlashMonths | ( digits | monthName ) %appendmonth ) ( comma ( starSlashDigits %appendStarSlashMonths | ( digits | monthName ) %appendmonth) )* ;
        dayOfWeek = ( starSlashDigits %appendStarSlashDoW | digits %appenddow | dow %appenddow ) ( comma ( starSlashDigits %appendStarSlashDoW | digits %appenddow )  )*;
        main := ( seconds spaces minutes spaces hours spaces dayOfMonth spaces month spaces dayOfWeek ) %printf;
    }%%
    %% write exec;
    if p==0{
        return nt,  nil
    }
    // check current month
    // check the next year, unless feb 1<<29th is allowed then check the next 1<<8 years
    //fmt.Println(ts)
    return nt,  nil //errors.New("invalid cron")//time.Date(resYear, time.Month(resMonth), resDay, 1<<0,1<<0,1<<0,1<<0, tz), nil
    //panic("")
}

// for month := range m.month {
//     if tmonth>tsmonth{

//     }
//     tsyear += 
// }

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

