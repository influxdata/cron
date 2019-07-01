
//line parse.rl:1
// -*-go-*-
package cron

import (
    "strconv"
    "fmt"
    "time"
)




//line parse.rl:13

//line parse.rl:14

//line parse.go:20
var _parse_eof_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	38, 40, 42, 44, 
}

const parse_start int = 1
const parse_first_final int = 64
const parse_error int = 0

const parse_en_main int = 1


//line parse.rl:15


func parse(s string, tz *time.Location)(nextTime, error){
    nt:=nextTime{loc:tz}
    cs, p, pe, eof:= 0, 0, len(s), len(s)
    mark := 0
    _ = mark
    m:=uint64(0)
    var err error
    
//line parse.go:51
	{
	cs = parse_start
	}

//line parse.rl:25
    //m,h := 1<<0,1<<0
    
//line parse.rl:375

    
//line parse.go:62
	{
	if p == pe {
		goto _test_eof
	}
	if cs == 0 {
		goto _out
	}
_resume:
	switch cs {
	case 1:
		if ( s)[p] == 42 {
			goto tr0;
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr2;
		}
		goto tr1;
	case 0:
		goto _out
	case 2:
		switch ( s)[p] {
		case 32:
			goto tr3;
		case 44:
			goto tr4;
		case 47:
			goto tr5;
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr3;
		}
		goto tr1;
	case 3:
		switch ( s)[p] {
		case 32:
			goto tr6;
		case 42:
			goto tr7;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr8;
			}
		case ( s)[p] >= 9:
			goto tr6;
		}
		goto tr1;
	case 4:
		switch ( s)[p] {
		case 32:
			goto tr9;
		case 44:
			goto tr10;
		case 47:
			goto tr11;
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr9;
		}
		goto tr1;
	case 5:
		switch ( s)[p] {
		case 32:
			goto tr12;
		case 42:
			goto tr13;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr14;
			}
		case ( s)[p] >= 9:
			goto tr12;
		}
		goto tr1;
	case 6:
		switch ( s)[p] {
		case 32:
			goto tr15;
		case 44:
			goto tr16;
		case 47:
			goto tr17;
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr15;
		}
		goto tr1;
	case 7:
		switch ( s)[p] {
		case 32:
			goto tr18;
		case 42:
			goto tr19;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr20;
			}
		case ( s)[p] >= 9:
			goto tr18;
		}
		goto tr1;
	case 8:
		switch ( s)[p] {
		case 32:
			goto tr21;
		case 44:
			goto tr22;
		case 47:
			goto tr23;
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr21;
		}
		goto tr1;
	case 9:
		switch ( s)[p] {
		case 32:
			goto tr24;
		case 42:
			goto tr25;
		case 65:
			goto tr27;
		case 68:
			goto tr28;
		case 70:
			goto tr29;
		case 74:
			goto tr30;
		case 77:
			goto tr31;
		case 78:
			goto tr32;
		case 79:
			goto tr33;
		case 83:
			goto tr34;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr26;
			}
		case ( s)[p] >= 9:
			goto tr24;
		}
		goto tr1;
	case 10:
		switch ( s)[p] {
		case 32:
			goto tr35;
		case 44:
			goto tr36;
		case 47:
			goto tr37;
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr35;
		}
		goto tr1;
	case 11:
		switch ( s)[p] {
		case 32:
			goto tr38;
		case 42:
			goto tr39;
		case 70:
			goto tr41;
		case 77:
			goto tr42;
		case 83:
			goto tr43;
		case 84:
			goto tr44;
		case 87:
			goto tr45;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr40;
			}
		case ( s)[p] >= 9:
			goto tr38;
		}
		goto tr1;
	case 64:
		switch ( s)[p] {
		case 44:
			goto tr120;
		case 47:
			goto tr121;
		}
		goto tr1;
	case 12:
		if ( s)[p] == 42 {
			goto tr39;
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr40;
		}
		goto tr1;
	case 65:
		if ( s)[p] == 44 {
			goto tr122;
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr123;
		}
		goto tr1;
	case 13:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr46;
		}
		goto tr1;
	case 66:
		if ( s)[p] == 44 {
			goto tr124;
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr125;
		}
		goto tr1;
	case 14:
		if ( s)[p] == 82 {
			goto tr47;
		}
		goto tr1;
	case 15:
		if ( s)[p] == 73 {
			goto tr48;
		}
		goto tr1;
	case 67:
		if ( s)[p] == 44 {
			goto tr126;
		}
		goto tr1;
	case 16:
		if ( s)[p] == 79 {
			goto tr49;
		}
		goto tr1;
	case 17:
		if ( s)[p] == 78 {
			goto tr50;
		}
		goto tr1;
	case 18:
		switch ( s)[p] {
		case 65:
			goto tr51;
		case 85:
			goto tr52;
		}
		goto tr1;
	case 19:
		if ( s)[p] == 84 {
			goto tr53;
		}
		goto tr1;
	case 20:
		if ( s)[p] == 78 {
			goto tr54;
		}
		goto tr1;
	case 21:
		switch ( s)[p] {
		case 72:
			goto tr55;
		case 85:
			goto tr56;
		}
		goto tr1;
	case 22:
		if ( s)[p] == 85 {
			goto tr57;
		}
		goto tr1;
	case 23:
		if ( s)[p] == 69 {
			goto tr58;
		}
		goto tr1;
	case 24:
		if ( s)[p] == 69 {
			goto tr59;
		}
		goto tr1;
	case 25:
		if ( s)[p] == 68 {
			goto tr60;
		}
		goto tr1;
	case 26:
		switch ( s)[p] {
		case 42:
			goto tr25;
		case 65:
			goto tr27;
		case 68:
			goto tr28;
		case 70:
			goto tr29;
		case 74:
			goto tr30;
		case 77:
			goto tr31;
		case 78:
			goto tr32;
		case 79:
			goto tr33;
		case 83:
			goto tr34;
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr26;
		}
		goto tr1;
	case 27:
		switch ( s)[p] {
		case 32:
			goto tr61;
		case 44:
			goto tr62;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr63;
			}
		case ( s)[p] >= 9:
			goto tr61;
		}
		goto tr1;
	case 28:
		switch ( s)[p] {
		case 80:
			goto tr64;
		case 85:
			goto tr65;
		}
		goto tr1;
	case 29:
		if ( s)[p] == 82 {
			goto tr66;
		}
		goto tr1;
	case 30:
		switch ( s)[p] {
		case 32:
			goto tr67;
		case 44:
			goto tr68;
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr67;
		}
		goto tr1;
	case 31:
		if ( s)[p] == 71 {
			goto tr69;
		}
		goto tr1;
	case 32:
		if ( s)[p] == 69 {
			goto tr70;
		}
		goto tr1;
	case 33:
		if ( s)[p] == 67 {
			goto tr71;
		}
		goto tr1;
	case 34:
		if ( s)[p] == 69 {
			goto tr72;
		}
		goto tr1;
	case 35:
		if ( s)[p] == 66 {
			goto tr73;
		}
		goto tr1;
	case 36:
		switch ( s)[p] {
		case 65:
			goto tr74;
		case 85:
			goto tr75;
		}
		goto tr1;
	case 37:
		if ( s)[p] == 78 {
			goto tr76;
		}
		goto tr1;
	case 38:
		switch ( s)[p] {
		case 76:
			goto tr77;
		case 78:
			goto tr78;
		}
		goto tr1;
	case 39:
		if ( s)[p] == 65 {
			goto tr79;
		}
		goto tr1;
	case 40:
		switch ( s)[p] {
		case 82:
			goto tr80;
		case 89:
			goto tr81;
		}
		goto tr1;
	case 41:
		if ( s)[p] == 79 {
			goto tr82;
		}
		goto tr1;
	case 42:
		if ( s)[p] == 86 {
			goto tr83;
		}
		goto tr1;
	case 43:
		if ( s)[p] == 67 {
			goto tr84;
		}
		goto tr1;
	case 44:
		if ( s)[p] == 84 {
			goto tr85;
		}
		goto tr1;
	case 45:
		if ( s)[p] == 69 {
			goto tr86;
		}
		goto tr1;
	case 46:
		if ( s)[p] == 80 {
			goto tr87;
		}
		goto tr1;
	case 47:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr88;
		}
		goto tr1;
	case 48:
		switch ( s)[p] {
		case 32:
			goto tr89;
		case 44:
			goto tr90;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr91;
			}
		case ( s)[p] >= 9:
			goto tr89;
		}
		goto tr1;
	case 49:
		if ( s)[p] == 42 {
			goto tr19;
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr20;
		}
		goto tr1;
	case 50:
		switch ( s)[p] {
		case 32:
			goto tr92;
		case 44:
			goto tr93;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr94;
			}
		case ( s)[p] >= 9:
			goto tr92;
		}
		goto tr1;
	case 51:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr95;
		}
		goto tr1;
	case 52:
		switch ( s)[p] {
		case 32:
			goto tr96;
		case 44:
			goto tr97;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr98;
			}
		case ( s)[p] >= 9:
			goto tr96;
		}
		goto tr1;
	case 53:
		if ( s)[p] == 42 {
			goto tr13;
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr14;
		}
		goto tr1;
	case 54:
		switch ( s)[p] {
		case 32:
			goto tr99;
		case 44:
			goto tr100;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr101;
			}
		case ( s)[p] >= 9:
			goto tr99;
		}
		goto tr1;
	case 55:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr102;
		}
		goto tr1;
	case 56:
		switch ( s)[p] {
		case 32:
			goto tr103;
		case 44:
			goto tr104;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr105;
			}
		case ( s)[p] >= 9:
			goto tr103;
		}
		goto tr1;
	case 57:
		if ( s)[p] == 42 {
			goto tr7;
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr8;
		}
		goto tr1;
	case 58:
		switch ( s)[p] {
		case 32:
			goto tr106;
		case 44:
			goto tr107;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr108;
			}
		case ( s)[p] >= 9:
			goto tr106;
		}
		goto tr1;
	case 59:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr109;
		}
		goto tr1;
	case 60:
		switch ( s)[p] {
		case 32:
			goto tr110;
		case 44:
			goto tr111;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr112;
			}
		case ( s)[p] >= 9:
			goto tr110;
		}
		goto tr1;
	case 61:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr113;
		}
		goto tr1;
	case 62:
		switch ( s)[p] {
		case 32:
			goto tr114;
		case 44:
			goto tr115;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr116;
			}
		case ( s)[p] >= 9:
			goto tr114;
		}
		goto tr1;
	case 63:
		switch ( s)[p] {
		case 32:
			goto tr117;
		case 44:
			goto tr118;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr119;
			}
		case ( s)[p] >= 9:
			goto tr117;
		}
		goto tr1;
	}

	tr1: cs = 0; goto _again
	tr4: cs = 1; goto f2
	tr115: cs = 1; goto f35
	tr118: cs = 1; goto f36
	tr0: cs = 2; goto f0
	tr6: cs = 3; goto _again
	tr3: cs = 3; goto f2
	tr114: cs = 3; goto f35
	tr117: cs = 3; goto f36
	tr7: cs = 4; goto f0
	tr12: cs = 5; goto _again
	tr9: cs = 5; goto f3
	tr106: cs = 5; goto f33
	tr110: cs = 5; goto f34
	tr13: cs = 6; goto f0
	tr18: cs = 7; goto _again
	tr15: cs = 7; goto f4
	tr99: cs = 7; goto f31
	tr103: cs = 7; goto f32
	tr19: cs = 8; goto f0
	tr24: cs = 9; goto _again
	tr21: cs = 9; goto f5
	tr92: cs = 9; goto f29
	tr96: cs = 9; goto f30
	tr25: cs = 10; goto f0
	tr38: cs = 11; goto _again
	tr35: cs = 11; goto f6
	tr61: cs = 11; goto f14
	tr67: cs = 11; goto f16
	tr89: cs = 11; goto f28
	tr120: cs = 12; goto f38
	tr122: cs = 12; goto f40
	tr124: cs = 12; goto f42
	tr126: cs = 12; goto f44
	tr121: cs = 13; goto _again
	tr41: cs = 14; goto _again
	tr47: cs = 15; goto _again
	tr42: cs = 16; goto _again
	tr49: cs = 17; goto _again
	tr43: cs = 18; goto _again
	tr51: cs = 19; goto _again
	tr52: cs = 20; goto _again
	tr44: cs = 21; goto _again
	tr55: cs = 22; goto _again
	tr56: cs = 23; goto _again
	tr45: cs = 24; goto _again
	tr59: cs = 25; goto _again
	tr36: cs = 26; goto f6
	tr62: cs = 26; goto f14
	tr68: cs = 26; goto f16
	tr90: cs = 26; goto f28
	tr63: cs = 27; goto _again
	tr26: cs = 27; goto f1
	tr27: cs = 28; goto _again
	tr64: cs = 29; goto _again
	tr66: cs = 30; goto f15
	tr69: cs = 30; goto f17
	tr71: cs = 30; goto f18
	tr73: cs = 30; goto f19
	tr76: cs = 30; goto f20
	tr77: cs = 30; goto f21
	tr78: cs = 30; goto f22
	tr80: cs = 30; goto f23
	tr81: cs = 30; goto f24
	tr83: cs = 30; goto f25
	tr85: cs = 30; goto f26
	tr87: cs = 30; goto f27
	tr65: cs = 31; goto _again
	tr28: cs = 32; goto _again
	tr70: cs = 33; goto _again
	tr29: cs = 34; goto _again
	tr72: cs = 35; goto _again
	tr30: cs = 36; goto _again
	tr74: cs = 37; goto _again
	tr75: cs = 38; goto _again
	tr31: cs = 39; goto _again
	tr79: cs = 40; goto _again
	tr32: cs = 41; goto _again
	tr82: cs = 42; goto _again
	tr33: cs = 43; goto _again
	tr84: cs = 44; goto _again
	tr34: cs = 45; goto _again
	tr86: cs = 46; goto _again
	tr37: cs = 47; goto _again
	tr91: cs = 48; goto _again
	tr88: cs = 48; goto f1
	tr22: cs = 49; goto f5
	tr93: cs = 49; goto f29
	tr97: cs = 49; goto f30
	tr94: cs = 50; goto _again
	tr20: cs = 50; goto f1
	tr23: cs = 51; goto _again
	tr98: cs = 52; goto _again
	tr95: cs = 52; goto f1
	tr16: cs = 53; goto f4
	tr100: cs = 53; goto f31
	tr104: cs = 53; goto f32
	tr101: cs = 54; goto _again
	tr14: cs = 54; goto f1
	tr17: cs = 55; goto _again
	tr105: cs = 56; goto _again
	tr102: cs = 56; goto f1
	tr10: cs = 57; goto f3
	tr107: cs = 57; goto f33
	tr111: cs = 57; goto f34
	tr108: cs = 58; goto _again
	tr8: cs = 58; goto f1
	tr11: cs = 59; goto _again
	tr112: cs = 60; goto _again
	tr109: cs = 60; goto f1
	tr5: cs = 61; goto _again
	tr116: cs = 62; goto _again
	tr113: cs = 62; goto f1
	tr119: cs = 63; goto _again
	tr2: cs = 63; goto f1
	tr39: cs = 64; goto f0
	tr123: cs = 65; goto _again
	tr40: cs = 65; goto f1
	tr125: cs = 66; goto _again
	tr46: cs = 66; goto f1
	tr48: cs = 67; goto f7
	tr50: cs = 67; goto f8
	tr53: cs = 67; goto f9
	tr54: cs = 67; goto f10
	tr57: cs = 67; goto f11
	tr58: cs = 67; goto f12
	tr60: cs = 67; goto f13

f1:
//line parse.rl:27

            mark = p;
        
	goto _again
f16:
//line parse.rl:42

            if m>12 || m<1{
                return nt, fmt.Errorf("invalid month %d", m)
            }
            nt.month |= 1<<(m-1)
        
	goto _again
f44:
//line parse.rl:55

            if m>1<<6 {
                return nt, fmt.Errorf("invalid day of week %d", m)
            }
            nt.dow |= 1<<m
        
	goto _again
f2:
//line parse.rl:76

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
        
	goto _again
f3:
//line parse.rl:147

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
        
	goto _again
f4:
//line parse.rl:219

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
        
	goto _again
f6:
//line parse.rl:254

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
        
	goto _again
f38:
//line parse.rl:277

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
        
	goto _again
f5:
//line parse.rl:295

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
        
	goto _again
f10:
//line parse.rl:352
m=1
	goto _again
f8:
//line parse.rl:352
m=1<<1
	goto _again
f12:
//line parse.rl:352
m=1<<2
	goto _again
f13:
//line parse.rl:352
m=1<<3
	goto _again
f11:
//line parse.rl:352
m=1<<4
	goto _again
f7:
//line parse.rl:352
m=1<<5
	goto _again
f9:
//line parse.rl:352
m=1<<6
	goto _again
f20:
//line parse.rl:353
m=1
	goto _again
f19:
//line parse.rl:353
m=1<<1
	goto _again
f23:
//line parse.rl:353
m=1<<2
	goto _again
f15:
//line parse.rl:353
m=1<<3
	goto _again
f24:
//line parse.rl:353
m=1<<4
	goto _again
f22:
//line parse.rl:353
m=1<<5
	goto _again
f21:
//line parse.rl:353
m=1<<6
	goto _again
f17:
//line parse.rl:353
m=1<<7
	goto _again
f27:
//line parse.rl:353
m=1<<8
	goto _again
f26:
//line parse.rl:353
m=1<<9
	goto _again
f25:
//line parse.rl:353
m=1<<10
	goto _again
f18:
//line parse.rl:353
m=1<<11
	goto _again
f0:
//line parse.rl:363
m=1;
	goto _again
f36:
//line parse.rl:340

            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil{
                return nt,  err
            }
        
//line parse.rl:35

            if m>1<<60 {
                return nt, fmt.Errorf("invalid minute %d", m)
            }
            nt.second |= 1<<m
        
	goto _again
f14:
//line parse.rl:340

            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil{
                return nt,  err
            }
        
//line parse.rl:42

            if m>12 || m<1{
                return nt, fmt.Errorf("invalid month %d", m)
            }
            nt.month |= 1<<(m-1)
        
	goto _again
f33:
//line parse.rl:340

            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil{
                return nt,  err
            }
        
//line parse.rl:48

            if m>1<<60 {
                return nt, fmt.Errorf("invalid minute %d", m)
            }
            nt.minute |= 1<<m
        
	goto _again
f40:
//line parse.rl:340

            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil{
                return nt,  err
            }
        
//line parse.rl:55

            if m>1<<6 {
                return nt, fmt.Errorf("invalid day of week %d", m)
            }
            nt.dow |= 1<<m
        
	goto _again
f31:
//line parse.rl:340

            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil{
                return nt,  err
            }
        
//line parse.rl:62

            if m>24 || m <0 {
                return nt, fmt.Errorf("invalid day of hour %d", m)
            }
            nt.hour |= 1<<m
        
	goto _again
f29:
//line parse.rl:340

            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil{
                return nt,  err
            }
        
//line parse.rl:69

            if m>24 || m <1 {
                return nt,  fmt.Errorf("invalid day of hour %d", m)
            }
            nt.dom |= 1<<(m-1)
        
	goto _again
f35:
//line parse.rl:340

            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil{
                return nt,  err
            }
        
//line parse.rl:76

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
        
	goto _again
f34:
//line parse.rl:340

            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil{
                return nt,  err
            }
        
//line parse.rl:147

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
        
	goto _again
f32:
//line parse.rl:340

            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil{
                return nt,  err
            }
        
//line parse.rl:219

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
        
	goto _again
f28:
//line parse.rl:340

            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil{
                return nt,  err
            }
        
//line parse.rl:254

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
        
	goto _again
f42:
//line parse.rl:340

            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil{
                return nt,  err
            }
        
//line parse.rl:277

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
        
	goto _again
f30:
//line parse.rl:340

            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil{
                return nt,  err
            }
        
//line parse.rl:295

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
        
	goto _again

_again:
	if cs == 0 {
		goto _out
	}
	if p++; p != pe {
		goto _resume
	}
	_test_eof: {}
	if p == eof {
		switch _parse_eof_actions[cs] {
		case 44:
//line parse.rl:55

            if m>1<<6 {
                return nt, fmt.Errorf("invalid day of week %d", m)
            }
            nt.dow |= 1<<m
        
//line parse.rl:30
 
            fmt.Println(nt)
            //  fmt.Printf("\nhere%d cs: %d, p:%d, pe: %d, eof:%d, mark:%d sl:%s\n",  1<<0, cs, p, pe, eof, mark, data[mark:p])
        
		case 38:
//line parse.rl:277

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
        
//line parse.rl:30
 
            fmt.Println(nt)
            //  fmt.Printf("\nhere%d cs: %d, p:%d, pe: %d, eof:%d, mark:%d sl:%s\n",  1<<0, cs, p, pe, eof, mark, data[mark:p])
        
		case 40:
//line parse.rl:340

            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil{
                return nt,  err
            }
        
//line parse.rl:55

            if m>1<<6 {
                return nt, fmt.Errorf("invalid day of week %d", m)
            }
            nt.dow |= 1<<m
        
//line parse.rl:30
 
            fmt.Println(nt)
            //  fmt.Printf("\nhere%d cs: %d, p:%d, pe: %d, eof:%d, mark:%d sl:%s\n",  1<<0, cs, p, pe, eof, mark, data[mark:p])
        
		case 42:
//line parse.rl:340

            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil{
                return nt,  err
            }
        
//line parse.rl:277

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
        
//line parse.rl:30
 
            fmt.Println(nt)
            //  fmt.Printf("\nhere%d cs: %d, p:%d, pe: %d, eof:%d, mark:%d sl:%s\n",  1<<0, cs, p, pe, eof, mark, data[mark:p])
        
//line parse.go:1725
		}
	}

	_out: {}
	}

//line parse.rl:377
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

