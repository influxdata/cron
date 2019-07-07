
//line parse.rl:1
// -*-go-*-
package cron

import (
    "strconv"
    "fmt"
    "time"
    "errors"
)



//line parse.rl:13

//line parse.rl:14

//line parse.go:20
var _parse_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 8, 
	1, 9, 1, 10, 1, 11, 1, 12, 
	1, 14, 1, 15, 1, 16, 1, 17, 
	1, 18, 1, 19, 1, 20, 1, 21, 
	1, 22, 1, 23, 1, 24, 1, 25, 
	1, 26, 1, 27, 1, 28, 1, 29, 
	1, 30, 1, 31, 1, 32, 1, 33, 
	1, 34, 1, 37, 2, 13, 1, 2, 
	13, 2, 2, 13, 3, 2, 13, 4, 
	2, 13, 5, 2, 13, 8, 2, 13, 
	9, 2, 13, 10, 2, 13, 11, 2, 
	13, 12, 2, 13, 35, 2, 13, 36, 
	2, 34, 6, 2, 34, 7, 2, 37, 
	0, 3, 13, 35, 6, 3, 13, 35, 
	7, 3, 13, 36, 6, 3, 13, 36, 
	7, 3, 13, 38, 6, 3, 13, 39, 
	7, 
}

var _parse_eof_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	11, 13, 73, 85, 64, 5, 82, 
}

const parse_start int = 1
const parse_first_final int = 72
const parse_error int = 0

const parse_en_main int = 1


//line parse.rl:15

func parse(s string, tz *time.Location)(nextTime, error){
    nt:=nextTime{loc:tz}
    cs, p, pe, eof:= 0, 0,len(s), len(s)

    mark := 0
    _ = mark
    m, d, start, end:=uint64(0), uint64(0), uint64(0), uint64(0)
    _ = d
    // TODO(docmerlin): handle ranges
    var err error
    
//line parse.go:74
	{
	cs = parse_start
	}

//line parse.rl:27
    //m,h := 1<<0,1<<0
    
//line parse.rl:448

    
//line parse.go:85
	{
	var _acts int
	var _nacts uint

	if p == pe {
		goto _test_eof
	}
	if cs == 0 {
		goto _out
	}
_resume:
	switch cs {
	case 1:
		switch ( s)[p] {
		case 32:
			goto tr0;
		case 42:
			goto tr2;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr3;
			}
		case ( s)[p] >= 9:
			goto tr0;
		}
		goto tr1;
	case 0:
		goto _out
	case 2:
		switch ( s)[p] {
		case 32:
			goto tr4;
		case 44:
			goto tr5;
		case 47:
			goto tr6;
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr4;
		}
		goto tr1;
	case 3:
		switch ( s)[p] {
		case 32:
			goto tr7;
		case 42:
			goto tr8;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr9;
			}
		case ( s)[p] >= 9:
			goto tr7;
		}
		goto tr1;
	case 4:
		switch ( s)[p] {
		case 32:
			goto tr10;
		case 44:
			goto tr11;
		case 47:
			goto tr12;
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr10;
		}
		goto tr1;
	case 5:
		switch ( s)[p] {
		case 32:
			goto tr13;
		case 42:
			goto tr14;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr15;
			}
		case ( s)[p] >= 9:
			goto tr13;
		}
		goto tr1;
	case 6:
		switch ( s)[p] {
		case 32:
			goto tr16;
		case 44:
			goto tr17;
		case 47:
			goto tr18;
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr16;
		}
		goto tr1;
	case 7:
		switch ( s)[p] {
		case 32:
			goto tr19;
		case 42:
			goto tr20;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr21;
			}
		case ( s)[p] >= 9:
			goto tr19;
		}
		goto tr1;
	case 8:
		switch ( s)[p] {
		case 32:
			goto tr22;
		case 44:
			goto tr23;
		case 47:
			goto tr24;
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr22;
		}
		goto tr1;
	case 9:
		switch ( s)[p] {
		case 32:
			goto tr25;
		case 42:
			goto tr26;
		case 65:
			goto tr28;
		case 68:
			goto tr29;
		case 70:
			goto tr30;
		case 74:
			goto tr31;
		case 77:
			goto tr32;
		case 78:
			goto tr33;
		case 79:
			goto tr34;
		case 83:
			goto tr35;
		case 97:
			goto tr28;
		case 100:
			goto tr29;
		case 102:
			goto tr30;
		case 106:
			goto tr31;
		case 109:
			goto tr32;
		case 110:
			goto tr33;
		case 111:
			goto tr34;
		case 115:
			goto tr35;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr27;
			}
		case ( s)[p] >= 9:
			goto tr25;
		}
		goto tr1;
	case 10:
		switch ( s)[p] {
		case 32:
			goto tr36;
		case 44:
			goto tr37;
		case 47:
			goto tr38;
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr36;
		}
		goto tr1;
	case 11:
		switch ( s)[p] {
		case 32:
			goto tr39;
		case 42:
			goto tr40;
		case 70:
			goto tr42;
		case 77:
			goto tr43;
		case 83:
			goto tr44;
		case 84:
			goto tr45;
		case 87:
			goto tr46;
		case 102:
			goto tr42;
		case 109:
			goto tr43;
		case 115:
			goto tr44;
		case 116:
			goto tr45;
		case 119:
			goto tr46;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr41;
			}
		case ( s)[p] >= 9:
			goto tr39;
		}
		goto tr1;
	case 72:
		switch ( s)[p] {
		case 32:
			goto tr139;
		case 44:
			goto tr140;
		case 47:
			goto tr141;
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr139;
		}
		goto tr1;
	case 12:
		switch ( s)[p] {
		case 32:
			goto tr47;
		case 42:
			goto tr48;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr49;
			}
		case ( s)[p] >= 9:
			goto tr47;
		}
		goto tr1;
	case 73:
		switch ( s)[p] {
		case 44:
			goto tr142;
		case 47:
			goto tr143;
		}
		goto tr1;
	case 13:
		if ( s)[p] == 42 {
			goto tr48;
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr49;
		}
		goto tr1;
	case 74:
		if ( s)[p] == 44 {
			goto tr144;
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr145;
		}
		goto tr1;
	case 14:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr50;
		}
		goto tr1;
	case 75:
		if ( s)[p] == 44 {
			goto tr146;
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr147;
		}
		goto tr1;
	case 15:
		switch ( s)[p] {
		case 42:
			goto tr40;
		case 70:
			goto tr42;
		case 77:
			goto tr43;
		case 83:
			goto tr44;
		case 84:
			goto tr45;
		case 87:
			goto tr46;
		case 102:
			goto tr42;
		case 109:
			goto tr43;
		case 115:
			goto tr44;
		case 116:
			goto tr45;
		case 119:
			goto tr46;
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr41;
		}
		goto tr1;
	case 76:
		switch ( s)[p] {
		case 32:
			goto tr148;
		case 44:
			goto tr149;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr150;
			}
		case ( s)[p] >= 9:
			goto tr148;
		}
		goto tr1;
	case 16:
		switch ( s)[p] {
		case 82:
			goto tr51;
		case 114:
			goto tr51;
		}
		goto tr1;
	case 17:
		switch ( s)[p] {
		case 73:
			goto tr52;
		case 105:
			goto tr52;
		}
		goto tr1;
	case 77:
		switch ( s)[p] {
		case 32:
			goto tr151;
		case 44:
			goto tr152;
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr151;
		}
		goto tr1;
	case 18:
		switch ( s)[p] {
		case 79:
			goto tr53;
		case 111:
			goto tr53;
		}
		goto tr1;
	case 19:
		switch ( s)[p] {
		case 78:
			goto tr54;
		case 110:
			goto tr54;
		}
		goto tr1;
	case 20:
		switch ( s)[p] {
		case 65:
			goto tr55;
		case 85:
			goto tr56;
		case 97:
			goto tr55;
		case 117:
			goto tr56;
		}
		goto tr1;
	case 21:
		switch ( s)[p] {
		case 84:
			goto tr57;
		case 116:
			goto tr57;
		}
		goto tr1;
	case 22:
		switch ( s)[p] {
		case 78:
			goto tr58;
		case 110:
			goto tr58;
		}
		goto tr1;
	case 23:
		switch ( s)[p] {
		case 72:
			goto tr59;
		case 85:
			goto tr60;
		case 104:
			goto tr59;
		case 117:
			goto tr60;
		}
		goto tr1;
	case 24:
		switch ( s)[p] {
		case 85:
			goto tr61;
		case 117:
			goto tr61;
		}
		goto tr1;
	case 25:
		switch ( s)[p] {
		case 69:
			goto tr62;
		case 101:
			goto tr62;
		}
		goto tr1;
	case 26:
		switch ( s)[p] {
		case 69:
			goto tr63;
		case 101:
			goto tr63;
		}
		goto tr1;
	case 27:
		switch ( s)[p] {
		case 68:
			goto tr64;
		case 100:
			goto tr64;
		}
		goto tr1;
	case 28:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr65;
		}
		goto tr1;
	case 78:
		switch ( s)[p] {
		case 32:
			goto tr153;
		case 44:
			goto tr154;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr155;
			}
		case ( s)[p] >= 9:
			goto tr153;
		}
		goto tr1;
	case 29:
		switch ( s)[p] {
		case 42:
			goto tr26;
		case 65:
			goto tr28;
		case 68:
			goto tr29;
		case 70:
			goto tr30;
		case 74:
			goto tr31;
		case 77:
			goto tr32;
		case 78:
			goto tr33;
		case 79:
			goto tr34;
		case 83:
			goto tr35;
		case 97:
			goto tr28;
		case 100:
			goto tr29;
		case 102:
			goto tr30;
		case 106:
			goto tr31;
		case 109:
			goto tr32;
		case 110:
			goto tr33;
		case 111:
			goto tr34;
		case 115:
			goto tr35;
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr27;
		}
		goto tr1;
	case 30:
		switch ( s)[p] {
		case 32:
			goto tr66;
		case 44:
			goto tr67;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr68;
			}
		case ( s)[p] >= 9:
			goto tr66;
		}
		goto tr1;
	case 31:
		switch ( s)[p] {
		case 80:
			goto tr69;
		case 85:
			goto tr70;
		case 112:
			goto tr69;
		case 117:
			goto tr70;
		}
		goto tr1;
	case 32:
		switch ( s)[p] {
		case 82:
			goto tr71;
		case 114:
			goto tr71;
		}
		goto tr1;
	case 33:
		switch ( s)[p] {
		case 32:
			goto tr72;
		case 44:
			goto tr73;
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr72;
		}
		goto tr1;
	case 34:
		switch ( s)[p] {
		case 71:
			goto tr74;
		case 103:
			goto tr74;
		}
		goto tr1;
	case 35:
		switch ( s)[p] {
		case 69:
			goto tr75;
		case 101:
			goto tr75;
		}
		goto tr1;
	case 36:
		switch ( s)[p] {
		case 67:
			goto tr76;
		case 99:
			goto tr76;
		}
		goto tr1;
	case 37:
		switch ( s)[p] {
		case 69:
			goto tr77;
		case 101:
			goto tr77;
		}
		goto tr1;
	case 38:
		switch ( s)[p] {
		case 66:
			goto tr78;
		case 98:
			goto tr78;
		}
		goto tr1;
	case 39:
		switch ( s)[p] {
		case 65:
			goto tr79;
		case 85:
			goto tr80;
		case 97:
			goto tr79;
		case 117:
			goto tr80;
		}
		goto tr1;
	case 40:
		switch ( s)[p] {
		case 78:
			goto tr81;
		case 110:
			goto tr81;
		}
		goto tr1;
	case 41:
		switch ( s)[p] {
		case 76:
			goto tr82;
		case 78:
			goto tr83;
		case 108:
			goto tr82;
		case 110:
			goto tr83;
		}
		goto tr1;
	case 42:
		switch ( s)[p] {
		case 65:
			goto tr84;
		case 97:
			goto tr84;
		}
		goto tr1;
	case 43:
		switch ( s)[p] {
		case 82:
			goto tr85;
		case 89:
			goto tr86;
		case 114:
			goto tr85;
		case 121:
			goto tr86;
		}
		goto tr1;
	case 44:
		switch ( s)[p] {
		case 79:
			goto tr87;
		case 111:
			goto tr87;
		}
		goto tr1;
	case 45:
		switch ( s)[p] {
		case 86:
			goto tr88;
		case 118:
			goto tr88;
		}
		goto tr1;
	case 46:
		switch ( s)[p] {
		case 67:
			goto tr89;
		case 99:
			goto tr89;
		}
		goto tr1;
	case 47:
		switch ( s)[p] {
		case 84:
			goto tr90;
		case 116:
			goto tr90;
		}
		goto tr1;
	case 48:
		switch ( s)[p] {
		case 69:
			goto tr91;
		case 101:
			goto tr91;
		}
		goto tr1;
	case 49:
		switch ( s)[p] {
		case 80:
			goto tr92;
		case 112:
			goto tr92;
		}
		goto tr1;
	case 50:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr93;
		}
		goto tr1;
	case 51:
		switch ( s)[p] {
		case 32:
			goto tr94;
		case 44:
			goto tr95;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr96;
			}
		case ( s)[p] >= 9:
			goto tr94;
		}
		goto tr1;
	case 52:
		if ( s)[p] == 42 {
			goto tr20;
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr21;
		}
		goto tr1;
	case 53:
		switch ( s)[p] {
		case 32:
			goto tr97;
		case 44:
			goto tr98;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr99;
			}
		case ( s)[p] >= 9:
			goto tr97;
		}
		goto tr1;
	case 54:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr100;
		}
		goto tr1;
	case 55:
		switch ( s)[p] {
		case 32:
			goto tr101;
		case 44:
			goto tr102;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr103;
			}
		case ( s)[p] >= 9:
			goto tr101;
		}
		goto tr1;
	case 56:
		if ( s)[p] == 42 {
			goto tr14;
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr15;
		}
		goto tr1;
	case 57:
		switch ( s)[p] {
		case 32:
			goto tr104;
		case 44:
			goto tr105;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr106;
			}
		case ( s)[p] >= 9:
			goto tr104;
		}
		goto tr1;
	case 58:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr107;
		}
		goto tr1;
	case 59:
		switch ( s)[p] {
		case 32:
			goto tr108;
		case 44:
			goto tr109;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr110;
			}
		case ( s)[p] >= 9:
			goto tr108;
		}
		goto tr1;
	case 60:
		if ( s)[p] == 42 {
			goto tr8;
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr9;
		}
		goto tr1;
	case 61:
		switch ( s)[p] {
		case 32:
			goto tr111;
		case 44:
			goto tr112;
		case 45:
			goto tr113;
		case 47:
			goto tr114;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr115;
			}
		case ( s)[p] >= 9:
			goto tr111;
		}
		goto tr1;
	case 62:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr116;
		}
		goto tr1;
	case 63:
		switch ( s)[p] {
		case 32:
			goto tr117;
		case 44:
			goto tr118;
		case 47:
			goto tr119;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr120;
			}
		case ( s)[p] >= 9:
			goto tr117;
		}
		goto tr1;
	case 64:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr121;
		}
		goto tr1;
	case 65:
		switch ( s)[p] {
		case 32:
			goto tr122;
		case 44:
			goto tr123;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr124;
			}
		case ( s)[p] >= 9:
			goto tr122;
		}
		goto tr1;
	case 66:
		if ( s)[p] == 42 {
			goto tr2;
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr3;
		}
		goto tr1;
	case 67:
		switch ( s)[p] {
		case 32:
			goto tr125;
		case 44:
			goto tr126;
		case 45:
			goto tr127;
		case 47:
			goto tr128;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr129;
			}
		case ( s)[p] >= 9:
			goto tr125;
		}
		goto tr1;
	case 68:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr130;
		}
		goto tr1;
	case 69:
		switch ( s)[p] {
		case 32:
			goto tr131;
		case 44:
			goto tr132;
		case 47:
			goto tr133;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr134;
			}
		case ( s)[p] >= 9:
			goto tr131;
		}
		goto tr1;
	case 70:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr135;
		}
		goto tr1;
	case 71:
		switch ( s)[p] {
		case 32:
			goto tr136;
		case 44:
			goto tr137;
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr138;
			}
		case ( s)[p] >= 9:
			goto tr136;
		}
		goto tr1;
	}

	tr1: cs = 0; goto _again
	tr0: cs = 1; goto _again
	tr2: cs = 2; goto f0
	tr7: cs = 3; goto _again
	tr4: cs = 3; goto f2
	tr125: cs = 3; goto f41
	tr131: cs = 3; goto f42
	tr136: cs = 3; goto f43
	tr8: cs = 4; goto f0
	tr13: cs = 5; goto _again
	tr10: cs = 5; goto f4
	tr111: cs = 5; goto f36
	tr117: cs = 5; goto f38
	tr122: cs = 5; goto f40
	tr14: cs = 6; goto f5
	tr19: cs = 7; goto _again
	tr16: cs = 7; goto f7
	tr104: cs = 7; goto f34
	tr108: cs = 7; goto f35
	tr20: cs = 8; goto f5
	tr25: cs = 9; goto _again
	tr22: cs = 9; goto f8
	tr97: cs = 9; goto f32
	tr101: cs = 9; goto f33
	tr26: cs = 10; goto f5
	tr39: cs = 11; goto _again
	tr36: cs = 11; goto f9
	tr66: cs = 11; goto f17
	tr72: cs = 11; goto f19
	tr94: cs = 11; goto f31
	tr47: cs = 12; goto _again
	tr139: cs = 12; goto f44
	tr148: cs = 12; goto f48
	tr151: cs = 12; goto f49
	tr153: cs = 12; goto f50
	tr142: cs = 13; goto f45
	tr144: cs = 13; goto f46
	tr146: cs = 13; goto f47
	tr143: cs = 14; goto _again
	tr140: cs = 15; goto f44
	tr149: cs = 15; goto f48
	tr152: cs = 15; goto f49
	tr154: cs = 15; goto f50
	tr42: cs = 16; goto _again
	tr51: cs = 17; goto _again
	tr43: cs = 18; goto _again
	tr53: cs = 19; goto _again
	tr44: cs = 20; goto _again
	tr55: cs = 21; goto _again
	tr56: cs = 22; goto _again
	tr45: cs = 23; goto _again
	tr59: cs = 24; goto _again
	tr60: cs = 25; goto _again
	tr46: cs = 26; goto _again
	tr63: cs = 27; goto _again
	tr141: cs = 28; goto _again
	tr37: cs = 29; goto f9
	tr67: cs = 29; goto f17
	tr73: cs = 29; goto f19
	tr95: cs = 29; goto f31
	tr68: cs = 30; goto _again
	tr27: cs = 30; goto f6
	tr28: cs = 31; goto _again
	tr69: cs = 32; goto _again
	tr71: cs = 33; goto f18
	tr74: cs = 33; goto f20
	tr76: cs = 33; goto f21
	tr78: cs = 33; goto f22
	tr81: cs = 33; goto f23
	tr82: cs = 33; goto f24
	tr83: cs = 33; goto f25
	tr85: cs = 33; goto f26
	tr86: cs = 33; goto f27
	tr88: cs = 33; goto f28
	tr90: cs = 33; goto f29
	tr92: cs = 33; goto f30
	tr70: cs = 34; goto _again
	tr29: cs = 35; goto _again
	tr75: cs = 36; goto _again
	tr30: cs = 37; goto _again
	tr77: cs = 38; goto _again
	tr31: cs = 39; goto _again
	tr79: cs = 40; goto _again
	tr80: cs = 41; goto _again
	tr32: cs = 42; goto _again
	tr84: cs = 43; goto _again
	tr33: cs = 44; goto _again
	tr87: cs = 45; goto _again
	tr34: cs = 46; goto _again
	tr89: cs = 47; goto _again
	tr35: cs = 48; goto _again
	tr91: cs = 49; goto _again
	tr38: cs = 50; goto _again
	tr96: cs = 51; goto _again
	tr93: cs = 51; goto f6
	tr23: cs = 52; goto f8
	tr98: cs = 52; goto f32
	tr102: cs = 52; goto f33
	tr99: cs = 53; goto _again
	tr21: cs = 53; goto f6
	tr24: cs = 54; goto _again
	tr103: cs = 55; goto _again
	tr100: cs = 55; goto f6
	tr17: cs = 56; goto f7
	tr105: cs = 56; goto f34
	tr109: cs = 56; goto f35
	tr106: cs = 57; goto _again
	tr15: cs = 57; goto f6
	tr18: cs = 58; goto _again
	tr110: cs = 59; goto _again
	tr107: cs = 59; goto f6
	tr11: cs = 60; goto f4
	tr112: cs = 60; goto f36
	tr118: cs = 60; goto f38
	tr123: cs = 60; goto f40
	tr115: cs = 61; goto _again
	tr9: cs = 61; goto f1
	tr113: cs = 62; goto f37
	tr120: cs = 63; goto _again
	tr116: cs = 63; goto f6
	tr12: cs = 64; goto f3
	tr114: cs = 64; goto f37
	tr119: cs = 64; goto f39
	tr124: cs = 65; goto _again
	tr121: cs = 65; goto f6
	tr5: cs = 66; goto f2
	tr126: cs = 66; goto f41
	tr132: cs = 66; goto f42
	tr137: cs = 66; goto f43
	tr129: cs = 67; goto _again
	tr3: cs = 67; goto f1
	tr127: cs = 68; goto f37
	tr134: cs = 69; goto _again
	tr130: cs = 69; goto f6
	tr6: cs = 70; goto f3
	tr128: cs = 70; goto f37
	tr133: cs = 70; goto f39
	tr138: cs = 71; goto _again
	tr135: cs = 71; goto f6
	tr40: cs = 72; goto f5
	tr48: cs = 73; goto f5
	tr145: cs = 74; goto _again
	tr49: cs = 74; goto f6
	tr147: cs = 75; goto _again
	tr50: cs = 75; goto f6
	tr150: cs = 76; goto _again
	tr41: cs = 76; goto f6
	tr52: cs = 77; goto f10
	tr54: cs = 77; goto f11
	tr57: cs = 77; goto f12
	tr58: cs = 77; goto f13
	tr61: cs = 77; goto f14
	tr62: cs = 77; goto f15
	tr64: cs = 77; goto f16
	tr155: cs = 78; goto _again
	tr65: cs = 78; goto f6

	f6: _acts = 1; goto execFuncs
	f19: _acts = 3; goto execFuncs
	f49: _acts = 5; goto execFuncs
	f7: _acts = 7; goto execFuncs
	f9: _acts = 9; goto execFuncs
	f44: _acts = 11; goto execFuncs
	f45: _acts = 13; goto execFuncs
	f8: _acts = 15; goto execFuncs
	f13: _acts = 17; goto execFuncs
	f11: _acts = 19; goto execFuncs
	f15: _acts = 21; goto execFuncs
	f16: _acts = 23; goto execFuncs
	f14: _acts = 25; goto execFuncs
	f10: _acts = 27; goto execFuncs
	f12: _acts = 29; goto execFuncs
	f23: _acts = 31; goto execFuncs
	f22: _acts = 33; goto execFuncs
	f26: _acts = 35; goto execFuncs
	f18: _acts = 37; goto execFuncs
	f27: _acts = 39; goto execFuncs
	f25: _acts = 41; goto execFuncs
	f24: _acts = 43; goto execFuncs
	f20: _acts = 45; goto execFuncs
	f30: _acts = 47; goto execFuncs
	f29: _acts = 49; goto execFuncs
	f28: _acts = 51; goto execFuncs
	f21: _acts = 53; goto execFuncs
	f5: _acts = 55; goto execFuncs
	f3: _acts = 57; goto execFuncs
	f0: _acts = 59; goto execFuncs
	f17: _acts = 61; goto execFuncs
	f48: _acts = 64; goto execFuncs
	f34: _acts = 67; goto execFuncs
	f32: _acts = 70; goto execFuncs
	f46: _acts = 73; goto execFuncs
	f35: _acts = 76; goto execFuncs
	f31: _acts = 79; goto execFuncs
	f50: _acts = 82; goto execFuncs
	f47: _acts = 85; goto execFuncs
	f33: _acts = 88; goto execFuncs
	f37: _acts = 91; goto execFuncs
	f39: _acts = 94; goto execFuncs
	f2: _acts = 97; goto execFuncs
	f4: _acts = 100; goto execFuncs
	f1: _acts = 103; goto execFuncs
	f41: _acts = 106; goto execFuncs
	f36: _acts = 110; goto execFuncs
	f42: _acts = 114; goto execFuncs
	f38: _acts = 118; goto execFuncs
	f43: _acts = 122; goto execFuncs
	f40: _acts = 126; goto execFuncs

execFuncs:
	_nacts = uint(_parse_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _parse_actions[_acts - 1] {
		case 0:
//line parse.rl:29

            mark = p;
        
		case 1:
//line parse.rl:41

            if m>12 || m<1{
                return nt, fmt.Errorf("invalid month %d", m)
            }
            nt.month |= 1<<(m-1) // we zero index months
        
		case 2:
//line parse.rl:55

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
        
		case 3:
//line parse.rl:69

            if m>24 || m <0 {
                return nt, fmt.Errorf("invalid hour of day %d", m)
            }
            nt.hour |= 1<<m
        
		case 4:
//line parse.rl:76

            if m>30 || m <0 {
                return nt,  fmt.Errorf("invalid day of month %d", m)
            }
            nt.dom |= 1<<(m-1)// again we zero index day
        
		case 5:
//line parse.rl:83

            if m<1970 || m>2099{
                return nt, errors.New("only years on interval [1970, 2099] are supported")
            }
            nt.year.set(int(m))
        
		case 6:
//line parse.rl:90

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
        
		case 7:
//line parse.rl:177

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
        
		case 8:
//line parse.rl:264

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
        
		case 9:
//line parse.rl:298

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
        
		case 10:
//line parse.rl:321

            const sundaysAtFirst = uint64(1 | 1<<7 | 1<<14 | 1<<21 | 1<<28 | 1<<35 | 1<<42)
            {
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
        
		case 11:
//line parse.rl:343

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
        
		case 12:
//line parse.rl:360

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
        
		case 13:
//line parse.rl:403

            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
		case 14:
//line parse.rl:415
m=0
		case 15:
//line parse.rl:415
m=1
		case 16:
//line parse.rl:415
m=2
		case 17:
//line parse.rl:415
m=3
		case 18:
//line parse.rl:415
m=4
		case 19:
//line parse.rl:415
m=5
		case 20:
//line parse.rl:415
m=6
		case 21:
//line parse.rl:416
m=1
		case 22:
//line parse.rl:416
m=2
		case 23:
//line parse.rl:416
m=3
		case 24:
//line parse.rl:416
m=4
		case 25:
//line parse.rl:416
m=5
		case 26:
//line parse.rl:416
m=6
		case 27:
//line parse.rl:416
m=7
		case 28:
//line parse.rl:416
m=8
		case 29:
//line parse.rl:416
m=9
		case 30:
//line parse.rl:416
m=10
		case 31:
//line parse.rl:416
m=11
		case 32:
//line parse.rl:416
m=12
		case 33:
//line parse.rl:424
m=1
		case 34:
//line parse.rl:428
 start=0;end=59;m=1;d=1; 
		case 35:
//line parse.rl:428
 start=m; end=0;d=0;
		case 36:
//line parse.rl:428
 end=m; d=1;
		case 37:
//line parse.rl:428
d=0;
		case 38:
//line parse.rl:429
d=m
		case 39:
//line parse.rl:434
d=m
//line parse.go:1713
		}
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
		__acts := int(_parse_eof_actions[cs])
		__nacts := uint(_parse_actions[__acts]); __acts++
		for ; __nacts > 0; __nacts-- {
			__acts++
			switch _parse_actions[__acts - 1] {
			case 2:
//line parse.rl:55

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
        
			case 5:
//line parse.rl:83

            if m<1970 || m>2099{
                return nt, errors.New("only years on interval [1970, 2099] are supported")
            }
            nt.year.set(int(m))
        
			case 10:
//line parse.rl:321

            const sundaysAtFirst = uint64(1 | 1<<7 | 1<<14 | 1<<21 | 1<<28 | 1<<35 | 1<<42)
            {
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
        
			case 11:
//line parse.rl:343

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
        
			case 13:
//line parse.rl:403

            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.go:1801
			}
		}
	}

	_out: {}
	}

//line parse.rl:450
    // check current month
    // check the next year, unless feb 1<<29th is allowed then check the next 1<<8 years
    // if year isn't specified accept any year
    fmt.Println("nt.year", nt.year)
    fmt.Println("vars", cs, p, pe, eof)
    fmt.Printf("in parser: year %b, %b \nmonth %b, \ndom %b, \ndow %b, \nhour %b, \nmin %b, \ns %b\n", nt.year.low, nt.year.high, nt.month, nt.dom, nt.dow, nt.hour, nt.minute, nt.second)
    if nt.minute==0||nt.hour==0||nt.dom==0||nt.month==0||nt.dow==0 {
        return nt, errors.New("failed to parse cron string")
    }

    if nt.year.high==0 && nt.year.low==0 {
        nt.year.high = ^uint64(0)
        nt.year.low = ^uint64(0)
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

