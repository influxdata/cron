
//line parse.rl:1
// -*-go-*-
package cron

import (
    "strconv"
    "fmt"
    "time"
    "errors"
)


//line parse.rl:12

//line parse.rl:13

//line parse.go:19
const parse_start int = 1
const parse_first_final int = 97
const parse_error int = 0

const parse_en_main int = 1
const parse_en_main_sevenPos int = 18


//line parse.rl:14

func parse(s string, tz *time.Location)(nextTime, error){
    nt:=nextTime{loc:tz}
    cs, p, pe, eof:= 0, 0,len(s), len(s)

    mark, backtrack := 0,0
    _ = mark
    _ = backtrack
    m, d, start, end:=uint64(0), uint64(0), uint64(0), uint64(0)
    _ = d
    // TODO(docmerlin): handle ranges
    var err error
    
//line parse.go:42
	{
	cs = parse_start
	}

//line parse.rl:27
    //m,h := 1<<0,1<<0
    
//line parse.rl:483

    
//line parse.go:53
	{
	if p == pe {
		goto _test_eof
	}
	goto _resume

_again:
	switch cs {
	case 1:
		goto st1
	case 0:
		goto st0
	case 2:
		goto st2
	case 3:
		goto st3
	case 4:
		goto st4
	case 5:
		goto st5
	case 6:
		goto st6
	case 7:
		goto st7
	case 8:
		goto st8
	case 9:
		goto st9
	case 10:
		goto st10
	case 11:
		goto st11
	case 12:
		goto st12
	case 13:
		goto st13
	case 14:
		goto st14
	case 15:
		goto st15
	case 16:
		goto st16
	case 17:
		goto st17
	case 18:
		goto st18
	case 19:
		goto st19
	case 20:
		goto st20
	case 21:
		goto st21
	case 22:
		goto st22
	case 23:
		goto st23
	case 24:
		goto st24
	case 25:
		goto st25
	case 26:
		goto st26
	case 27:
		goto st27
	case 28:
		goto st28
	case 29:
		goto st29
	case 30:
		goto st30
	case 31:
		goto st31
	case 32:
		goto st32
	case 33:
		goto st33
	case 34:
		goto st34
	case 35:
		goto st35
	case 36:
		goto st36
	case 37:
		goto st37
	case 38:
		goto st38
	case 39:
		goto st39
	case 40:
		goto st40
	case 41:
		goto st41
	case 42:
		goto st42
	case 43:
		goto st43
	case 44:
		goto st44
	case 45:
		goto st45
	case 46:
		goto st46
	case 47:
		goto st47
	case 48:
		goto st48
	case 49:
		goto st49
	case 50:
		goto st50
	case 51:
		goto st51
	case 52:
		goto st52
	case 53:
		goto st53
	case 54:
		goto st54
	case 55:
		goto st55
	case 56:
		goto st56
	case 57:
		goto st57
	case 58:
		goto st58
	case 59:
		goto st59
	case 60:
		goto st60
	case 61:
		goto st61
	case 62:
		goto st62
	case 63:
		goto st63
	case 64:
		goto st64
	case 65:
		goto st65
	case 66:
		goto st66
	case 67:
		goto st67
	case 68:
		goto st68
	case 69:
		goto st69
	case 70:
		goto st70
	case 71:
		goto st71
	case 72:
		goto st72
	case 73:
		goto st73
	case 74:
		goto st74
	case 75:
		goto st75
	case 76:
		goto st76
	case 77:
		goto st77
	case 78:
		goto st78
	case 79:
		goto st79
	case 80:
		goto st80
	case 81:
		goto st81
	case 82:
		goto st82
	case 83:
		goto st83
	case 84:
		goto st84
	case 85:
		goto st85
	case 86:
		goto st86
	case 87:
		goto st87
	case 88:
		goto st88
	case 89:
		goto st89
	case 90:
		goto st90
	case 91:
		goto st91
	case 92:
		goto st92
	case 93:
		goto st93
	case 94:
		goto st94
	case 95:
		goto st95
	case 96:
		goto st96
	}

	if p++; p == pe {
		goto _test_eof
	}
_resume:
	switch cs {
	case 1:
		goto st_case_1
	case 0:
		goto st_case_0
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
	case 47:
		goto st_case_47
	case 48:
		goto st_case_48
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 51:
		goto st_case_51
	case 52:
		goto st_case_52
	case 53:
		goto st_case_53
	case 54:
		goto st_case_54
	case 55:
		goto st_case_55
	case 56:
		goto st_case_56
	case 57:
		goto st_case_57
	case 58:
		goto st_case_58
	case 59:
		goto st_case_59
	case 60:
		goto st_case_60
	case 61:
		goto st_case_61
	case 62:
		goto st_case_62
	case 63:
		goto st_case_63
	case 64:
		goto st_case_64
	case 65:
		goto st_case_65
	case 66:
		goto st_case_66
	case 67:
		goto st_case_67
	case 68:
		goto st_case_68
	case 69:
		goto st_case_69
	case 70:
		goto st_case_70
	case 71:
		goto st_case_71
	case 72:
		goto st_case_72
	case 73:
		goto st_case_73
	case 74:
		goto st_case_74
	case 75:
		goto st_case_75
	case 76:
		goto st_case_76
	case 77:
		goto st_case_77
	case 78:
		goto st_case_78
	case 79:
		goto st_case_79
	case 80:
		goto st_case_80
	case 81:
		goto st_case_81
	case 82:
		goto st_case_82
	case 83:
		goto st_case_83
	case 84:
		goto st_case_84
	case 85:
		goto st_case_85
	case 86:
		goto st_case_86
	case 87:
		goto st_case_87
	case 88:
		goto st_case_88
	case 89:
		goto st_case_89
	case 90:
		goto st_case_90
	case 91:
		goto st_case_91
	case 92:
		goto st_case_92
	case 93:
		goto st_case_93
	case 94:
		goto st_case_94
	case 95:
		goto st_case_95
	case 96:
		goto st_case_96
	}
	goto st_out
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch ( s)[p] {
		case 32:
			goto st1
		case 42:
			goto tr2
		}
		switch {
		case ( s)[p] < 47:
			switch {
			case ( s)[p] > 13:
				if 44 <= ( s)[p] && ( s)[p] <= 45 {
					goto tr2
				}
			case ( s)[p] >= 9:
				goto st1
			}
		case ( s)[p] > 57:
			switch {
			case ( s)[p] > 90:
				if 97 <= ( s)[p] && ( s)[p] <= 122 {
					goto tr2
				}
			case ( s)[p] >= 65:
				goto tr2
			}
		default:
			goto tr2
		}
		goto tr0
tr0:
	cs = 0
//line parse.rl:463
 
                if p!=pe{
                    fmt.Println(p,pe, backtrack)
                }
                if p==pe{
                    if backtrack == 7{
                        fmt.Println("here case 7")
                        p=mark
                        fmt.Println("here case 7", p, pe, mark)
                        //fhold;
                        fmt.Println("here case 7.b")
                        //cs = fentry(sevenPos);
                        p = ( p) - 1

                        cs = 18;
                        //fgoto *fentry(sevenPos);
                    }
                }
            
	goto _again
//line parse.go:516
st_case_0:
	st0:
		cs = 0
		goto _out
tr2:
//line parse.rl:29

            mark = p;
        
	goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line parse.go:532
		switch ( s)[p] {
		case 32:
			goto tr3
		case 42:
			goto st2
		}
		switch {
		case ( s)[p] < 47:
			switch {
			case ( s)[p] > 13:
				if 44 <= ( s)[p] && ( s)[p] <= 45 {
					goto st2
				}
			case ( s)[p] >= 9:
				goto tr3
			}
		case ( s)[p] > 57:
			switch {
			case ( s)[p] > 90:
				if 97 <= ( s)[p] && ( s)[p] <= 122 {
					goto st2
				}
			case ( s)[p] >= 65:
				goto st2
			}
		default:
			goto st2
		}
		goto tr0
tr3:
//line parse.rl:403

            backtrack++
        
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line parse.go:573
		switch ( s)[p] {
		case 32:
			goto st3
		case 42:
			goto st4
		}
		switch {
		case ( s)[p] < 47:
			switch {
			case ( s)[p] > 13:
				if 44 <= ( s)[p] && ( s)[p] <= 45 {
					goto st4
				}
			case ( s)[p] >= 9:
				goto st3
			}
		case ( s)[p] > 57:
			switch {
			case ( s)[p] > 90:
				if 97 <= ( s)[p] && ( s)[p] <= 122 {
					goto st4
				}
			case ( s)[p] >= 65:
				goto st4
			}
		default:
			goto st4
		}
		goto tr0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch ( s)[p] {
		case 32:
			goto tr7
		case 42:
			goto st4
		}
		switch {
		case ( s)[p] < 47:
			switch {
			case ( s)[p] > 13:
				if 44 <= ( s)[p] && ( s)[p] <= 45 {
					goto st4
				}
			case ( s)[p] >= 9:
				goto tr7
			}
		case ( s)[p] > 57:
			switch {
			case ( s)[p] > 90:
				if 97 <= ( s)[p] && ( s)[p] <= 122 {
					goto st4
				}
			case ( s)[p] >= 65:
				goto st4
			}
		default:
			goto st4
		}
		goto tr0
tr7:
//line parse.rl:403

            backtrack++
        
	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line parse.go:648
		switch ( s)[p] {
		case 32:
			goto st5
		case 42:
			goto st6
		}
		switch {
		case ( s)[p] < 47:
			switch {
			case ( s)[p] > 13:
				if 44 <= ( s)[p] && ( s)[p] <= 45 {
					goto st6
				}
			case ( s)[p] >= 9:
				goto st5
			}
		case ( s)[p] > 57:
			switch {
			case ( s)[p] > 90:
				if 97 <= ( s)[p] && ( s)[p] <= 122 {
					goto st6
				}
			case ( s)[p] >= 65:
				goto st6
			}
		default:
			goto st6
		}
		goto tr0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch ( s)[p] {
		case 32:
			goto tr10
		case 42:
			goto st6
		}
		switch {
		case ( s)[p] < 47:
			switch {
			case ( s)[p] > 13:
				if 44 <= ( s)[p] && ( s)[p] <= 45 {
					goto st6
				}
			case ( s)[p] >= 9:
				goto tr10
			}
		case ( s)[p] > 57:
			switch {
			case ( s)[p] > 90:
				if 97 <= ( s)[p] && ( s)[p] <= 122 {
					goto st6
				}
			case ( s)[p] >= 65:
				goto st6
			}
		default:
			goto st6
		}
		goto tr0
tr10:
//line parse.rl:403

            backtrack++
        
	goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line parse.go:723
		switch ( s)[p] {
		case 32:
			goto st7
		case 42:
			goto st8
		}
		switch {
		case ( s)[p] < 47:
			switch {
			case ( s)[p] > 13:
				if 44 <= ( s)[p] && ( s)[p] <= 45 {
					goto st8
				}
			case ( s)[p] >= 9:
				goto st7
			}
		case ( s)[p] > 57:
			switch {
			case ( s)[p] > 90:
				if 97 <= ( s)[p] && ( s)[p] <= 122 {
					goto st8
				}
			case ( s)[p] >= 65:
				goto st8
			}
		default:
			goto st8
		}
		goto tr0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		switch ( s)[p] {
		case 32:
			goto tr13
		case 42:
			goto st8
		}
		switch {
		case ( s)[p] < 47:
			switch {
			case ( s)[p] > 13:
				if 44 <= ( s)[p] && ( s)[p] <= 45 {
					goto st8
				}
			case ( s)[p] >= 9:
				goto tr13
			}
		case ( s)[p] > 57:
			switch {
			case ( s)[p] > 90:
				if 97 <= ( s)[p] && ( s)[p] <= 122 {
					goto st8
				}
			case ( s)[p] >= 65:
				goto st8
			}
		default:
			goto st8
		}
		goto tr0
tr13:
//line parse.rl:403

            backtrack++
        
	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line parse.go:798
		switch ( s)[p] {
		case 32:
			goto st9
		case 42:
			goto st10
		}
		switch {
		case ( s)[p] < 47:
			switch {
			case ( s)[p] > 13:
				if 44 <= ( s)[p] && ( s)[p] <= 45 {
					goto st10
				}
			case ( s)[p] >= 9:
				goto st9
			}
		case ( s)[p] > 57:
			switch {
			case ( s)[p] > 90:
				if 97 <= ( s)[p] && ( s)[p] <= 122 {
					goto st10
				}
			case ( s)[p] >= 65:
				goto st10
			}
		default:
			goto st10
		}
		goto tr0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch ( s)[p] {
		case 32:
			goto tr16
		case 42:
			goto st10
		}
		switch {
		case ( s)[p] < 47:
			switch {
			case ( s)[p] > 13:
				if 44 <= ( s)[p] && ( s)[p] <= 45 {
					goto st10
				}
			case ( s)[p] >= 9:
				goto tr16
			}
		case ( s)[p] > 57:
			switch {
			case ( s)[p] > 90:
				if 97 <= ( s)[p] && ( s)[p] <= 122 {
					goto st10
				}
			case ( s)[p] >= 65:
				goto st10
			}
		default:
			goto st10
		}
		goto tr0
tr16:
//line parse.rl:403

            backtrack++
        
	goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line parse.go:873
		switch ( s)[p] {
		case 32:
			goto st11
		case 42:
			goto st12
		}
		switch {
		case ( s)[p] < 47:
			switch {
			case ( s)[p] > 13:
				if 44 <= ( s)[p] && ( s)[p] <= 45 {
					goto st12
				}
			case ( s)[p] >= 9:
				goto st11
			}
		case ( s)[p] > 57:
			switch {
			case ( s)[p] > 90:
				if 97 <= ( s)[p] && ( s)[p] <= 122 {
					goto st12
				}
			case ( s)[p] >= 65:
				goto st12
			}
		default:
			goto st12
		}
		goto tr0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch ( s)[p] {
		case 32:
			goto tr19
		case 42:
			goto st12
		}
		switch {
		case ( s)[p] < 47:
			switch {
			case ( s)[p] > 13:
				if 44 <= ( s)[p] && ( s)[p] <= 45 {
					goto st12
				}
			case ( s)[p] >= 9:
				goto tr19
			}
		case ( s)[p] > 57:
			switch {
			case ( s)[p] > 90:
				if 97 <= ( s)[p] && ( s)[p] <= 122 {
					goto st12
				}
			case ( s)[p] >= 65:
				goto st12
			}
		default:
			goto st12
		}
		goto tr0
tr19:
//line parse.rl:403

            backtrack++
        
	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line parse.go:948
		switch ( s)[p] {
		case 32:
			goto st13
		case 42:
			goto st14
		}
		switch {
		case ( s)[p] < 47:
			switch {
			case ( s)[p] > 13:
				if 44 <= ( s)[p] && ( s)[p] <= 45 {
					goto st14
				}
			case ( s)[p] >= 9:
				goto st13
			}
		case ( s)[p] > 57:
			switch {
			case ( s)[p] > 90:
				if 97 <= ( s)[p] && ( s)[p] <= 122 {
					goto st14
				}
			case ( s)[p] >= 65:
				goto st14
			}
		default:
			goto st14
		}
		goto tr0
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		switch ( s)[p] {
		case 32:
			goto tr22
		case 42:
			goto st14
		}
		switch {
		case ( s)[p] < 47:
			switch {
			case ( s)[p] > 13:
				if 44 <= ( s)[p] && ( s)[p] <= 45 {
					goto st14
				}
			case ( s)[p] >= 9:
				goto tr22
			}
		case ( s)[p] > 57:
			switch {
			case ( s)[p] > 90:
				if 97 <= ( s)[p] && ( s)[p] <= 122 {
					goto st14
				}
			case ( s)[p] >= 65:
				goto st14
			}
		default:
			goto st14
		}
		goto tr0
tr22:
//line parse.rl:403

            backtrack++
        
	goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line parse.go:1023
		switch ( s)[p] {
		case 32:
			goto st15
		case 42:
			goto st16
		}
		switch {
		case ( s)[p] < 47:
			switch {
			case ( s)[p] > 13:
				if 44 <= ( s)[p] && ( s)[p] <= 45 {
					goto st16
				}
			case ( s)[p] >= 9:
				goto st15
			}
		case ( s)[p] > 57:
			switch {
			case ( s)[p] > 90:
				if 97 <= ( s)[p] && ( s)[p] <= 122 {
					goto st16
				}
			case ( s)[p] >= 65:
				goto st16
			}
		default:
			goto st16
		}
		goto tr0
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		switch ( s)[p] {
		case 32:
			goto tr25
		case 42:
			goto st16
		}
		switch {
		case ( s)[p] < 47:
			switch {
			case ( s)[p] > 13:
				if 44 <= ( s)[p] && ( s)[p] <= 45 {
					goto st16
				}
			case ( s)[p] >= 9:
				goto tr25
			}
		case ( s)[p] > 57:
			switch {
			case ( s)[p] > 90:
				if 97 <= ( s)[p] && ( s)[p] <= 122 {
					goto st16
				}
			case ( s)[p] >= 65:
				goto st16
			}
		default:
			goto st16
		}
		goto tr0
tr25:
//line parse.rl:403

            backtrack++
        
	goto st17
tr26:
//line parse.rl:407

            return nt, fmt.Errorf("too many positions in cron")
        
	goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line parse.go:1104
		if ( s)[p] == 32 {
			goto tr26
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr26
		}
		goto tr0
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		if ( s)[p] == 42 {
			goto tr27
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr29
		}
		goto st0
tr27:
//line parse.rl:481
panic("")
//line parse.rl:441
d=0;
	goto st19
tr172:
//line parse.rl:441
d=0;
	goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line parse.go:1139
		switch ( s)[p] {
		case 32:
			goto tr30
		case 44:
			goto tr31
		case 47:
			goto tr32
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr30
		}
		goto st0
tr30:
//line parse.rl:441
 start=0;end=59;m=1;d=1; 
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
        
	goto st20
tr174:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:441
 start=m; end=0;d=0;
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
        
	goto st20
tr180:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:441
 end=m; d=1;
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
        
	goto st20
tr185:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:442
d=m
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
        
	goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line parse.go:1545
		switch ( s)[p] {
		case 32:
			goto st20
		case 42:
			goto tr34
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr35
			}
		case ( s)[p] >= 9:
			goto st20
		}
		goto st0
tr34:
//line parse.rl:441
d=0;
	goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line parse.go:1570
		switch ( s)[p] {
		case 32:
			goto tr36
		case 44:
			goto tr37
		case 47:
			goto tr38
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr36
		}
		goto st0
tr36:
//line parse.rl:441
 start=0;end=59;m=1;d=1; 
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
        
	goto st22
tr158:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:441
 start=m; end=0;d=0;
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
        
	goto st22
tr164:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:441
 end=m; d=1;
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
        
	goto st22
tr169:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:447
d=m
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
        
	goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line parse.go:1976
		switch ( s)[p] {
		case 32:
			goto st22
		case 42:
			goto tr40
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr41
			}
		case ( s)[p] >= 9:
			goto st22
		}
		goto st0
tr40:
//line parse.rl:437
m=1
	goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line parse.go:2001
		switch ( s)[p] {
		case 32:
			goto tr42
		case 44:
			goto tr43
		case 47:
			goto st83
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr42
		}
		goto st0
tr42:
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
        
	goto st24
tr151:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:69

            if m>24 || m <0 {
                return nt, fmt.Errorf("invalid hour of day %d", m)
            }
            nt.hour |= 1<<m
        
	goto st24
tr155:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
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
        
	goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line parse.go:2116
		switch ( s)[p] {
		case 32:
			goto st24
		case 42:
			goto tr46
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr47
			}
		case ( s)[p] >= 9:
			goto st24
		}
		goto st0
tr46:
//line parse.rl:437
m=1
	goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line parse.go:2141
		switch ( s)[p] {
		case 32:
			goto tr48
		case 44:
			goto tr49
		case 47:
			goto st79
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr48
		}
		goto st0
tr48:
//line parse.rl:361

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
        
	goto st26
tr144:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:76

            if m>30 || m <0 {
                return nt,  fmt.Errorf("invalid day of month %d", m)
            }
            nt.dom |= 1<<(m-1)// again we zero index day
        
	goto st26
tr148:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:361

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
        
	goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line parse.go:2272
		switch ( s)[p] {
		case 32:
			goto st26
		case 42:
			goto tr52
		case 65:
			goto st56
		case 68:
			goto st60
		case 70:
			goto st62
		case 74:
			goto st64
		case 77:
			goto st67
		case 78:
			goto st69
		case 79:
			goto st71
		case 83:
			goto st73
		case 97:
			goto st56
		case 100:
			goto st60
		case 102:
			goto st62
		case 106:
			goto st64
		case 109:
			goto st67
		case 110:
			goto st69
		case 111:
			goto st71
		case 115:
			goto st73
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr53
			}
		case ( s)[p] >= 9:
			goto st26
		}
		goto st0
tr52:
//line parse.rl:437
m=1
	goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
//line parse.go:2329
		switch ( s)[p] {
		case 32:
			goto tr62
		case 44:
			goto tr63
		case 47:
			goto st75
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr62
		}
		goto st0
tr62:
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
        
	goto st28
tr113:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:41

            if m>12 || m<1{
                return nt, fmt.Errorf("invalid month %d", m)
            }
            nt.month |= 1<<(m-1) // we zero index months
        
	goto st28
tr119:
//line parse.rl:41

            if m>12 || m<1{
                return nt, fmt.Errorf("invalid month %d", m)
            }
            nt.month |= 1<<(m-1) // we zero index months
        
	goto st28
tr141:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
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
        
	goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line parse.go:2431
		switch ( s)[p] {
		case 32:
			goto st28
		case 42:
			goto tr66
		case 70:
			goto st39
		case 77:
			goto st42
		case 83:
			goto st44
		case 84:
			goto st47
		case 87:
			goto st50
		case 102:
			goto st39
		case 109:
			goto st42
		case 115:
			goto st44
		case 116:
			goto st47
		case 119:
			goto st50
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr67
			}
		case ( s)[p] >= 9:
			goto st28
		}
		goto st0
tr66:
//line parse.rl:437
m=1
	goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
//line parse.go:2476
		switch ( s)[p] {
		case 32:
			goto tr73
		case 44:
			goto tr74
		case 47:
			goto st52
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr73
		}
		goto st0
tr73:
//line parse.rl:321

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
        
	goto st30
tr90:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
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
        
	goto st30
tr95:
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
        
	goto st30
tr110:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:321

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
        
	goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
//line parse.go:2592
		switch ( s)[p] {
		case 32:
			goto st30
		case 42:
			goto tr77
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr78
			}
		case ( s)[p] >= 9:
			goto st30
		}
		goto st0
tr77:
//line parse.rl:437
m=1
	goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
//line parse.go:2617
		switch ( s)[p] {
		case 32:
			goto tr79
		case 44:
			goto tr80
		case 47:
			goto st35
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr79
		}
		goto st0
tr79:
//line parse.rl:344

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
        
	goto st32
tr83:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:83

            if m<1970 || m>2099{
                return nt, errors.New("only years on interval [1970, 2099] are supported")
            }
            nt.year.set(int(m))
        
	goto st32
tr87:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:344

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
        
	goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
//line parse.go:2692
		if ( s)[p] == 32 {
			goto st32
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto st32
		}
		goto st0
tr80:
//line parse.rl:344

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
        
	goto st33
tr84:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:83

            if m<1970 || m>2099{
                return nt, errors.New("only years on interval [1970, 2099] are supported")
            }
            nt.year.set(int(m))
        
	goto st33
tr88:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:344

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
        
	goto st33
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
//line parse.go:2762
		if ( s)[p] == 42 {
			goto tr77
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr78
		}
		goto st0
tr78:
//line parse.rl:29

            mark = p;
        
	goto st34
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
//line parse.go:2781
		switch ( s)[p] {
		case 32:
			goto tr83
		case 44:
			goto tr84
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st34
			}
		case ( s)[p] >= 9:
			goto tr83
		}
		goto st0
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr86
		}
		goto st0
tr86:
//line parse.rl:29

            mark = p;
        
	goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
//line parse.go:2817
		switch ( s)[p] {
		case 32:
			goto tr87
		case 44:
			goto tr88
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st36
			}
		case ( s)[p] >= 9:
			goto tr87
		}
		goto st0
tr74:
//line parse.rl:321

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
        
	goto st37
tr91:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
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
        
	goto st37
tr96:
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
        
	goto st37
tr111:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:321

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
        
	goto st37
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
//line parse.go:2936
		switch ( s)[p] {
		case 42:
			goto tr66
		case 70:
			goto st39
		case 77:
			goto st42
		case 83:
			goto st44
		case 84:
			goto st47
		case 87:
			goto st50
		case 102:
			goto st39
		case 109:
			goto st42
		case 115:
			goto st44
		case 116:
			goto st47
		case 119:
			goto st50
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr67
		}
		goto st0
tr67:
//line parse.rl:29

            mark = p;
        
	goto st38
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
//line parse.go:2976
		switch ( s)[p] {
		case 32:
			goto tr90
		case 44:
			goto tr91
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st38
			}
		case ( s)[p] >= 9:
			goto tr90
		}
		goto st0
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		switch ( s)[p] {
		case 82:
			goto st40
		case 114:
			goto st40
		}
		goto st0
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		switch ( s)[p] {
		case 73:
			goto tr94
		case 105:
			goto tr94
		}
		goto st0
tr94:
//line parse.rl:428
m=5
	goto st41
tr98:
//line parse.rl:428
m=1
	goto st41
tr101:
//line parse.rl:428
m=6
	goto st41
tr102:
//line parse.rl:428
m=0
	goto st41
tr105:
//line parse.rl:428
m=4
	goto st41
tr106:
//line parse.rl:428
m=2
	goto st41
tr108:
//line parse.rl:428
m=3
	goto st41
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
//line parse.go:3049
		switch ( s)[p] {
		case 32:
			goto tr95
		case 44:
			goto tr96
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr95
		}
		goto st0
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		switch ( s)[p] {
		case 79:
			goto st43
		case 111:
			goto st43
		}
		goto st0
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		switch ( s)[p] {
		case 78:
			goto tr98
		case 110:
			goto tr98
		}
		goto st0
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
		switch ( s)[p] {
		case 65:
			goto st45
		case 85:
			goto st46
		case 97:
			goto st45
		case 117:
			goto st46
		}
		goto st0
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		switch ( s)[p] {
		case 84:
			goto tr101
		case 116:
			goto tr101
		}
		goto st0
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		switch ( s)[p] {
		case 78:
			goto tr102
		case 110:
			goto tr102
		}
		goto st0
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		switch ( s)[p] {
		case 72:
			goto st48
		case 85:
			goto st49
		case 104:
			goto st48
		case 117:
			goto st49
		}
		goto st0
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		switch ( s)[p] {
		case 85:
			goto tr105
		case 117:
			goto tr105
		}
		goto st0
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		switch ( s)[p] {
		case 69:
			goto tr106
		case 101:
			goto tr106
		}
		goto st0
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		switch ( s)[p] {
		case 69:
			goto st51
		case 101:
			goto st51
		}
		goto st0
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
		switch ( s)[p] {
		case 68:
			goto tr108
		case 100:
			goto tr108
		}
		goto st0
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr109
		}
		goto st0
tr109:
//line parse.rl:29

            mark = p;
        
	goto st53
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
//line parse.go:3208
		switch ( s)[p] {
		case 32:
			goto tr110
		case 44:
			goto tr111
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st53
			}
		case ( s)[p] >= 9:
			goto tr110
		}
		goto st0
tr63:
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
        
	goto st54
tr114:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:41

            if m>12 || m<1{
                return nt, fmt.Errorf("invalid month %d", m)
            }
            nt.month |= 1<<(m-1) // we zero index months
        
	goto st54
tr120:
//line parse.rl:41

            if m>12 || m<1{
                return nt, fmt.Errorf("invalid month %d", m)
            }
            nt.month |= 1<<(m-1) // we zero index months
        
	goto st54
tr142:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
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
        
	goto st54
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
//line parse.go:3313
		switch ( s)[p] {
		case 42:
			goto tr52
		case 65:
			goto st56
		case 68:
			goto st60
		case 70:
			goto st62
		case 74:
			goto st64
		case 77:
			goto st67
		case 78:
			goto st69
		case 79:
			goto st71
		case 83:
			goto st73
		case 97:
			goto st56
		case 100:
			goto st60
		case 102:
			goto st62
		case 106:
			goto st64
		case 109:
			goto st67
		case 110:
			goto st69
		case 111:
			goto st71
		case 115:
			goto st73
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr53
		}
		goto st0
tr53:
//line parse.rl:29

            mark = p;
        
	goto st55
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
//line parse.go:3365
		switch ( s)[p] {
		case 32:
			goto tr113
		case 44:
			goto tr114
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st55
			}
		case ( s)[p] >= 9:
			goto tr113
		}
		goto st0
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		switch ( s)[p] {
		case 80:
			goto st57
		case 85:
			goto st59
		case 112:
			goto st57
		case 117:
			goto st59
		}
		goto st0
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		switch ( s)[p] {
		case 82:
			goto tr118
		case 114:
			goto tr118
		}
		goto st0
tr118:
//line parse.rl:429
m=4
	goto st58
tr121:
//line parse.rl:429
m=8
	goto st58
tr123:
//line parse.rl:429
m=12
	goto st58
tr125:
//line parse.rl:429
m=2
	goto st58
tr128:
//line parse.rl:429
m=1
	goto st58
tr129:
//line parse.rl:429
m=7
	goto st58
tr130:
//line parse.rl:429
m=6
	goto st58
tr132:
//line parse.rl:429
m=3
	goto st58
tr133:
//line parse.rl:429
m=5
	goto st58
tr135:
//line parse.rl:429
m=11
	goto st58
tr137:
//line parse.rl:429
m=10
	goto st58
tr139:
//line parse.rl:429
m=9
	goto st58
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
//line parse.go:3462
		switch ( s)[p] {
		case 32:
			goto tr119
		case 44:
			goto tr120
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr119
		}
		goto st0
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		switch ( s)[p] {
		case 71:
			goto tr121
		case 103:
			goto tr121
		}
		goto st0
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
		switch ( s)[p] {
		case 69:
			goto st61
		case 101:
			goto st61
		}
		goto st0
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		switch ( s)[p] {
		case 67:
			goto tr123
		case 99:
			goto tr123
		}
		goto st0
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		switch ( s)[p] {
		case 69:
			goto st63
		case 101:
			goto st63
		}
		goto st0
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		switch ( s)[p] {
		case 66:
			goto tr125
		case 98:
			goto tr125
		}
		goto st0
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		switch ( s)[p] {
		case 65:
			goto st65
		case 85:
			goto st66
		case 97:
			goto st65
		case 117:
			goto st66
		}
		goto st0
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		switch ( s)[p] {
		case 78:
			goto tr128
		case 110:
			goto tr128
		}
		goto st0
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
		switch ( s)[p] {
		case 76:
			goto tr129
		case 78:
			goto tr130
		case 108:
			goto tr129
		case 110:
			goto tr130
		}
		goto st0
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
		switch ( s)[p] {
		case 65:
			goto st68
		case 97:
			goto st68
		}
		goto st0
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
		switch ( s)[p] {
		case 82:
			goto tr132
		case 89:
			goto tr133
		case 114:
			goto tr132
		case 121:
			goto tr133
		}
		goto st0
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		switch ( s)[p] {
		case 79:
			goto st70
		case 111:
			goto st70
		}
		goto st0
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
		switch ( s)[p] {
		case 86:
			goto tr135
		case 118:
			goto tr135
		}
		goto st0
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
		switch ( s)[p] {
		case 67:
			goto st72
		case 99:
			goto st72
		}
		goto st0
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
		switch ( s)[p] {
		case 84:
			goto tr137
		case 116:
			goto tr137
		}
		goto st0
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
		switch ( s)[p] {
		case 69:
			goto st74
		case 101:
			goto st74
		}
		goto st0
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
		switch ( s)[p] {
		case 80:
			goto tr139
		case 112:
			goto tr139
		}
		goto st0
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr140
		}
		goto st0
tr140:
//line parse.rl:29

            mark = p;
        
	goto st76
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
//line parse.go:3697
		switch ( s)[p] {
		case 32:
			goto tr141
		case 44:
			goto tr142
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st76
			}
		case ( s)[p] >= 9:
			goto tr141
		}
		goto st0
tr49:
//line parse.rl:361

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
        
	goto st77
tr145:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:76

            if m>30 || m <0 {
                return nt,  fmt.Errorf("invalid day of month %d", m)
            }
            nt.dom |= 1<<(m-1)// again we zero index day
        
	goto st77
tr149:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:361

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
        
	goto st77
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
//line parse.go:3831
		if ( s)[p] == 42 {
			goto tr46
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr47
		}
		goto st0
tr47:
//line parse.rl:29

            mark = p;
        
	goto st78
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
//line parse.go:3850
		switch ( s)[p] {
		case 32:
			goto tr144
		case 44:
			goto tr145
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st78
			}
		case ( s)[p] >= 9:
			goto tr144
		}
		goto st0
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr147
		}
		goto st0
tr147:
//line parse.rl:29

            mark = p;
        
	goto st80
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
//line parse.go:3886
		switch ( s)[p] {
		case 32:
			goto tr148
		case 44:
			goto tr149
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st80
			}
		case ( s)[p] >= 9:
			goto tr148
		}
		goto st0
tr43:
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
        
	goto st81
tr152:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:69

            if m>24 || m <0 {
                return nt, fmt.Errorf("invalid hour of day %d", m)
            }
            nt.hour |= 1<<m
        
	goto st81
tr156:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
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
        
	goto st81
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
//line parse.go:4004
		if ( s)[p] == 42 {
			goto tr40
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr41
		}
		goto st0
tr41:
//line parse.rl:29

            mark = p;
        
	goto st82
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
//line parse.go:4023
		switch ( s)[p] {
		case 32:
			goto tr151
		case 44:
			goto tr152
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st82
			}
		case ( s)[p] >= 9:
			goto tr151
		}
		goto st0
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr154
		}
		goto st0
tr154:
//line parse.rl:29

            mark = p;
        
	goto st84
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
//line parse.go:4059
		switch ( s)[p] {
		case 32:
			goto tr155
		case 44:
			goto tr156
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st84
			}
		case ( s)[p] >= 9:
			goto tr155
		}
		goto st0
tr37:
//line parse.rl:441
 start=0;end=59;m=1;d=1; 
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
        
	goto st85
tr159:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:441
 start=m; end=0;d=0;
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
        
	goto st85
tr165:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:441
 end=m; d=1;
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
        
	goto st85
tr170:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:447
d=m
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
        
	goto st85
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
//line parse.go:4468
		if ( s)[p] == 42 {
			goto tr34
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr35
		}
		goto st0
tr35:
//line parse.rl:441
d=0;
//line parse.rl:29

            mark = p;
        
	goto st86
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
//line parse.go:4489
		switch ( s)[p] {
		case 32:
			goto tr158
		case 44:
			goto tr159
		case 45:
			goto tr160
		case 47:
			goto tr161
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st86
			}
		case ( s)[p] >= 9:
			goto tr158
		}
		goto st0
tr160:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:441
 start=m; end=0;d=0;
	goto st87
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
//line parse.go:4526
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr163
		}
		goto st0
tr163:
//line parse.rl:29

            mark = p;
        
	goto st88
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
//line parse.go:4542
		switch ( s)[p] {
		case 32:
			goto tr164
		case 44:
			goto tr165
		case 47:
			goto tr166
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st88
			}
		case ( s)[p] >= 9:
			goto tr164
		}
		goto st0
tr38:
//line parse.rl:441
 start=0;end=59;m=1;d=1; 
	goto st89
tr161:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:441
 start=m; end=0;d=0;
	goto st89
tr166:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:441
 end=m; d=1;
	goto st89
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
//line parse.go:4593
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr168
		}
		goto st0
tr168:
//line parse.rl:29

            mark = p;
        
	goto st90
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
//line parse.go:4609
		switch ( s)[p] {
		case 32:
			goto tr169
		case 44:
			goto tr170
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st90
			}
		case ( s)[p] >= 9:
			goto tr169
		}
		goto st0
tr31:
//line parse.rl:441
 start=0;end=59;m=1;d=1; 
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
        
	goto st91
tr175:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:441
 start=m; end=0;d=0;
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
        
	goto st91
tr181:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:441
 end=m; d=1;
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
        
	goto st91
tr186:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:442
d=m
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
        
	goto st91
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
//line parse.go:5018
		if ( s)[p] == 42 {
			goto tr172
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr173
		}
		goto st0
tr29:
//line parse.rl:481
panic("")
//line parse.rl:441
d=0;
//line parse.rl:29

            mark = p;
        
	goto st92
tr173:
//line parse.rl:441
d=0;
//line parse.rl:29

            mark = p;
        
	goto st92
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
//line parse.go:5049
		switch ( s)[p] {
		case 32:
			goto tr174
		case 44:
			goto tr175
		case 45:
			goto tr176
		case 47:
			goto tr177
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st92
			}
		case ( s)[p] >= 9:
			goto tr174
		}
		goto st0
tr176:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:441
 start=m; end=0;d=0;
	goto st93
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
//line parse.go:5086
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr179
		}
		goto st0
tr179:
//line parse.rl:29

            mark = p;
        
	goto st94
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
//line parse.go:5102
		switch ( s)[p] {
		case 32:
			goto tr180
		case 44:
			goto tr181
		case 47:
			goto tr182
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st94
			}
		case ( s)[p] >= 9:
			goto tr180
		}
		goto st0
tr32:
//line parse.rl:441
 start=0;end=59;m=1;d=1; 
	goto st95
tr177:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:441
 start=m; end=0;d=0;
	goto st95
tr182:
//line parse.rl:416

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:441
 end=m; d=1;
	goto st95
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
//line parse.go:5153
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr184
		}
		goto st0
tr184:
//line parse.rl:29

            mark = p;
        
	goto st96
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
//line parse.go:5169
		switch ( s)[p] {
		case 32:
			goto tr185
		case 44:
			goto tr186
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st96
			}
		case ( s)[p] >= 9:
			goto tr185
		}
		goto st0
	st_out:
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof45: cs = 45; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof
	_test_eof53: cs = 53; goto _test_eof
	_test_eof54: cs = 54; goto _test_eof
	_test_eof55: cs = 55; goto _test_eof
	_test_eof56: cs = 56; goto _test_eof
	_test_eof57: cs = 57; goto _test_eof
	_test_eof58: cs = 58; goto _test_eof
	_test_eof59: cs = 59; goto _test_eof
	_test_eof60: cs = 60; goto _test_eof
	_test_eof61: cs = 61; goto _test_eof
	_test_eof62: cs = 62; goto _test_eof
	_test_eof63: cs = 63; goto _test_eof
	_test_eof64: cs = 64; goto _test_eof
	_test_eof65: cs = 65; goto _test_eof
	_test_eof66: cs = 66; goto _test_eof
	_test_eof67: cs = 67; goto _test_eof
	_test_eof68: cs = 68; goto _test_eof
	_test_eof69: cs = 69; goto _test_eof
	_test_eof70: cs = 70; goto _test_eof
	_test_eof71: cs = 71; goto _test_eof
	_test_eof72: cs = 72; goto _test_eof
	_test_eof73: cs = 73; goto _test_eof
	_test_eof74: cs = 74; goto _test_eof
	_test_eof75: cs = 75; goto _test_eof
	_test_eof76: cs = 76; goto _test_eof
	_test_eof77: cs = 77; goto _test_eof
	_test_eof78: cs = 78; goto _test_eof
	_test_eof79: cs = 79; goto _test_eof
	_test_eof80: cs = 80; goto _test_eof
	_test_eof81: cs = 81; goto _test_eof
	_test_eof82: cs = 82; goto _test_eof
	_test_eof83: cs = 83; goto _test_eof
	_test_eof84: cs = 84; goto _test_eof
	_test_eof85: cs = 85; goto _test_eof
	_test_eof86: cs = 86; goto _test_eof
	_test_eof87: cs = 87; goto _test_eof
	_test_eof88: cs = 88; goto _test_eof
	_test_eof89: cs = 89; goto _test_eof
	_test_eof90: cs = 90; goto _test_eof
	_test_eof91: cs = 91; goto _test_eof
	_test_eof92: cs = 92; goto _test_eof
	_test_eof93: cs = 93; goto _test_eof
	_test_eof94: cs = 94; goto _test_eof
	_test_eof95: cs = 95; goto _test_eof
	_test_eof96: cs = 96; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17:
//line parse.rl:463
 
                if p!=pe{
                    fmt.Println(p,pe, backtrack)
                }
                if p==pe{
                    if backtrack == 7{
                        fmt.Println("here case 7")
                        p=mark
                        fmt.Println("here case 7", p, pe, mark)
                        //fhold;
                        fmt.Println("here case 7.b")
                        //cs = fentry(sevenPos);
                        p = ( p) - 1

                        cs = 18;
                        //fgoto *fentry(sevenPos);
                    }
                }
            
//line parse.go:5307
		}
	}

	_out: {}
	}

//line parse.rl:485
    // check current month
    // check the next year, unless feb 1<<29th is allowed then check the next 1<<8 years
    // if year isn't specified accept any year
    fmt.Println("nt.year", nt.year)
    fmt.Println("vars", cs, p, pe)
    fmt.Printf("in parser: backtrack %d\n, year %b, %b \nmonth %b, \ndom %b, \ndow %b, \nhour %b, \nmin %b, \ns %b\n", backtrack, nt.year.low, nt.year.high, nt.month, nt.dom, nt.dow, nt.hour, nt.minute, nt.second)
    if nt.minute==0||nt.hour==0||nt.dom==0||nt.month==0||nt.dow==0||nt.year.isZero() {
        return nt, errors.New("failed to parse cron string")
    }

    if nt.year.high==0 && nt.year.low==0 {
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

