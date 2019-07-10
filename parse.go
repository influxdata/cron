
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
const parse_start int = 211
const parse_first_final int = 211
const parse_error int = 0

const parse_en_sevenPos int = 11
const parse_en_sixPos int = 85
const parse_en_fivePos int = 152
const parse_en_main int = 211


//line parse.rl:14

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
    
//line parse.go:53
	{
	cs = parse_start
	top = 0
	ts = 0
	te = 0
	act = 0
	}

//line parse.rl:36
    //m,h := 1<<0,1<<0
    
//line parse.rl:496

    //
    
//line parse.go:69
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 211:
		goto st_case_211
	case 0:
		goto st_case_0
	case 212:
		goto st_case_212
	case 1:
		goto st_case_1
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
	case 213:
		goto st_case_213
	case 9:
		goto st_case_9
	case 214:
		goto st_case_214
	case 10:
		goto st_case_10
	case 215:
		goto st_case_215
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
	case 216:
		goto st_case_216
	case 217:
		goto st_case_217
	case 24:
		goto st_case_24
	case 218:
		goto st_case_218
	case 25:
		goto st_case_25
	case 219:
		goto st_case_219
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
	case 220:
		goto st_case_220
	case 221:
		goto st_case_221
	case 96:
		goto st_case_96
	case 222:
		goto st_case_222
	case 97:
		goto st_case_97
	case 98:
		goto st_case_98
	case 223:
		goto st_case_223
	case 99:
		goto st_case_99
	case 100:
		goto st_case_100
	case 101:
		goto st_case_101
	case 102:
		goto st_case_102
	case 103:
		goto st_case_103
	case 104:
		goto st_case_104
	case 105:
		goto st_case_105
	case 106:
		goto st_case_106
	case 107:
		goto st_case_107
	case 108:
		goto st_case_108
	case 109:
		goto st_case_109
	case 224:
		goto st_case_224
	case 110:
		goto st_case_110
	case 111:
		goto st_case_111
	case 112:
		goto st_case_112
	case 113:
		goto st_case_113
	case 114:
		goto st_case_114
	case 115:
		goto st_case_115
	case 116:
		goto st_case_116
	case 117:
		goto st_case_117
	case 118:
		goto st_case_118
	case 119:
		goto st_case_119
	case 120:
		goto st_case_120
	case 121:
		goto st_case_121
	case 122:
		goto st_case_122
	case 123:
		goto st_case_123
	case 124:
		goto st_case_124
	case 125:
		goto st_case_125
	case 126:
		goto st_case_126
	case 127:
		goto st_case_127
	case 128:
		goto st_case_128
	case 129:
		goto st_case_129
	case 130:
		goto st_case_130
	case 131:
		goto st_case_131
	case 132:
		goto st_case_132
	case 133:
		goto st_case_133
	case 134:
		goto st_case_134
	case 135:
		goto st_case_135
	case 136:
		goto st_case_136
	case 137:
		goto st_case_137
	case 138:
		goto st_case_138
	case 139:
		goto st_case_139
	case 140:
		goto st_case_140
	case 141:
		goto st_case_141
	case 142:
		goto st_case_142
	case 143:
		goto st_case_143
	case 144:
		goto st_case_144
	case 145:
		goto st_case_145
	case 146:
		goto st_case_146
	case 147:
		goto st_case_147
	case 148:
		goto st_case_148
	case 149:
		goto st_case_149
	case 150:
		goto st_case_150
	case 151:
		goto st_case_151
	case 152:
		goto st_case_152
	case 153:
		goto st_case_153
	case 154:
		goto st_case_154
	case 155:
		goto st_case_155
	case 156:
		goto st_case_156
	case 157:
		goto st_case_157
	case 158:
		goto st_case_158
	case 159:
		goto st_case_159
	case 160:
		goto st_case_160
	case 225:
		goto st_case_225
	case 226:
		goto st_case_226
	case 161:
		goto st_case_161
	case 227:
		goto st_case_227
	case 162:
		goto st_case_162
	case 163:
		goto st_case_163
	case 228:
		goto st_case_228
	case 164:
		goto st_case_164
	case 165:
		goto st_case_165
	case 166:
		goto st_case_166
	case 167:
		goto st_case_167
	case 168:
		goto st_case_168
	case 169:
		goto st_case_169
	case 170:
		goto st_case_170
	case 171:
		goto st_case_171
	case 172:
		goto st_case_172
	case 173:
		goto st_case_173
	case 174:
		goto st_case_174
	case 229:
		goto st_case_229
	case 175:
		goto st_case_175
	case 176:
		goto st_case_176
	case 177:
		goto st_case_177
	case 178:
		goto st_case_178
	case 179:
		goto st_case_179
	case 180:
		goto st_case_180
	case 181:
		goto st_case_181
	case 182:
		goto st_case_182
	case 183:
		goto st_case_183
	case 184:
		goto st_case_184
	case 185:
		goto st_case_185
	case 186:
		goto st_case_186
	case 187:
		goto st_case_187
	case 188:
		goto st_case_188
	case 189:
		goto st_case_189
	case 190:
		goto st_case_190
	case 191:
		goto st_case_191
	case 192:
		goto st_case_192
	case 193:
		goto st_case_193
	case 194:
		goto st_case_194
	case 195:
		goto st_case_195
	case 196:
		goto st_case_196
	case 197:
		goto st_case_197
	case 198:
		goto st_case_198
	case 199:
		goto st_case_199
	case 200:
		goto st_case_200
	case 201:
		goto st_case_201
	case 202:
		goto st_case_202
	case 203:
		goto st_case_203
	case 204:
		goto st_case_204
	case 205:
		goto st_case_205
	case 206:
		goto st_case_206
	case 207:
		goto st_case_207
	case 208:
		goto st_case_208
	case 209:
		goto st_case_209
	case 210:
		goto st_case_210
	}
	goto st_out
tr10:
//line parse.rl:476
p = (te) - 1
{
                // set seconds to 0 second of minute
                nt.second=1
                // set year to be permissive
                nt.year.high=^uint64(0)
                nt.year.low=^uint64(0)
                p = ( mark) - 1

                {stack[top] = 211; top++; goto st152 }
                }
	goto st211
tr13:
//line parse.rl:486
p = (te) - 1
{
                // set year to be permissive
                nt.year.high=^uint64(0)
                nt.year.low=^uint64(0)
                p = ( mark) - 1

                {goto st85 }
                }
	goto st211
tr412:
//line parse.rl:38
te = p
p--
{
            mark = p;
        }
	goto st211
tr413:
//line parse.rl:476
te = p
p--
{
                // set seconds to 0 second of minute
                nt.second=1
                // set year to be permissive
                nt.year.high=^uint64(0)
                nt.year.low=^uint64(0)
                p = ( mark) - 1

                {stack[top] = 211; top++; goto st152 }
                }
	goto st211
tr414:
//line parse.rl:486
te = p
p--
{
                // set year to be permissive
                nt.year.high=^uint64(0)
                nt.year.low=^uint64(0)
                p = ( mark) - 1

                {goto st85 }
                }
	goto st211
tr415:
//line parse.rl:494
te = p
p--
{p = ( mark) - 1
{stack[top] = 211; top++; goto st11 }}
	goto st211
	st211:
//line NONE:1
ts = 0

		if p++; p == pe {
			goto _test_eof211
		}
	st_case_211:
//line NONE:1
ts = p

//line parse.go:617
		switch ( s)[p] {
		case 32:
			goto st212
		case 42:
			goto tr411
		}
		switch {
		case ( s)[p] < 47:
			switch {
			case ( s)[p] > 13:
				if 44 <= ( s)[p] && ( s)[p] <= 45 {
					goto tr411
				}
			case ( s)[p] >= 9:
				goto st212
			}
		case ( s)[p] > 57:
			switch {
			case ( s)[p] > 90:
				if 97 <= ( s)[p] && ( s)[p] <= 122 {
					goto tr411
				}
			case ( s)[p] >= 65:
				goto tr411
			}
		default:
			goto tr411
		}
		goto st0
st_case_0:
	st0:
		cs = 0
		goto _out
	st212:
		if p++; p == pe {
			goto _test_eof212
		}
	st_case_212:
		if ( s)[p] == 32 {
			goto st212
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto st212
		}
		goto tr412
tr411:
//line parse.rl:38

            mark = p;
        
	goto st1
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
//line parse.go:674
		switch ( s)[p] {
		case 32:
			goto st2
		case 42:
			goto st1
		}
		switch {
		case ( s)[p] < 47:
			switch {
			case ( s)[p] > 13:
				if 44 <= ( s)[p] && ( s)[p] <= 45 {
					goto st1
				}
			case ( s)[p] >= 9:
				goto st2
			}
		case ( s)[p] > 57:
			switch {
			case ( s)[p] > 90:
				if 97 <= ( s)[p] && ( s)[p] <= 122 {
					goto st1
				}
			case ( s)[p] >= 65:
				goto st1
			}
		default:
			goto st1
		}
		goto st0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch ( s)[p] {
		case 32:
			goto st2
		case 42:
			goto st3
		}
		switch {
		case ( s)[p] < 47:
			switch {
			case ( s)[p] > 13:
				if 44 <= ( s)[p] && ( s)[p] <= 45 {
					goto st3
				}
			case ( s)[p] >= 9:
				goto st2
			}
		case ( s)[p] > 57:
			switch {
			case ( s)[p] > 90:
				if 97 <= ( s)[p] && ( s)[p] <= 122 {
					goto st3
				}
			case ( s)[p] >= 65:
				goto st3
			}
		default:
			goto st3
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch ( s)[p] {
		case 32:
			goto st4
		case 42:
			goto st3
		}
		switch {
		case ( s)[p] < 47:
			switch {
			case ( s)[p] > 13:
				if 44 <= ( s)[p] && ( s)[p] <= 45 {
					goto st3
				}
			case ( s)[p] >= 9:
				goto st4
			}
		case ( s)[p] > 57:
			switch {
			case ( s)[p] > 90:
				if 97 <= ( s)[p] && ( s)[p] <= 122 {
					goto st3
				}
			case ( s)[p] >= 65:
				goto st3
			}
		default:
			goto st3
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch ( s)[p] {
		case 32:
			goto st4
		case 42:
			goto st5
		}
		switch {
		case ( s)[p] < 47:
			switch {
			case ( s)[p] > 13:
				if 44 <= ( s)[p] && ( s)[p] <= 45 {
					goto st5
				}
			case ( s)[p] >= 9:
				goto st4
			}
		case ( s)[p] > 57:
			switch {
			case ( s)[p] > 90:
				if 97 <= ( s)[p] && ( s)[p] <= 122 {
					goto st5
				}
			case ( s)[p] >= 65:
				goto st5
			}
		default:
			goto st5
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch ( s)[p] {
		case 32:
			goto st6
		case 42:
			goto st5
		}
		switch {
		case ( s)[p] < 47:
			switch {
			case ( s)[p] > 13:
				if 44 <= ( s)[p] && ( s)[p] <= 45 {
					goto st5
				}
			case ( s)[p] >= 9:
				goto st6
			}
		case ( s)[p] > 57:
			switch {
			case ( s)[p] > 90:
				if 97 <= ( s)[p] && ( s)[p] <= 122 {
					goto st5
				}
			case ( s)[p] >= 65:
				goto st5
			}
		default:
			goto st5
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch ( s)[p] {
		case 32:
			goto st6
		case 42:
			goto st7
		}
		switch {
		case ( s)[p] < 47:
			switch {
			case ( s)[p] > 13:
				if 44 <= ( s)[p] && ( s)[p] <= 45 {
					goto st7
				}
			case ( s)[p] >= 9:
				goto st6
			}
		case ( s)[p] > 57:
			switch {
			case ( s)[p] > 90:
				if 97 <= ( s)[p] && ( s)[p] <= 122 {
					goto st7
				}
			case ( s)[p] >= 65:
				goto st7
			}
		default:
			goto st7
		}
		goto st0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch ( s)[p] {
		case 32:
			goto st8
		case 42:
			goto st7
		}
		switch {
		case ( s)[p] < 47:
			switch {
			case ( s)[p] > 13:
				if 44 <= ( s)[p] && ( s)[p] <= 45 {
					goto st7
				}
			case ( s)[p] >= 9:
				goto st8
			}
		case ( s)[p] > 57:
			switch {
			case ( s)[p] > 90:
				if 97 <= ( s)[p] && ( s)[p] <= 122 {
					goto st7
				}
			case ( s)[p] >= 65:
				goto st7
			}
		default:
			goto st7
		}
		goto st0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		switch ( s)[p] {
		case 32:
			goto st8
		case 42:
			goto tr9
		}
		switch {
		case ( s)[p] < 47:
			switch {
			case ( s)[p] > 13:
				if 44 <= ( s)[p] && ( s)[p] <= 45 {
					goto tr9
				}
			case ( s)[p] >= 9:
				goto st8
			}
		case ( s)[p] > 57:
			switch {
			case ( s)[p] > 90:
				if 97 <= ( s)[p] && ( s)[p] <= 122 {
					goto tr9
				}
			case ( s)[p] >= 65:
				goto tr9
			}
		default:
			goto tr9
		}
		goto st0
tr9:
//line NONE:1
te = p+1

	goto st213
	st213:
		if p++; p == pe {
			goto _test_eof213
		}
	st_case_213:
//line parse.go:952
		switch ( s)[p] {
		case 32:
			goto st9
		case 42:
			goto tr9
		}
		switch {
		case ( s)[p] < 47:
			switch {
			case ( s)[p] > 13:
				if 44 <= ( s)[p] && ( s)[p] <= 45 {
					goto tr9
				}
			case ( s)[p] >= 9:
				goto st9
			}
		case ( s)[p] > 57:
			switch {
			case ( s)[p] > 90:
				if 97 <= ( s)[p] && ( s)[p] <= 122 {
					goto tr9
				}
			case ( s)[p] >= 65:
				goto tr9
			}
		default:
			goto tr9
		}
		goto tr413
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		switch ( s)[p] {
		case 32:
			goto st9
		case 42:
			goto tr12
		}
		switch {
		case ( s)[p] < 47:
			switch {
			case ( s)[p] > 13:
				if 44 <= ( s)[p] && ( s)[p] <= 45 {
					goto tr12
				}
			case ( s)[p] >= 9:
				goto st9
			}
		case ( s)[p] > 57:
			switch {
			case ( s)[p] > 90:
				if 97 <= ( s)[p] && ( s)[p] <= 122 {
					goto tr12
				}
			case ( s)[p] >= 65:
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr10
tr12:
//line NONE:1
te = p+1

	goto st214
	st214:
		if p++; p == pe {
			goto _test_eof214
		}
	st_case_214:
//line parse.go:1026
		switch ( s)[p] {
		case 32:
			goto st10
		case 42:
			goto tr12
		}
		switch {
		case ( s)[p] < 47:
			switch {
			case ( s)[p] > 13:
				if 44 <= ( s)[p] && ( s)[p] <= 45 {
					goto tr12
				}
			case ( s)[p] >= 9:
				goto st10
			}
		case ( s)[p] > 57:
			switch {
			case ( s)[p] > 90:
				if 97 <= ( s)[p] && ( s)[p] <= 122 {
					goto tr12
				}
			case ( s)[p] >= 65:
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr414
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch ( s)[p] {
		case 32:
			goto st10
		case 42:
			goto st215
		}
		switch {
		case ( s)[p] < 47:
			switch {
			case ( s)[p] > 13:
				if 44 <= ( s)[p] && ( s)[p] <= 45 {
					goto st215
				}
			case ( s)[p] >= 9:
				goto st10
			}
		case ( s)[p] > 57:
			switch {
			case ( s)[p] > 90:
				if 97 <= ( s)[p] && ( s)[p] <= 122 {
					goto st215
				}
			case ( s)[p] >= 65:
				goto st215
			}
		default:
			goto st215
		}
		goto tr13
	st215:
		if p++; p == pe {
			goto _test_eof215
		}
	st_case_215:
		if ( s)[p] == 42 {
			goto st215
		}
		switch {
		case ( s)[p] < 47:
			if 44 <= ( s)[p] && ( s)[p] <= 45 {
				goto st215
			}
		case ( s)[p] > 57:
			switch {
			case ( s)[p] > 90:
				if 97 <= ( s)[p] && ( s)[p] <= 122 {
					goto st215
				}
			case ( s)[p] >= 65:
				goto st215
			}
		default:
			goto st215
		}
		goto tr415
tr19:
//line parse.rl:450
 start=0;end=59;m=1;d=1; 
//line parse.rl:99

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
        
	goto st11
tr152:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:451
d=m
//line parse.rl:99

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
        
	goto st11
tr155:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 start=m; end=0;d=0;
//line parse.rl:99

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
        
	goto st11
tr161:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 end=m; d=1;
//line parse.rl:99

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
        
	goto st11
	st11:
//line NONE:1
ts = 0

		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line parse.go:1512
		if ( s)[p] == 42 {
			goto tr16
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr17
		}
		goto st0
tr16:
//line parse.rl:450
d=0;
	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line parse.go:1529
		switch ( s)[p] {
		case 32:
			goto tr18
		case 44:
			goto tr19
		case 47:
			goto tr20
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr18
		}
		goto st0
tr18:
//line parse.rl:450
 start=0;end=59;m=1;d=1; 
//line parse.rl:99

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
        
	goto st13
tr151:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:451
d=m
//line parse.rl:99

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
        
	goto st13
tr154:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 start=m; end=0;d=0;
//line parse.rl:99

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
        
	goto st13
tr160:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 end=m; d=1;
//line parse.rl:99

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
        
	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line parse.go:1935
		switch ( s)[p] {
		case 32:
			goto st13
		case 42:
			goto tr22
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr23
			}
		case ( s)[p] >= 9:
			goto st13
		}
		goto st0
tr22:
//line parse.rl:450
d=0;
	goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line parse.go:1960
		switch ( s)[p] {
		case 32:
			goto tr24
		case 44:
			goto tr25
		case 47:
			goto tr26
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr24
		}
		goto st0
tr24:
//line parse.rl:450
 start=0;end=59;m=1;d=1; 
//line parse.rl:186

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
        
	goto st15
tr136:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 start=m; end=0;d=0;
//line parse.rl:186

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
        
	goto st15
tr142:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 end=m; d=1;
//line parse.rl:186

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
        
	goto st15
tr147:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:456
d=m
//line parse.rl:186

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
        
	goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line parse.go:2366
		switch ( s)[p] {
		case 32:
			goto st15
		case 42:
			goto tr28
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr29
			}
		case ( s)[p] >= 9:
			goto st15
		}
		goto st0
tr28:
//line parse.rl:446
m=1
	goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line parse.go:2391
		switch ( s)[p] {
		case 32:
			goto tr30
		case 44:
			goto tr31
		case 47:
			goto st72
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr30
		}
		goto st0
tr30:
//line parse.rl:273

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
        
	goto st17
tr129:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:78

            if m>24 || m <0 {
                return nt, fmt.Errorf("invalid hour of day %d", m)
            }
            nt.hour |= 1<<m
        
	goto st17
tr133:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:273

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
        
	goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line parse.go:2506
		switch ( s)[p] {
		case 32:
			goto st17
		case 42:
			goto tr34
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr35
			}
		case ( s)[p] >= 9:
			goto st17
		}
		goto st0
tr34:
//line parse.rl:446
m=1
	goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line parse.go:2531
		switch ( s)[p] {
		case 32:
			goto tr36
		case 44:
			goto tr37
		case 47:
			goto st68
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr36
		}
		goto st0
tr36:
//line parse.rl:370

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
        
	goto st19
tr122:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:85

            if m>30 || m <0 {
                return nt,  fmt.Errorf("invalid day of month %d", m)
            }
            nt.dom |= 1<<(m-1)// again we zero index day
        
	goto st19
tr126:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:370

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
        
	goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line parse.go:2662
		switch ( s)[p] {
		case 32:
			goto st19
		case 42:
			goto tr40
		case 65:
			goto st45
		case 68:
			goto st49
		case 70:
			goto st51
		case 74:
			goto st53
		case 77:
			goto st56
		case 78:
			goto st58
		case 79:
			goto st60
		case 83:
			goto st62
		case 97:
			goto st45
		case 100:
			goto st49
		case 102:
			goto st51
		case 106:
			goto st53
		case 109:
			goto st56
		case 110:
			goto st58
		case 111:
			goto st60
		case 115:
			goto st62
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr41
			}
		case ( s)[p] >= 9:
			goto st19
		}
		goto st0
tr40:
//line parse.rl:446
m=1
	goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line parse.go:2719
		switch ( s)[p] {
		case 32:
			goto tr50
		case 44:
			goto tr51
		case 47:
			goto st64
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr50
		}
		goto st0
tr50:
//line parse.rl:307

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
        
	goto st21
tr91:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:50

            if m>12 || m<1{
                return nt, fmt.Errorf("invalid month %d", m)
            }
            nt.month |= 1<<(m-1) // we zero index months
        
	goto st21
tr97:
//line parse.rl:50

            if m>12 || m<1{
                return nt, fmt.Errorf("invalid month %d", m)
            }
            nt.month |= 1<<(m-1) // we zero index months
        
	goto st21
tr119:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:307

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
        
	goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line parse.go:2821
		switch ( s)[p] {
		case 32:
			goto st21
		case 42:
			goto tr54
		case 70:
			goto st28
		case 77:
			goto st31
		case 83:
			goto st33
		case 84:
			goto st36
		case 87:
			goto st39
		case 102:
			goto st28
		case 109:
			goto st31
		case 115:
			goto st33
		case 116:
			goto st36
		case 119:
			goto st39
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr55
			}
		case ( s)[p] >= 9:
			goto st21
		}
		goto st0
tr54:
//line parse.rl:446
m=1
	goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line parse.go:2866
		switch ( s)[p] {
		case 32:
			goto tr61
		case 44:
			goto tr62
		case 47:
			goto st41
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr61
		}
		goto st0
tr61:
//line parse.rl:330

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
        
	goto st23
tr68:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:64

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
        
	goto st23
tr73:
//line parse.rl:64

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
        
	goto st23
tr88:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:330

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
        
	goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line parse.go:2982
		switch ( s)[p] {
		case 32:
			goto st23
		case 42:
			goto tr65
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr66
			}
		case ( s)[p] >= 9:
			goto st23
		}
		goto st0
tr65:
//line parse.rl:446
m=1
	goto st216
	st216:
		if p++; p == pe {
			goto _test_eof216
		}
	st_case_216:
//line parse.go:3007
		switch ( s)[p] {
		case 32:
			goto tr416
		case 44:
			goto tr417
		case 47:
			goto st25
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr416
		}
		goto st0
tr416:
//line parse.rl:353

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
        
	goto st217
tr420:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:92

            if m<1970 || m>2099{
                return nt, errors.New("only years on interval [1970, 2099] are supported")
            }
            nt.year.set(int(m))
        
	goto st217
tr423:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:353

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
        
	goto st217
	st217:
		if p++; p == pe {
			goto _test_eof217
		}
	st_case_217:
//line parse.go:3082
		if ( s)[p] == 32 {
			goto st217
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto st217
		}
		goto st0
tr417:
//line parse.rl:353

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
        
	goto st24
tr421:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:92

            if m<1970 || m>2099{
                return nt, errors.New("only years on interval [1970, 2099] are supported")
            }
            nt.year.set(int(m))
        
	goto st24
tr424:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:353

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
        
	goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line parse.go:3152
		if ( s)[p] == 42 {
			goto tr65
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr66
		}
		goto st0
tr66:
//line parse.rl:38

            mark = p;
        
	goto st218
	st218:
		if p++; p == pe {
			goto _test_eof218
		}
	st_case_218:
//line parse.go:3171
		switch ( s)[p] {
		case 32:
			goto tr420
		case 44:
			goto tr421
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st218
			}
		case ( s)[p] >= 9:
			goto tr420
		}
		goto st0
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr67
		}
		goto st0
tr67:
//line parse.rl:38

            mark = p;
        
	goto st219
	st219:
		if p++; p == pe {
			goto _test_eof219
		}
	st_case_219:
//line parse.go:3207
		switch ( s)[p] {
		case 32:
			goto tr423
		case 44:
			goto tr424
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st219
			}
		case ( s)[p] >= 9:
			goto tr423
		}
		goto st0
tr62:
//line parse.rl:330

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
        
	goto st26
tr69:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:64

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
        
	goto st26
tr74:
//line parse.rl:64

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
        
	goto st26
tr89:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:330

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
        
	goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line parse.go:3326
		switch ( s)[p] {
		case 42:
			goto tr54
		case 70:
			goto st28
		case 77:
			goto st31
		case 83:
			goto st33
		case 84:
			goto st36
		case 87:
			goto st39
		case 102:
			goto st28
		case 109:
			goto st31
		case 115:
			goto st33
		case 116:
			goto st36
		case 119:
			goto st39
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr55
		}
		goto st0
tr55:
//line parse.rl:38

            mark = p;
        
	goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
//line parse.go:3366
		switch ( s)[p] {
		case 32:
			goto tr68
		case 44:
			goto tr69
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st27
			}
		case ( s)[p] >= 9:
			goto tr68
		}
		goto st0
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		switch ( s)[p] {
		case 82:
			goto st29
		case 114:
			goto st29
		}
		goto st0
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		switch ( s)[p] {
		case 73:
			goto tr72
		case 105:
			goto tr72
		}
		goto st0
tr72:
//line parse.rl:437
m=5
	goto st30
tr76:
//line parse.rl:437
m=1
	goto st30
tr79:
//line parse.rl:437
m=6
	goto st30
tr80:
//line parse.rl:437
m=0
	goto st30
tr83:
//line parse.rl:437
m=4
	goto st30
tr84:
//line parse.rl:437
m=2
	goto st30
tr86:
//line parse.rl:437
m=3
	goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
//line parse.go:3439
		switch ( s)[p] {
		case 32:
			goto tr73
		case 44:
			goto tr74
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr73
		}
		goto st0
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		switch ( s)[p] {
		case 79:
			goto st32
		case 111:
			goto st32
		}
		goto st0
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		switch ( s)[p] {
		case 78:
			goto tr76
		case 110:
			goto tr76
		}
		goto st0
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		switch ( s)[p] {
		case 65:
			goto st34
		case 85:
			goto st35
		case 97:
			goto st34
		case 117:
			goto st35
		}
		goto st0
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		switch ( s)[p] {
		case 84:
			goto tr79
		case 116:
			goto tr79
		}
		goto st0
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		switch ( s)[p] {
		case 78:
			goto tr80
		case 110:
			goto tr80
		}
		goto st0
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		switch ( s)[p] {
		case 72:
			goto st37
		case 85:
			goto st38
		case 104:
			goto st37
		case 117:
			goto st38
		}
		goto st0
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		switch ( s)[p] {
		case 85:
			goto tr83
		case 117:
			goto tr83
		}
		goto st0
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		switch ( s)[p] {
		case 69:
			goto tr84
		case 101:
			goto tr84
		}
		goto st0
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		switch ( s)[p] {
		case 69:
			goto st40
		case 101:
			goto st40
		}
		goto st0
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		switch ( s)[p] {
		case 68:
			goto tr86
		case 100:
			goto tr86
		}
		goto st0
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr87
		}
		goto st0
tr87:
//line parse.rl:38

            mark = p;
        
	goto st42
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
//line parse.go:3598
		switch ( s)[p] {
		case 32:
			goto tr88
		case 44:
			goto tr89
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st42
			}
		case ( s)[p] >= 9:
			goto tr88
		}
		goto st0
tr51:
//line parse.rl:307

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
        
	goto st43
tr92:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:50

            if m>12 || m<1{
                return nt, fmt.Errorf("invalid month %d", m)
            }
            nt.month |= 1<<(m-1) // we zero index months
        
	goto st43
tr98:
//line parse.rl:50

            if m>12 || m<1{
                return nt, fmt.Errorf("invalid month %d", m)
            }
            nt.month |= 1<<(m-1) // we zero index months
        
	goto st43
tr120:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:307

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
        
	goto st43
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
//line parse.go:3703
		switch ( s)[p] {
		case 42:
			goto tr40
		case 65:
			goto st45
		case 68:
			goto st49
		case 70:
			goto st51
		case 74:
			goto st53
		case 77:
			goto st56
		case 78:
			goto st58
		case 79:
			goto st60
		case 83:
			goto st62
		case 97:
			goto st45
		case 100:
			goto st49
		case 102:
			goto st51
		case 106:
			goto st53
		case 109:
			goto st56
		case 110:
			goto st58
		case 111:
			goto st60
		case 115:
			goto st62
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr41
		}
		goto st0
tr41:
//line parse.rl:38

            mark = p;
        
	goto st44
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
//line parse.go:3755
		switch ( s)[p] {
		case 32:
			goto tr91
		case 44:
			goto tr92
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st44
			}
		case ( s)[p] >= 9:
			goto tr91
		}
		goto st0
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		switch ( s)[p] {
		case 80:
			goto st46
		case 85:
			goto st48
		case 112:
			goto st46
		case 117:
			goto st48
		}
		goto st0
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		switch ( s)[p] {
		case 82:
			goto tr96
		case 114:
			goto tr96
		}
		goto st0
tr96:
//line parse.rl:438
m=4
	goto st47
tr99:
//line parse.rl:438
m=8
	goto st47
tr101:
//line parse.rl:438
m=12
	goto st47
tr103:
//line parse.rl:438
m=2
	goto st47
tr106:
//line parse.rl:438
m=1
	goto st47
tr107:
//line parse.rl:438
m=7
	goto st47
tr108:
//line parse.rl:438
m=6
	goto st47
tr110:
//line parse.rl:438
m=3
	goto st47
tr111:
//line parse.rl:438
m=5
	goto st47
tr113:
//line parse.rl:438
m=11
	goto st47
tr115:
//line parse.rl:438
m=10
	goto st47
tr117:
//line parse.rl:438
m=9
	goto st47
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
//line parse.go:3852
		switch ( s)[p] {
		case 32:
			goto tr97
		case 44:
			goto tr98
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr97
		}
		goto st0
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		switch ( s)[p] {
		case 71:
			goto tr99
		case 103:
			goto tr99
		}
		goto st0
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		switch ( s)[p] {
		case 69:
			goto st50
		case 101:
			goto st50
		}
		goto st0
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		switch ( s)[p] {
		case 67:
			goto tr101
		case 99:
			goto tr101
		}
		goto st0
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
		switch ( s)[p] {
		case 69:
			goto st52
		case 101:
			goto st52
		}
		goto st0
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
		switch ( s)[p] {
		case 66:
			goto tr103
		case 98:
			goto tr103
		}
		goto st0
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
		switch ( s)[p] {
		case 65:
			goto st54
		case 85:
			goto st55
		case 97:
			goto st54
		case 117:
			goto st55
		}
		goto st0
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		switch ( s)[p] {
		case 78:
			goto tr106
		case 110:
			goto tr106
		}
		goto st0
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
		switch ( s)[p] {
		case 76:
			goto tr107
		case 78:
			goto tr108
		case 108:
			goto tr107
		case 110:
			goto tr108
		}
		goto st0
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		switch ( s)[p] {
		case 65:
			goto st57
		case 97:
			goto st57
		}
		goto st0
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		switch ( s)[p] {
		case 82:
			goto tr110
		case 89:
			goto tr111
		case 114:
			goto tr110
		case 121:
			goto tr111
		}
		goto st0
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		switch ( s)[p] {
		case 79:
			goto st59
		case 111:
			goto st59
		}
		goto st0
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		switch ( s)[p] {
		case 86:
			goto tr113
		case 118:
			goto tr113
		}
		goto st0
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
		switch ( s)[p] {
		case 67:
			goto st61
		case 99:
			goto st61
		}
		goto st0
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		switch ( s)[p] {
		case 84:
			goto tr115
		case 116:
			goto tr115
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
		case 80:
			goto tr117
		case 112:
			goto tr117
		}
		goto st0
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr118
		}
		goto st0
tr118:
//line parse.rl:38

            mark = p;
        
	goto st65
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
//line parse.go:4087
		switch ( s)[p] {
		case 32:
			goto tr119
		case 44:
			goto tr120
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st65
			}
		case ( s)[p] >= 9:
			goto tr119
		}
		goto st0
tr37:
//line parse.rl:370

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
        
	goto st66
tr123:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:85

            if m>30 || m <0 {
                return nt,  fmt.Errorf("invalid day of month %d", m)
            }
            nt.dom |= 1<<(m-1)// again we zero index day
        
	goto st66
tr127:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:370

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
        
	goto st66
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
//line parse.go:4221
		if ( s)[p] == 42 {
			goto tr34
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr35
		}
		goto st0
tr35:
//line parse.rl:38

            mark = p;
        
	goto st67
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
//line parse.go:4240
		switch ( s)[p] {
		case 32:
			goto tr122
		case 44:
			goto tr123
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st67
			}
		case ( s)[p] >= 9:
			goto tr122
		}
		goto st0
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr125
		}
		goto st0
tr125:
//line parse.rl:38

            mark = p;
        
	goto st69
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
//line parse.go:4276
		switch ( s)[p] {
		case 32:
			goto tr126
		case 44:
			goto tr127
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st69
			}
		case ( s)[p] >= 9:
			goto tr126
		}
		goto st0
tr31:
//line parse.rl:273

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
        
	goto st70
tr130:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:78

            if m>24 || m <0 {
                return nt, fmt.Errorf("invalid hour of day %d", m)
            }
            nt.hour |= 1<<m
        
	goto st70
tr134:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:273

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
        
	goto st70
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
//line parse.go:4394
		if ( s)[p] == 42 {
			goto tr28
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr29
		}
		goto st0
tr29:
//line parse.rl:38

            mark = p;
        
	goto st71
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
//line parse.go:4413
		switch ( s)[p] {
		case 32:
			goto tr129
		case 44:
			goto tr130
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st71
			}
		case ( s)[p] >= 9:
			goto tr129
		}
		goto st0
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr132
		}
		goto st0
tr132:
//line parse.rl:38

            mark = p;
        
	goto st73
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
//line parse.go:4449
		switch ( s)[p] {
		case 32:
			goto tr133
		case 44:
			goto tr134
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st73
			}
		case ( s)[p] >= 9:
			goto tr133
		}
		goto st0
tr25:
//line parse.rl:450
 start=0;end=59;m=1;d=1; 
//line parse.rl:186

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
        
	goto st74
tr137:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 start=m; end=0;d=0;
//line parse.rl:186

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
        
	goto st74
tr143:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 end=m; d=1;
//line parse.rl:186

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
        
	goto st74
tr148:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:456
d=m
//line parse.rl:186

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
        
	goto st74
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
//line parse.go:4858
		if ( s)[p] == 42 {
			goto tr22
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr23
		}
		goto st0
tr23:
//line parse.rl:450
d=0;
//line parse.rl:38

            mark = p;
        
	goto st75
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
//line parse.go:4879
		switch ( s)[p] {
		case 32:
			goto tr136
		case 44:
			goto tr137
		case 45:
			goto tr138
		case 47:
			goto tr139
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st75
			}
		case ( s)[p] >= 9:
			goto tr136
		}
		goto st0
tr138:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 start=m; end=0;d=0;
	goto st76
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
//line parse.go:4916
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr141
		}
		goto st0
tr141:
//line parse.rl:38

            mark = p;
        
	goto st77
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
//line parse.go:4932
		switch ( s)[p] {
		case 32:
			goto tr142
		case 44:
			goto tr143
		case 47:
			goto tr144
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st77
			}
		case ( s)[p] >= 9:
			goto tr142
		}
		goto st0
tr26:
//line parse.rl:450
 start=0;end=59;m=1;d=1; 
	goto st78
tr139:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 start=m; end=0;d=0;
	goto st78
tr144:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 end=m; d=1;
	goto st78
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
//line parse.go:4983
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr146
		}
		goto st0
tr146:
//line parse.rl:38

            mark = p;
        
	goto st79
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
//line parse.go:4999
		switch ( s)[p] {
		case 32:
			goto tr147
		case 44:
			goto tr148
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st79
			}
		case ( s)[p] >= 9:
			goto tr147
		}
		goto st0
tr20:
//line parse.rl:450
 start=0;end=59;m=1;d=1; 
	goto st80
tr157:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 start=m; end=0;d=0;
	goto st80
tr162:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 end=m; d=1;
	goto st80
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
//line parse.go:5048
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr150
		}
		goto st0
tr150:
//line parse.rl:38

            mark = p;
        
	goto st81
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
//line parse.go:5064
		switch ( s)[p] {
		case 32:
			goto tr151
		case 44:
			goto tr152
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st81
			}
		case ( s)[p] >= 9:
			goto tr151
		}
		goto st0
tr17:
//line parse.rl:450
d=0;
//line parse.rl:38

            mark = p;
        
	goto st82
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
//line parse.go:5093
		switch ( s)[p] {
		case 32:
			goto tr154
		case 44:
			goto tr155
		case 45:
			goto tr156
		case 47:
			goto tr157
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st82
			}
		case ( s)[p] >= 9:
			goto tr154
		}
		goto st0
tr156:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 start=m; end=0;d=0;
	goto st83
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
//line parse.go:5130
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr159
		}
		goto st0
tr159:
//line parse.rl:38

            mark = p;
        
	goto st84
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
//line parse.go:5146
		switch ( s)[p] {
		case 32:
			goto tr160
		case 44:
			goto tr161
		case 47:
			goto tr162
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st84
			}
		case ( s)[p] >= 9:
			goto tr160
		}
		goto st0
tr167:
//line parse.rl:450
 start=0;end=59;m=1;d=1; 
//line parse.rl:99

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
        
	goto st85
tr285:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:451
d=m
//line parse.rl:99

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
        
	goto st85
tr288:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 start=m; end=0;d=0;
//line parse.rl:99

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
        
	goto st85
tr294:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 end=m; d=1;
//line parse.rl:99

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
        
	goto st85
	st85:
//line NONE:1
ts = 0

		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
//line parse.go:5560
		if ( s)[p] == 42 {
			goto tr164
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr165
		}
		goto st0
tr164:
//line parse.rl:450
d=0;
	goto st86
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
//line parse.go:5577
		switch ( s)[p] {
		case 32:
			goto tr166
		case 44:
			goto tr167
		case 47:
			goto tr168
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr166
		}
		goto st0
tr166:
//line parse.rl:450
 start=0;end=59;m=1;d=1; 
//line parse.rl:99

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
        
	goto st87
tr284:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:451
d=m
//line parse.rl:99

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
        
	goto st87
tr287:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 start=m; end=0;d=0;
//line parse.rl:99

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
        
	goto st87
tr293:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 end=m; d=1;
//line parse.rl:99

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
        
	goto st87
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
//line parse.go:5983
		switch ( s)[p] {
		case 32:
			goto st87
		case 42:
			goto tr170
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr171
			}
		case ( s)[p] >= 9:
			goto st87
		}
		goto st0
tr170:
//line parse.rl:450
d=0;
	goto st88
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
//line parse.go:6008
		switch ( s)[p] {
		case 32:
			goto tr172
		case 44:
			goto tr173
		case 47:
			goto tr174
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr172
		}
		goto st0
tr172:
//line parse.rl:450
 start=0;end=59;m=1;d=1; 
//line parse.rl:186

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
        
	goto st89
tr269:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 start=m; end=0;d=0;
//line parse.rl:186

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
        
	goto st89
tr275:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 end=m; d=1;
//line parse.rl:186

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
        
	goto st89
tr280:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:456
d=m
//line parse.rl:186

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
        
	goto st89
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
//line parse.go:6414
		switch ( s)[p] {
		case 32:
			goto st89
		case 42:
			goto tr176
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr177
			}
		case ( s)[p] >= 9:
			goto st89
		}
		goto st0
tr176:
//line parse.rl:446
m=1
	goto st90
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
//line parse.go:6439
		switch ( s)[p] {
		case 32:
			goto tr178
		case 44:
			goto tr179
		case 47:
			goto st139
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr178
		}
		goto st0
tr178:
//line parse.rl:273

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
        
	goto st91
tr262:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:78

            if m>24 || m <0 {
                return nt, fmt.Errorf("invalid hour of day %d", m)
            }
            nt.hour |= 1<<m
        
	goto st91
tr266:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:273

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
        
	goto st91
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
//line parse.go:6554
		switch ( s)[p] {
		case 32:
			goto st91
		case 42:
			goto tr182
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr183
			}
		case ( s)[p] >= 9:
			goto st91
		}
		goto st0
tr182:
//line parse.rl:446
m=1
	goto st92
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
//line parse.go:6579
		switch ( s)[p] {
		case 32:
			goto tr184
		case 44:
			goto tr185
		case 47:
			goto st135
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr184
		}
		goto st0
tr184:
//line parse.rl:370

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
        
	goto st93
tr255:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:85

            if m>30 || m <0 {
                return nt,  fmt.Errorf("invalid day of month %d", m)
            }
            nt.dom |= 1<<(m-1)// again we zero index day
        
	goto st93
tr259:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:370

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
        
	goto st93
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
//line parse.go:6710
		switch ( s)[p] {
		case 32:
			goto st93
		case 42:
			goto tr188
		case 65:
			goto st112
		case 68:
			goto st116
		case 70:
			goto st118
		case 74:
			goto st120
		case 77:
			goto st123
		case 78:
			goto st125
		case 79:
			goto st127
		case 83:
			goto st129
		case 97:
			goto st112
		case 100:
			goto st116
		case 102:
			goto st118
		case 106:
			goto st120
		case 109:
			goto st123
		case 110:
			goto st125
		case 111:
			goto st127
		case 115:
			goto st129
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr189
			}
		case ( s)[p] >= 9:
			goto st93
		}
		goto st0
tr188:
//line parse.rl:446
m=1
	goto st94
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
//line parse.go:6767
		switch ( s)[p] {
		case 32:
			goto tr198
		case 44:
			goto tr199
		case 47:
			goto st131
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr198
		}
		goto st0
tr198:
//line parse.rl:307

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
        
	goto st95
tr224:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:50

            if m>12 || m<1{
                return nt, fmt.Errorf("invalid month %d", m)
            }
            nt.month |= 1<<(m-1) // we zero index months
        
	goto st95
tr230:
//line parse.rl:50

            if m>12 || m<1{
                return nt, fmt.Errorf("invalid month %d", m)
            }
            nt.month |= 1<<(m-1) // we zero index months
        
	goto st95
tr252:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:307

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
        
	goto st95
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
//line parse.go:6869
		switch ( s)[p] {
		case 32:
			goto st95
		case 42:
			goto tr202
		case 70:
			goto st97
		case 77:
			goto st99
		case 83:
			goto st101
		case 84:
			goto st104
		case 87:
			goto st107
		case 102:
			goto st97
		case 109:
			goto st99
		case 115:
			goto st101
		case 116:
			goto st104
		case 119:
			goto st107
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr203
			}
		case ( s)[p] >= 9:
			goto st95
		}
		goto st0
tr202:
//line parse.rl:446
m=1
	goto st220
	st220:
		if p++; p == pe {
			goto _test_eof220
		}
	st_case_220:
//line parse.go:6914
		switch ( s)[p] {
		case 32:
			goto tr426
		case 44:
			goto tr427
		case 47:
			goto st109
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr426
		}
		goto st0
tr426:
//line parse.rl:330

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
        
	goto st221
tr430:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:64

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
        
	goto st221
tr433:
//line parse.rl:64

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
        
	goto st221
tr435:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:330

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
        
	goto st221
	st221:
		if p++; p == pe {
			goto _test_eof221
		}
	st_case_221:
//line parse.go:7030
		if ( s)[p] == 32 {
			goto st221
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto st221
		}
		goto st0
tr427:
//line parse.rl:330

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
        
	goto st96
tr431:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:64

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
        
	goto st96
tr434:
//line parse.rl:64

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
        
	goto st96
tr436:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:330

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
        
	goto st96
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
//line parse.go:7141
		switch ( s)[p] {
		case 42:
			goto tr202
		case 70:
			goto st97
		case 77:
			goto st99
		case 83:
			goto st101
		case 84:
			goto st104
		case 87:
			goto st107
		case 102:
			goto st97
		case 109:
			goto st99
		case 115:
			goto st101
		case 116:
			goto st104
		case 119:
			goto st107
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr203
		}
		goto st0
tr203:
//line parse.rl:38

            mark = p;
        
	goto st222
	st222:
		if p++; p == pe {
			goto _test_eof222
		}
	st_case_222:
//line parse.go:7181
		switch ( s)[p] {
		case 32:
			goto tr430
		case 44:
			goto tr431
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st222
			}
		case ( s)[p] >= 9:
			goto tr430
		}
		goto st0
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
		switch ( s)[p] {
		case 82:
			goto st98
		case 114:
			goto st98
		}
		goto st0
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
		switch ( s)[p] {
		case 73:
			goto tr210
		case 105:
			goto tr210
		}
		goto st0
tr210:
//line parse.rl:437
m=5
	goto st223
tr212:
//line parse.rl:437
m=1
	goto st223
tr215:
//line parse.rl:437
m=6
	goto st223
tr216:
//line parse.rl:437
m=0
	goto st223
tr219:
//line parse.rl:437
m=4
	goto st223
tr220:
//line parse.rl:437
m=2
	goto st223
tr222:
//line parse.rl:437
m=3
	goto st223
	st223:
		if p++; p == pe {
			goto _test_eof223
		}
	st_case_223:
//line parse.go:7254
		switch ( s)[p] {
		case 32:
			goto tr433
		case 44:
			goto tr434
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr433
		}
		goto st0
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
		switch ( s)[p] {
		case 79:
			goto st100
		case 111:
			goto st100
		}
		goto st0
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
		switch ( s)[p] {
		case 78:
			goto tr212
		case 110:
			goto tr212
		}
		goto st0
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
		switch ( s)[p] {
		case 65:
			goto st102
		case 85:
			goto st103
		case 97:
			goto st102
		case 117:
			goto st103
		}
		goto st0
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
		switch ( s)[p] {
		case 84:
			goto tr215
		case 116:
			goto tr215
		}
		goto st0
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
		switch ( s)[p] {
		case 78:
			goto tr216
		case 110:
			goto tr216
		}
		goto st0
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
		switch ( s)[p] {
		case 72:
			goto st105
		case 85:
			goto st106
		case 104:
			goto st105
		case 117:
			goto st106
		}
		goto st0
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
		switch ( s)[p] {
		case 85:
			goto tr219
		case 117:
			goto tr219
		}
		goto st0
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
		switch ( s)[p] {
		case 69:
			goto tr220
		case 101:
			goto tr220
		}
		goto st0
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
		switch ( s)[p] {
		case 69:
			goto st108
		case 101:
			goto st108
		}
		goto st0
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
		switch ( s)[p] {
		case 68:
			goto tr222
		case 100:
			goto tr222
		}
		goto st0
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr223
		}
		goto st0
tr223:
//line parse.rl:38

            mark = p;
        
	goto st224
	st224:
		if p++; p == pe {
			goto _test_eof224
		}
	st_case_224:
//line parse.go:7413
		switch ( s)[p] {
		case 32:
			goto tr435
		case 44:
			goto tr436
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st224
			}
		case ( s)[p] >= 9:
			goto tr435
		}
		goto st0
tr199:
//line parse.rl:307

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
        
	goto st110
tr225:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:50

            if m>12 || m<1{
                return nt, fmt.Errorf("invalid month %d", m)
            }
            nt.month |= 1<<(m-1) // we zero index months
        
	goto st110
tr231:
//line parse.rl:50

            if m>12 || m<1{
                return nt, fmt.Errorf("invalid month %d", m)
            }
            nt.month |= 1<<(m-1) // we zero index months
        
	goto st110
tr253:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:307

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
        
	goto st110
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
//line parse.go:7518
		switch ( s)[p] {
		case 42:
			goto tr188
		case 65:
			goto st112
		case 68:
			goto st116
		case 70:
			goto st118
		case 74:
			goto st120
		case 77:
			goto st123
		case 78:
			goto st125
		case 79:
			goto st127
		case 83:
			goto st129
		case 97:
			goto st112
		case 100:
			goto st116
		case 102:
			goto st118
		case 106:
			goto st120
		case 109:
			goto st123
		case 110:
			goto st125
		case 111:
			goto st127
		case 115:
			goto st129
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr189
		}
		goto st0
tr189:
//line parse.rl:38

            mark = p;
        
	goto st111
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
//line parse.go:7570
		switch ( s)[p] {
		case 32:
			goto tr224
		case 44:
			goto tr225
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st111
			}
		case ( s)[p] >= 9:
			goto tr224
		}
		goto st0
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
		switch ( s)[p] {
		case 80:
			goto st113
		case 85:
			goto st115
		case 112:
			goto st113
		case 117:
			goto st115
		}
		goto st0
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
		switch ( s)[p] {
		case 82:
			goto tr229
		case 114:
			goto tr229
		}
		goto st0
tr229:
//line parse.rl:438
m=4
	goto st114
tr232:
//line parse.rl:438
m=8
	goto st114
tr234:
//line parse.rl:438
m=12
	goto st114
tr236:
//line parse.rl:438
m=2
	goto st114
tr239:
//line parse.rl:438
m=1
	goto st114
tr240:
//line parse.rl:438
m=7
	goto st114
tr241:
//line parse.rl:438
m=6
	goto st114
tr243:
//line parse.rl:438
m=3
	goto st114
tr244:
//line parse.rl:438
m=5
	goto st114
tr246:
//line parse.rl:438
m=11
	goto st114
tr248:
//line parse.rl:438
m=10
	goto st114
tr250:
//line parse.rl:438
m=9
	goto st114
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
//line parse.go:7667
		switch ( s)[p] {
		case 32:
			goto tr230
		case 44:
			goto tr231
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr230
		}
		goto st0
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
		switch ( s)[p] {
		case 71:
			goto tr232
		case 103:
			goto tr232
		}
		goto st0
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
		switch ( s)[p] {
		case 69:
			goto st117
		case 101:
			goto st117
		}
		goto st0
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
		switch ( s)[p] {
		case 67:
			goto tr234
		case 99:
			goto tr234
		}
		goto st0
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
		switch ( s)[p] {
		case 69:
			goto st119
		case 101:
			goto st119
		}
		goto st0
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
		switch ( s)[p] {
		case 66:
			goto tr236
		case 98:
			goto tr236
		}
		goto st0
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
		switch ( s)[p] {
		case 65:
			goto st121
		case 85:
			goto st122
		case 97:
			goto st121
		case 117:
			goto st122
		}
		goto st0
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
		switch ( s)[p] {
		case 78:
			goto tr239
		case 110:
			goto tr239
		}
		goto st0
	st122:
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
		switch ( s)[p] {
		case 76:
			goto tr240
		case 78:
			goto tr241
		case 108:
			goto tr240
		case 110:
			goto tr241
		}
		goto st0
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
		switch ( s)[p] {
		case 65:
			goto st124
		case 97:
			goto st124
		}
		goto st0
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
		switch ( s)[p] {
		case 82:
			goto tr243
		case 89:
			goto tr244
		case 114:
			goto tr243
		case 121:
			goto tr244
		}
		goto st0
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
		switch ( s)[p] {
		case 79:
			goto st126
		case 111:
			goto st126
		}
		goto st0
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
		switch ( s)[p] {
		case 86:
			goto tr246
		case 118:
			goto tr246
		}
		goto st0
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
		switch ( s)[p] {
		case 67:
			goto st128
		case 99:
			goto st128
		}
		goto st0
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
		switch ( s)[p] {
		case 84:
			goto tr248
		case 116:
			goto tr248
		}
		goto st0
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
		switch ( s)[p] {
		case 69:
			goto st130
		case 101:
			goto st130
		}
		goto st0
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
		switch ( s)[p] {
		case 80:
			goto tr250
		case 112:
			goto tr250
		}
		goto st0
	st131:
		if p++; p == pe {
			goto _test_eof131
		}
	st_case_131:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr251
		}
		goto st0
tr251:
//line parse.rl:38

            mark = p;
        
	goto st132
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
//line parse.go:7902
		switch ( s)[p] {
		case 32:
			goto tr252
		case 44:
			goto tr253
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st132
			}
		case ( s)[p] >= 9:
			goto tr252
		}
		goto st0
tr185:
//line parse.rl:370

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
        
	goto st133
tr256:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:85

            if m>30 || m <0 {
                return nt,  fmt.Errorf("invalid day of month %d", m)
            }
            nt.dom |= 1<<(m-1)// again we zero index day
        
	goto st133
tr260:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:370

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
        
	goto st133
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
//line parse.go:8036
		if ( s)[p] == 42 {
			goto tr182
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr183
		}
		goto st0
tr183:
//line parse.rl:38

            mark = p;
        
	goto st134
	st134:
		if p++; p == pe {
			goto _test_eof134
		}
	st_case_134:
//line parse.go:8055
		switch ( s)[p] {
		case 32:
			goto tr255
		case 44:
			goto tr256
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st134
			}
		case ( s)[p] >= 9:
			goto tr255
		}
		goto st0
	st135:
		if p++; p == pe {
			goto _test_eof135
		}
	st_case_135:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr258
		}
		goto st0
tr258:
//line parse.rl:38

            mark = p;
        
	goto st136
	st136:
		if p++; p == pe {
			goto _test_eof136
		}
	st_case_136:
//line parse.go:8091
		switch ( s)[p] {
		case 32:
			goto tr259
		case 44:
			goto tr260
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st136
			}
		case ( s)[p] >= 9:
			goto tr259
		}
		goto st0
tr179:
//line parse.rl:273

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
        
	goto st137
tr263:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:78

            if m>24 || m <0 {
                return nt, fmt.Errorf("invalid hour of day %d", m)
            }
            nt.hour |= 1<<m
        
	goto st137
tr267:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:273

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
        
	goto st137
	st137:
		if p++; p == pe {
			goto _test_eof137
		}
	st_case_137:
//line parse.go:8209
		if ( s)[p] == 42 {
			goto tr176
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr177
		}
		goto st0
tr177:
//line parse.rl:38

            mark = p;
        
	goto st138
	st138:
		if p++; p == pe {
			goto _test_eof138
		}
	st_case_138:
//line parse.go:8228
		switch ( s)[p] {
		case 32:
			goto tr262
		case 44:
			goto tr263
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st138
			}
		case ( s)[p] >= 9:
			goto tr262
		}
		goto st0
	st139:
		if p++; p == pe {
			goto _test_eof139
		}
	st_case_139:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr265
		}
		goto st0
tr265:
//line parse.rl:38

            mark = p;
        
	goto st140
	st140:
		if p++; p == pe {
			goto _test_eof140
		}
	st_case_140:
//line parse.go:8264
		switch ( s)[p] {
		case 32:
			goto tr266
		case 44:
			goto tr267
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st140
			}
		case ( s)[p] >= 9:
			goto tr266
		}
		goto st0
tr173:
//line parse.rl:450
 start=0;end=59;m=1;d=1; 
//line parse.rl:186

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
        
	goto st141
tr270:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 start=m; end=0;d=0;
//line parse.rl:186

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
        
	goto st141
tr276:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 end=m; d=1;
//line parse.rl:186

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
        
	goto st141
tr281:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:456
d=m
//line parse.rl:186

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
        
	goto st141
	st141:
		if p++; p == pe {
			goto _test_eof141
		}
	st_case_141:
//line parse.go:8673
		if ( s)[p] == 42 {
			goto tr170
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr171
		}
		goto st0
tr171:
//line parse.rl:450
d=0;
//line parse.rl:38

            mark = p;
        
	goto st142
	st142:
		if p++; p == pe {
			goto _test_eof142
		}
	st_case_142:
//line parse.go:8694
		switch ( s)[p] {
		case 32:
			goto tr269
		case 44:
			goto tr270
		case 45:
			goto tr271
		case 47:
			goto tr272
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st142
			}
		case ( s)[p] >= 9:
			goto tr269
		}
		goto st0
tr271:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 start=m; end=0;d=0;
	goto st143
	st143:
		if p++; p == pe {
			goto _test_eof143
		}
	st_case_143:
//line parse.go:8731
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr274
		}
		goto st0
tr274:
//line parse.rl:38

            mark = p;
        
	goto st144
	st144:
		if p++; p == pe {
			goto _test_eof144
		}
	st_case_144:
//line parse.go:8747
		switch ( s)[p] {
		case 32:
			goto tr275
		case 44:
			goto tr276
		case 47:
			goto tr277
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st144
			}
		case ( s)[p] >= 9:
			goto tr275
		}
		goto st0
tr174:
//line parse.rl:450
 start=0;end=59;m=1;d=1; 
	goto st145
tr272:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 start=m; end=0;d=0;
	goto st145
tr277:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 end=m; d=1;
	goto st145
	st145:
		if p++; p == pe {
			goto _test_eof145
		}
	st_case_145:
//line parse.go:8798
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr279
		}
		goto st0
tr279:
//line parse.rl:38

            mark = p;
        
	goto st146
	st146:
		if p++; p == pe {
			goto _test_eof146
		}
	st_case_146:
//line parse.go:8814
		switch ( s)[p] {
		case 32:
			goto tr280
		case 44:
			goto tr281
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st146
			}
		case ( s)[p] >= 9:
			goto tr280
		}
		goto st0
tr168:
//line parse.rl:450
 start=0;end=59;m=1;d=1; 
	goto st147
tr290:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 start=m; end=0;d=0;
	goto st147
tr295:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 end=m; d=1;
	goto st147
	st147:
		if p++; p == pe {
			goto _test_eof147
		}
	st_case_147:
//line parse.go:8863
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr283
		}
		goto st0
tr283:
//line parse.rl:38

            mark = p;
        
	goto st148
	st148:
		if p++; p == pe {
			goto _test_eof148
		}
	st_case_148:
//line parse.go:8879
		switch ( s)[p] {
		case 32:
			goto tr284
		case 44:
			goto tr285
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st148
			}
		case ( s)[p] >= 9:
			goto tr284
		}
		goto st0
tr165:
//line parse.rl:450
d=0;
//line parse.rl:38

            mark = p;
        
	goto st149
	st149:
		if p++; p == pe {
			goto _test_eof149
		}
	st_case_149:
//line parse.go:8908
		switch ( s)[p] {
		case 32:
			goto tr287
		case 44:
			goto tr288
		case 45:
			goto tr289
		case 47:
			goto tr290
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st149
			}
		case ( s)[p] >= 9:
			goto tr287
		}
		goto st0
tr289:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 start=m; end=0;d=0;
	goto st150
	st150:
		if p++; p == pe {
			goto _test_eof150
		}
	st_case_150:
//line parse.go:8945
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr292
		}
		goto st0
tr292:
//line parse.rl:38

            mark = p;
        
	goto st151
	st151:
		if p++; p == pe {
			goto _test_eof151
		}
	st_case_151:
//line parse.go:8961
		switch ( s)[p] {
		case 32:
			goto tr293
		case 44:
			goto tr294
		case 47:
			goto tr295
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st151
			}
		case ( s)[p] >= 9:
			goto tr293
		}
		goto st0
tr300:
//line parse.rl:450
 start=0;end=59;m=1;d=1; 
//line parse.rl:186

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
        
	goto st152
tr401:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 start=m; end=0;d=0;
//line parse.rl:186

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
        
	goto st152
tr407:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 end=m; d=1;
//line parse.rl:186

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
        
	goto st152
tr398:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:456
d=m
//line parse.rl:186

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
        
	goto st152
	st152:
//line NONE:1
ts = 0

		if p++; p == pe {
			goto _test_eof152
		}
	st_case_152:
//line parse.go:9375
		if ( s)[p] == 42 {
			goto tr297
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr298
		}
		goto st0
tr297:
//line parse.rl:450
d=0;
	goto st153
	st153:
		if p++; p == pe {
			goto _test_eof153
		}
	st_case_153:
//line parse.go:9392
		switch ( s)[p] {
		case 32:
			goto tr299
		case 44:
			goto tr300
		case 47:
			goto tr301
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr299
		}
		goto st0
tr299:
//line parse.rl:450
 start=0;end=59;m=1;d=1; 
//line parse.rl:186

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
        
	goto st154
tr400:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 start=m; end=0;d=0;
//line parse.rl:186

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
        
	goto st154
tr406:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 end=m; d=1;
//line parse.rl:186

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
        
	goto st154
tr397:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:456
d=m
//line parse.rl:186

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
        
	goto st154
	st154:
		if p++; p == pe {
			goto _test_eof154
		}
	st_case_154:
//line parse.go:9798
		switch ( s)[p] {
		case 32:
			goto st154
		case 42:
			goto tr303
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr304
			}
		case ( s)[p] >= 9:
			goto st154
		}
		goto st0
tr303:
//line parse.rl:446
m=1
	goto st155
	st155:
		if p++; p == pe {
			goto _test_eof155
		}
	st_case_155:
//line parse.go:9823
		switch ( s)[p] {
		case 32:
			goto tr305
		case 44:
			goto tr306
		case 47:
			goto st204
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr305
		}
		goto st0
tr305:
//line parse.rl:273

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
        
	goto st156
tr389:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:78

            if m>24 || m <0 {
                return nt, fmt.Errorf("invalid hour of day %d", m)
            }
            nt.hour |= 1<<m
        
	goto st156
tr393:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:273

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
        
	goto st156
	st156:
		if p++; p == pe {
			goto _test_eof156
		}
	st_case_156:
//line parse.go:9938
		switch ( s)[p] {
		case 32:
			goto st156
		case 42:
			goto tr309
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr310
			}
		case ( s)[p] >= 9:
			goto st156
		}
		goto st0
tr309:
//line parse.rl:446
m=1
	goto st157
	st157:
		if p++; p == pe {
			goto _test_eof157
		}
	st_case_157:
//line parse.go:9963
		switch ( s)[p] {
		case 32:
			goto tr311
		case 44:
			goto tr312
		case 47:
			goto st200
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr311
		}
		goto st0
tr311:
//line parse.rl:370

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
        
	goto st158
tr382:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:85

            if m>30 || m <0 {
                return nt,  fmt.Errorf("invalid day of month %d", m)
            }
            nt.dom |= 1<<(m-1)// again we zero index day
        
	goto st158
tr386:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:370

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
        
	goto st158
	st158:
		if p++; p == pe {
			goto _test_eof158
		}
	st_case_158:
//line parse.go:10094
		switch ( s)[p] {
		case 32:
			goto st158
		case 42:
			goto tr315
		case 65:
			goto st177
		case 68:
			goto st181
		case 70:
			goto st183
		case 74:
			goto st185
		case 77:
			goto st188
		case 78:
			goto st190
		case 79:
			goto st192
		case 83:
			goto st194
		case 97:
			goto st177
		case 100:
			goto st181
		case 102:
			goto st183
		case 106:
			goto st185
		case 109:
			goto st188
		case 110:
			goto st190
		case 111:
			goto st192
		case 115:
			goto st194
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr316
			}
		case ( s)[p] >= 9:
			goto st158
		}
		goto st0
tr315:
//line parse.rl:446
m=1
	goto st159
	st159:
		if p++; p == pe {
			goto _test_eof159
		}
	st_case_159:
//line parse.go:10151
		switch ( s)[p] {
		case 32:
			goto tr325
		case 44:
			goto tr326
		case 47:
			goto st196
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr325
		}
		goto st0
tr325:
//line parse.rl:307

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
        
	goto st160
tr351:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:50

            if m>12 || m<1{
                return nt, fmt.Errorf("invalid month %d", m)
            }
            nt.month |= 1<<(m-1) // we zero index months
        
	goto st160
tr357:
//line parse.rl:50

            if m>12 || m<1{
                return nt, fmt.Errorf("invalid month %d", m)
            }
            nt.month |= 1<<(m-1) // we zero index months
        
	goto st160
tr379:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:307

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
        
	goto st160
	st160:
		if p++; p == pe {
			goto _test_eof160
		}
	st_case_160:
//line parse.go:10253
		switch ( s)[p] {
		case 32:
			goto st160
		case 42:
			goto tr329
		case 70:
			goto st162
		case 77:
			goto st164
		case 83:
			goto st166
		case 84:
			goto st169
		case 87:
			goto st172
		case 102:
			goto st162
		case 109:
			goto st164
		case 115:
			goto st166
		case 116:
			goto st169
		case 119:
			goto st172
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto tr330
			}
		case ( s)[p] >= 9:
			goto st160
		}
		goto st0
tr329:
//line parse.rl:446
m=1
	goto st225
	st225:
		if p++; p == pe {
			goto _test_eof225
		}
	st_case_225:
//line parse.go:10298
		switch ( s)[p] {
		case 32:
			goto tr438
		case 44:
			goto tr439
		case 47:
			goto st174
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr438
		}
		goto st0
tr438:
//line parse.rl:330

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
        
	goto st226
tr442:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:64

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
        
	goto st226
tr445:
//line parse.rl:64

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
        
	goto st226
tr447:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:330

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
        
	goto st226
	st226:
		if p++; p == pe {
			goto _test_eof226
		}
	st_case_226:
//line parse.go:10414
		if ( s)[p] == 32 {
			goto st226
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto st226
		}
		goto st0
tr439:
//line parse.rl:330

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
        
	goto st161
tr443:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:64

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
        
	goto st161
tr446:
//line parse.rl:64

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
        
	goto st161
tr448:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:330

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
        
	goto st161
	st161:
		if p++; p == pe {
			goto _test_eof161
		}
	st_case_161:
//line parse.go:10525
		switch ( s)[p] {
		case 42:
			goto tr329
		case 70:
			goto st162
		case 77:
			goto st164
		case 83:
			goto st166
		case 84:
			goto st169
		case 87:
			goto st172
		case 102:
			goto st162
		case 109:
			goto st164
		case 115:
			goto st166
		case 116:
			goto st169
		case 119:
			goto st172
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr330
		}
		goto st0
tr330:
//line parse.rl:38

            mark = p;
        
	goto st227
	st227:
		if p++; p == pe {
			goto _test_eof227
		}
	st_case_227:
//line parse.go:10565
		switch ( s)[p] {
		case 32:
			goto tr442
		case 44:
			goto tr443
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st227
			}
		case ( s)[p] >= 9:
			goto tr442
		}
		goto st0
	st162:
		if p++; p == pe {
			goto _test_eof162
		}
	st_case_162:
		switch ( s)[p] {
		case 82:
			goto st163
		case 114:
			goto st163
		}
		goto st0
	st163:
		if p++; p == pe {
			goto _test_eof163
		}
	st_case_163:
		switch ( s)[p] {
		case 73:
			goto tr337
		case 105:
			goto tr337
		}
		goto st0
tr337:
//line parse.rl:437
m=5
	goto st228
tr339:
//line parse.rl:437
m=1
	goto st228
tr342:
//line parse.rl:437
m=6
	goto st228
tr343:
//line parse.rl:437
m=0
	goto st228
tr346:
//line parse.rl:437
m=4
	goto st228
tr347:
//line parse.rl:437
m=2
	goto st228
tr349:
//line parse.rl:437
m=3
	goto st228
	st228:
		if p++; p == pe {
			goto _test_eof228
		}
	st_case_228:
//line parse.go:10638
		switch ( s)[p] {
		case 32:
			goto tr445
		case 44:
			goto tr446
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr445
		}
		goto st0
	st164:
		if p++; p == pe {
			goto _test_eof164
		}
	st_case_164:
		switch ( s)[p] {
		case 79:
			goto st165
		case 111:
			goto st165
		}
		goto st0
	st165:
		if p++; p == pe {
			goto _test_eof165
		}
	st_case_165:
		switch ( s)[p] {
		case 78:
			goto tr339
		case 110:
			goto tr339
		}
		goto st0
	st166:
		if p++; p == pe {
			goto _test_eof166
		}
	st_case_166:
		switch ( s)[p] {
		case 65:
			goto st167
		case 85:
			goto st168
		case 97:
			goto st167
		case 117:
			goto st168
		}
		goto st0
	st167:
		if p++; p == pe {
			goto _test_eof167
		}
	st_case_167:
		switch ( s)[p] {
		case 84:
			goto tr342
		case 116:
			goto tr342
		}
		goto st0
	st168:
		if p++; p == pe {
			goto _test_eof168
		}
	st_case_168:
		switch ( s)[p] {
		case 78:
			goto tr343
		case 110:
			goto tr343
		}
		goto st0
	st169:
		if p++; p == pe {
			goto _test_eof169
		}
	st_case_169:
		switch ( s)[p] {
		case 72:
			goto st170
		case 85:
			goto st171
		case 104:
			goto st170
		case 117:
			goto st171
		}
		goto st0
	st170:
		if p++; p == pe {
			goto _test_eof170
		}
	st_case_170:
		switch ( s)[p] {
		case 85:
			goto tr346
		case 117:
			goto tr346
		}
		goto st0
	st171:
		if p++; p == pe {
			goto _test_eof171
		}
	st_case_171:
		switch ( s)[p] {
		case 69:
			goto tr347
		case 101:
			goto tr347
		}
		goto st0
	st172:
		if p++; p == pe {
			goto _test_eof172
		}
	st_case_172:
		switch ( s)[p] {
		case 69:
			goto st173
		case 101:
			goto st173
		}
		goto st0
	st173:
		if p++; p == pe {
			goto _test_eof173
		}
	st_case_173:
		switch ( s)[p] {
		case 68:
			goto tr349
		case 100:
			goto tr349
		}
		goto st0
	st174:
		if p++; p == pe {
			goto _test_eof174
		}
	st_case_174:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr350
		}
		goto st0
tr350:
//line parse.rl:38

            mark = p;
        
	goto st229
	st229:
		if p++; p == pe {
			goto _test_eof229
		}
	st_case_229:
//line parse.go:10797
		switch ( s)[p] {
		case 32:
			goto tr447
		case 44:
			goto tr448
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st229
			}
		case ( s)[p] >= 9:
			goto tr447
		}
		goto st0
tr326:
//line parse.rl:307

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
        
	goto st175
tr352:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:50

            if m>12 || m<1{
                return nt, fmt.Errorf("invalid month %d", m)
            }
            nt.month |= 1<<(m-1) // we zero index months
        
	goto st175
tr358:
//line parse.rl:50

            if m>12 || m<1{
                return nt, fmt.Errorf("invalid month %d", m)
            }
            nt.month |= 1<<(m-1) // we zero index months
        
	goto st175
tr380:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:307

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
        
	goto st175
	st175:
		if p++; p == pe {
			goto _test_eof175
		}
	st_case_175:
//line parse.go:10902
		switch ( s)[p] {
		case 42:
			goto tr315
		case 65:
			goto st177
		case 68:
			goto st181
		case 70:
			goto st183
		case 74:
			goto st185
		case 77:
			goto st188
		case 78:
			goto st190
		case 79:
			goto st192
		case 83:
			goto st194
		case 97:
			goto st177
		case 100:
			goto st181
		case 102:
			goto st183
		case 106:
			goto st185
		case 109:
			goto st188
		case 110:
			goto st190
		case 111:
			goto st192
		case 115:
			goto st194
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr316
		}
		goto st0
tr316:
//line parse.rl:38

            mark = p;
        
	goto st176
	st176:
		if p++; p == pe {
			goto _test_eof176
		}
	st_case_176:
//line parse.go:10954
		switch ( s)[p] {
		case 32:
			goto tr351
		case 44:
			goto tr352
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st176
			}
		case ( s)[p] >= 9:
			goto tr351
		}
		goto st0
	st177:
		if p++; p == pe {
			goto _test_eof177
		}
	st_case_177:
		switch ( s)[p] {
		case 80:
			goto st178
		case 85:
			goto st180
		case 112:
			goto st178
		case 117:
			goto st180
		}
		goto st0
	st178:
		if p++; p == pe {
			goto _test_eof178
		}
	st_case_178:
		switch ( s)[p] {
		case 82:
			goto tr356
		case 114:
			goto tr356
		}
		goto st0
tr356:
//line parse.rl:438
m=4
	goto st179
tr359:
//line parse.rl:438
m=8
	goto st179
tr361:
//line parse.rl:438
m=12
	goto st179
tr363:
//line parse.rl:438
m=2
	goto st179
tr366:
//line parse.rl:438
m=1
	goto st179
tr367:
//line parse.rl:438
m=7
	goto st179
tr368:
//line parse.rl:438
m=6
	goto st179
tr370:
//line parse.rl:438
m=3
	goto st179
tr371:
//line parse.rl:438
m=5
	goto st179
tr373:
//line parse.rl:438
m=11
	goto st179
tr375:
//line parse.rl:438
m=10
	goto st179
tr377:
//line parse.rl:438
m=9
	goto st179
	st179:
		if p++; p == pe {
			goto _test_eof179
		}
	st_case_179:
//line parse.go:11051
		switch ( s)[p] {
		case 32:
			goto tr357
		case 44:
			goto tr358
		}
		if 9 <= ( s)[p] && ( s)[p] <= 13 {
			goto tr357
		}
		goto st0
	st180:
		if p++; p == pe {
			goto _test_eof180
		}
	st_case_180:
		switch ( s)[p] {
		case 71:
			goto tr359
		case 103:
			goto tr359
		}
		goto st0
	st181:
		if p++; p == pe {
			goto _test_eof181
		}
	st_case_181:
		switch ( s)[p] {
		case 69:
			goto st182
		case 101:
			goto st182
		}
		goto st0
	st182:
		if p++; p == pe {
			goto _test_eof182
		}
	st_case_182:
		switch ( s)[p] {
		case 67:
			goto tr361
		case 99:
			goto tr361
		}
		goto st0
	st183:
		if p++; p == pe {
			goto _test_eof183
		}
	st_case_183:
		switch ( s)[p] {
		case 69:
			goto st184
		case 101:
			goto st184
		}
		goto st0
	st184:
		if p++; p == pe {
			goto _test_eof184
		}
	st_case_184:
		switch ( s)[p] {
		case 66:
			goto tr363
		case 98:
			goto tr363
		}
		goto st0
	st185:
		if p++; p == pe {
			goto _test_eof185
		}
	st_case_185:
		switch ( s)[p] {
		case 65:
			goto st186
		case 85:
			goto st187
		case 97:
			goto st186
		case 117:
			goto st187
		}
		goto st0
	st186:
		if p++; p == pe {
			goto _test_eof186
		}
	st_case_186:
		switch ( s)[p] {
		case 78:
			goto tr366
		case 110:
			goto tr366
		}
		goto st0
	st187:
		if p++; p == pe {
			goto _test_eof187
		}
	st_case_187:
		switch ( s)[p] {
		case 76:
			goto tr367
		case 78:
			goto tr368
		case 108:
			goto tr367
		case 110:
			goto tr368
		}
		goto st0
	st188:
		if p++; p == pe {
			goto _test_eof188
		}
	st_case_188:
		switch ( s)[p] {
		case 65:
			goto st189
		case 97:
			goto st189
		}
		goto st0
	st189:
		if p++; p == pe {
			goto _test_eof189
		}
	st_case_189:
		switch ( s)[p] {
		case 82:
			goto tr370
		case 89:
			goto tr371
		case 114:
			goto tr370
		case 121:
			goto tr371
		}
		goto st0
	st190:
		if p++; p == pe {
			goto _test_eof190
		}
	st_case_190:
		switch ( s)[p] {
		case 79:
			goto st191
		case 111:
			goto st191
		}
		goto st0
	st191:
		if p++; p == pe {
			goto _test_eof191
		}
	st_case_191:
		switch ( s)[p] {
		case 86:
			goto tr373
		case 118:
			goto tr373
		}
		goto st0
	st192:
		if p++; p == pe {
			goto _test_eof192
		}
	st_case_192:
		switch ( s)[p] {
		case 67:
			goto st193
		case 99:
			goto st193
		}
		goto st0
	st193:
		if p++; p == pe {
			goto _test_eof193
		}
	st_case_193:
		switch ( s)[p] {
		case 84:
			goto tr375
		case 116:
			goto tr375
		}
		goto st0
	st194:
		if p++; p == pe {
			goto _test_eof194
		}
	st_case_194:
		switch ( s)[p] {
		case 69:
			goto st195
		case 101:
			goto st195
		}
		goto st0
	st195:
		if p++; p == pe {
			goto _test_eof195
		}
	st_case_195:
		switch ( s)[p] {
		case 80:
			goto tr377
		case 112:
			goto tr377
		}
		goto st0
	st196:
		if p++; p == pe {
			goto _test_eof196
		}
	st_case_196:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr378
		}
		goto st0
tr378:
//line parse.rl:38

            mark = p;
        
	goto st197
	st197:
		if p++; p == pe {
			goto _test_eof197
		}
	st_case_197:
//line parse.go:11286
		switch ( s)[p] {
		case 32:
			goto tr379
		case 44:
			goto tr380
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st197
			}
		case ( s)[p] >= 9:
			goto tr379
		}
		goto st0
tr312:
//line parse.rl:370

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
        
	goto st198
tr383:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:85

            if m>30 || m <0 {
                return nt,  fmt.Errorf("invalid day of month %d", m)
            }
            nt.dom |= 1<<(m-1)// again we zero index day
        
	goto st198
tr387:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:370

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
        
	goto st198
	st198:
		if p++; p == pe {
			goto _test_eof198
		}
	st_case_198:
//line parse.go:11420
		if ( s)[p] == 42 {
			goto tr309
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr310
		}
		goto st0
tr310:
//line parse.rl:38

            mark = p;
        
	goto st199
	st199:
		if p++; p == pe {
			goto _test_eof199
		}
	st_case_199:
//line parse.go:11439
		switch ( s)[p] {
		case 32:
			goto tr382
		case 44:
			goto tr383
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st199
			}
		case ( s)[p] >= 9:
			goto tr382
		}
		goto st0
	st200:
		if p++; p == pe {
			goto _test_eof200
		}
	st_case_200:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr385
		}
		goto st0
tr385:
//line parse.rl:38

            mark = p;
        
	goto st201
	st201:
		if p++; p == pe {
			goto _test_eof201
		}
	st_case_201:
//line parse.go:11475
		switch ( s)[p] {
		case 32:
			goto tr386
		case 44:
			goto tr387
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st201
			}
		case ( s)[p] >= 9:
			goto tr386
		}
		goto st0
tr306:
//line parse.rl:273

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
        
	goto st202
tr390:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:78

            if m>24 || m <0 {
                return nt, fmt.Errorf("invalid hour of day %d", m)
            }
            nt.hour |= 1<<m
        
	goto st202
tr394:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:273

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
        
	goto st202
	st202:
		if p++; p == pe {
			goto _test_eof202
		}
	st_case_202:
//line parse.go:11593
		if ( s)[p] == 42 {
			goto tr303
		}
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr304
		}
		goto st0
tr304:
//line parse.rl:38

            mark = p;
        
	goto st203
	st203:
		if p++; p == pe {
			goto _test_eof203
		}
	st_case_203:
//line parse.go:11612
		switch ( s)[p] {
		case 32:
			goto tr389
		case 44:
			goto tr390
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st203
			}
		case ( s)[p] >= 9:
			goto tr389
		}
		goto st0
	st204:
		if p++; p == pe {
			goto _test_eof204
		}
	st_case_204:
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr392
		}
		goto st0
tr392:
//line parse.rl:38

            mark = p;
        
	goto st205
	st205:
		if p++; p == pe {
			goto _test_eof205
		}
	st_case_205:
//line parse.go:11648
		switch ( s)[p] {
		case 32:
			goto tr393
		case 44:
			goto tr394
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st205
			}
		case ( s)[p] >= 9:
			goto tr393
		}
		goto st0
tr301:
//line parse.rl:450
 start=0;end=59;m=1;d=1; 
	goto st206
tr403:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 start=m; end=0;d=0;
	goto st206
tr408:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 end=m; d=1;
	goto st206
	st206:
		if p++; p == pe {
			goto _test_eof206
		}
	st_case_206:
//line parse.go:11697
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr396
		}
		goto st0
tr396:
//line parse.rl:38

            mark = p;
        
	goto st207
	st207:
		if p++; p == pe {
			goto _test_eof207
		}
	st_case_207:
//line parse.go:11713
		switch ( s)[p] {
		case 32:
			goto tr397
		case 44:
			goto tr398
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st207
			}
		case ( s)[p] >= 9:
			goto tr397
		}
		goto st0
tr298:
//line parse.rl:450
d=0;
//line parse.rl:38

            mark = p;
        
	goto st208
	st208:
		if p++; p == pe {
			goto _test_eof208
		}
	st_case_208:
//line parse.go:11742
		switch ( s)[p] {
		case 32:
			goto tr400
		case 44:
			goto tr401
		case 45:
			goto tr402
		case 47:
			goto tr403
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st208
			}
		case ( s)[p] >= 9:
			goto tr400
		}
		goto st0
tr402:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:450
 start=m; end=0;d=0;
	goto st209
	st209:
		if p++; p == pe {
			goto _test_eof209
		}
	st_case_209:
//line parse.go:11779
		if 48 <= ( s)[p] && ( s)[p] <= 57 {
			goto tr405
		}
		goto st0
tr405:
//line parse.rl:38

            mark = p;
        
	goto st210
	st210:
		if p++; p == pe {
			goto _test_eof210
		}
	st_case_210:
//line parse.go:11795
		switch ( s)[p] {
		case 32:
			goto tr406
		case 44:
			goto tr407
		case 47:
			goto tr408
		}
		switch {
		case ( s)[p] > 13:
			if 48 <= ( s)[p] && ( s)[p] <= 57 {
				goto st210
			}
		case ( s)[p] >= 9:
			goto tr406
		}
		goto st0
	st_out:
	_test_eof211: cs = 211; goto _test_eof
	_test_eof212: cs = 212; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof213: cs = 213; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof214: cs = 214; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof215: cs = 215; goto _test_eof
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
	_test_eof216: cs = 216; goto _test_eof
	_test_eof217: cs = 217; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof218: cs = 218; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof219: cs = 219; goto _test_eof
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
	_test_eof220: cs = 220; goto _test_eof
	_test_eof221: cs = 221; goto _test_eof
	_test_eof96: cs = 96; goto _test_eof
	_test_eof222: cs = 222; goto _test_eof
	_test_eof97: cs = 97; goto _test_eof
	_test_eof98: cs = 98; goto _test_eof
	_test_eof223: cs = 223; goto _test_eof
	_test_eof99: cs = 99; goto _test_eof
	_test_eof100: cs = 100; goto _test_eof
	_test_eof101: cs = 101; goto _test_eof
	_test_eof102: cs = 102; goto _test_eof
	_test_eof103: cs = 103; goto _test_eof
	_test_eof104: cs = 104; goto _test_eof
	_test_eof105: cs = 105; goto _test_eof
	_test_eof106: cs = 106; goto _test_eof
	_test_eof107: cs = 107; goto _test_eof
	_test_eof108: cs = 108; goto _test_eof
	_test_eof109: cs = 109; goto _test_eof
	_test_eof224: cs = 224; goto _test_eof
	_test_eof110: cs = 110; goto _test_eof
	_test_eof111: cs = 111; goto _test_eof
	_test_eof112: cs = 112; goto _test_eof
	_test_eof113: cs = 113; goto _test_eof
	_test_eof114: cs = 114; goto _test_eof
	_test_eof115: cs = 115; goto _test_eof
	_test_eof116: cs = 116; goto _test_eof
	_test_eof117: cs = 117; goto _test_eof
	_test_eof118: cs = 118; goto _test_eof
	_test_eof119: cs = 119; goto _test_eof
	_test_eof120: cs = 120; goto _test_eof
	_test_eof121: cs = 121; goto _test_eof
	_test_eof122: cs = 122; goto _test_eof
	_test_eof123: cs = 123; goto _test_eof
	_test_eof124: cs = 124; goto _test_eof
	_test_eof125: cs = 125; goto _test_eof
	_test_eof126: cs = 126; goto _test_eof
	_test_eof127: cs = 127; goto _test_eof
	_test_eof128: cs = 128; goto _test_eof
	_test_eof129: cs = 129; goto _test_eof
	_test_eof130: cs = 130; goto _test_eof
	_test_eof131: cs = 131; goto _test_eof
	_test_eof132: cs = 132; goto _test_eof
	_test_eof133: cs = 133; goto _test_eof
	_test_eof134: cs = 134; goto _test_eof
	_test_eof135: cs = 135; goto _test_eof
	_test_eof136: cs = 136; goto _test_eof
	_test_eof137: cs = 137; goto _test_eof
	_test_eof138: cs = 138; goto _test_eof
	_test_eof139: cs = 139; goto _test_eof
	_test_eof140: cs = 140; goto _test_eof
	_test_eof141: cs = 141; goto _test_eof
	_test_eof142: cs = 142; goto _test_eof
	_test_eof143: cs = 143; goto _test_eof
	_test_eof144: cs = 144; goto _test_eof
	_test_eof145: cs = 145; goto _test_eof
	_test_eof146: cs = 146; goto _test_eof
	_test_eof147: cs = 147; goto _test_eof
	_test_eof148: cs = 148; goto _test_eof
	_test_eof149: cs = 149; goto _test_eof
	_test_eof150: cs = 150; goto _test_eof
	_test_eof151: cs = 151; goto _test_eof
	_test_eof152: cs = 152; goto _test_eof
	_test_eof153: cs = 153; goto _test_eof
	_test_eof154: cs = 154; goto _test_eof
	_test_eof155: cs = 155; goto _test_eof
	_test_eof156: cs = 156; goto _test_eof
	_test_eof157: cs = 157; goto _test_eof
	_test_eof158: cs = 158; goto _test_eof
	_test_eof159: cs = 159; goto _test_eof
	_test_eof160: cs = 160; goto _test_eof
	_test_eof225: cs = 225; goto _test_eof
	_test_eof226: cs = 226; goto _test_eof
	_test_eof161: cs = 161; goto _test_eof
	_test_eof227: cs = 227; goto _test_eof
	_test_eof162: cs = 162; goto _test_eof
	_test_eof163: cs = 163; goto _test_eof
	_test_eof228: cs = 228; goto _test_eof
	_test_eof164: cs = 164; goto _test_eof
	_test_eof165: cs = 165; goto _test_eof
	_test_eof166: cs = 166; goto _test_eof
	_test_eof167: cs = 167; goto _test_eof
	_test_eof168: cs = 168; goto _test_eof
	_test_eof169: cs = 169; goto _test_eof
	_test_eof170: cs = 170; goto _test_eof
	_test_eof171: cs = 171; goto _test_eof
	_test_eof172: cs = 172; goto _test_eof
	_test_eof173: cs = 173; goto _test_eof
	_test_eof174: cs = 174; goto _test_eof
	_test_eof229: cs = 229; goto _test_eof
	_test_eof175: cs = 175; goto _test_eof
	_test_eof176: cs = 176; goto _test_eof
	_test_eof177: cs = 177; goto _test_eof
	_test_eof178: cs = 178; goto _test_eof
	_test_eof179: cs = 179; goto _test_eof
	_test_eof180: cs = 180; goto _test_eof
	_test_eof181: cs = 181; goto _test_eof
	_test_eof182: cs = 182; goto _test_eof
	_test_eof183: cs = 183; goto _test_eof
	_test_eof184: cs = 184; goto _test_eof
	_test_eof185: cs = 185; goto _test_eof
	_test_eof186: cs = 186; goto _test_eof
	_test_eof187: cs = 187; goto _test_eof
	_test_eof188: cs = 188; goto _test_eof
	_test_eof189: cs = 189; goto _test_eof
	_test_eof190: cs = 190; goto _test_eof
	_test_eof191: cs = 191; goto _test_eof
	_test_eof192: cs = 192; goto _test_eof
	_test_eof193: cs = 193; goto _test_eof
	_test_eof194: cs = 194; goto _test_eof
	_test_eof195: cs = 195; goto _test_eof
	_test_eof196: cs = 196; goto _test_eof
	_test_eof197: cs = 197; goto _test_eof
	_test_eof198: cs = 198; goto _test_eof
	_test_eof199: cs = 199; goto _test_eof
	_test_eof200: cs = 200; goto _test_eof
	_test_eof201: cs = 201; goto _test_eof
	_test_eof202: cs = 202; goto _test_eof
	_test_eof203: cs = 203; goto _test_eof
	_test_eof204: cs = 204; goto _test_eof
	_test_eof205: cs = 205; goto _test_eof
	_test_eof206: cs = 206; goto _test_eof
	_test_eof207: cs = 207; goto _test_eof
	_test_eof208: cs = 208; goto _test_eof
	_test_eof209: cs = 209; goto _test_eof
	_test_eof210: cs = 210; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 212:
			goto tr412
		case 213:
			goto tr413
		case 9:
			goto tr10
		case 214:
			goto tr414
		case 10:
			goto tr13
		case 215:
			goto tr415
		case 223, 228:
//line parse.rl:64

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
        
		case 220, 225:
//line parse.rl:330

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
        
		case 216:
//line parse.rl:353

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
        
		case 222, 227:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:64

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
        
		case 218:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:92

            if m<1970 || m>2099{
                return nt, errors.New("only years on interval [1970, 2099] are supported")
            }
            nt.year.set(int(m))
        
		case 224, 229:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:330

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
        
		case 219:
//line parse.rl:425

            fmt.Println("print digit", s[mark:p])
            m, err = strconv.ParseUint(s[mark:p], 10, 64)
            if err!=nil {
                return nt,  fmt.Errorf("unable to parse %s into a number", s[mark:p])
            }
        
//line parse.rl:353

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
        
//line parse.go:12207
		}
	}

	_out: {}
	}

//line parse.rl:499
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

