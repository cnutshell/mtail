// Code generated by goyacc -v y.output -o parser.go -p mtail parser.y. DO NOT EDIT.

//line parser.y:5
package vm

import __yyfmt__ "fmt"

//line parser.y:5

import (
	"time"

	"github.com/golang/glog"
	"github.com/google/mtail/metrics"
)

//line parser.y:16
type mtailSymType struct {
	yys      int
	intVal   int64
	floatVal float64
	op       int
	text     string
	texts    []string
	flag     bool
	n        astNode
	kind     metrics.Kind
	duration time.Duration
}

const INVALID = 57346
const COUNTER = 57347
const GAUGE = 57348
const TIMER = 57349
const TEXT = 57350
const AFTER = 57351
const AS = 57352
const BY = 57353
const CONST = 57354
const HIDDEN = 57355
const DEF = 57356
const DEL = 57357
const NEXT = 57358
const OTHERWISE = 57359
const ELSE = 57360
const BUILTIN = 57361
const REGEX = 57362
const STRING = 57363
const CAPREF = 57364
const CAPREF_NAMED = 57365
const ID = 57366
const DECO = 57367
const INTLITERAL = 57368
const FLOATLITERAL = 57369
const DURATIONLITERAL = 57370
const INC = 57371
const DEC = 57372
const DIV = 57373
const MOD = 57374
const MUL = 57375
const MINUS = 57376
const PLUS = 57377
const POW = 57378
const SHL = 57379
const SHR = 57380
const LT = 57381
const GT = 57382
const LE = 57383
const GE = 57384
const EQ = 57385
const NE = 57386
const BITAND = 57387
const XOR = 57388
const BITOR = 57389
const NOT = 57390
const AND = 57391
const OR = 57392
const ADD_ASSIGN = 57393
const ASSIGN = 57394
const CONCAT = 57395
const MATCH = 57396
const NOT_MATCH = 57397
const LCURLY = 57398
const RCURLY = 57399
const LPAREN = 57400
const RPAREN = 57401
const LSQUARE = 57402
const RSQUARE = 57403
const COMMA = 57404
const NL = 57405

var mtailToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"INVALID",
	"COUNTER",
	"GAUGE",
	"TIMER",
	"TEXT",
	"AFTER",
	"AS",
	"BY",
	"CONST",
	"HIDDEN",
	"DEF",
	"DEL",
	"NEXT",
	"OTHERWISE",
	"ELSE",
	"BUILTIN",
	"REGEX",
	"STRING",
	"CAPREF",
	"CAPREF_NAMED",
	"ID",
	"DECO",
	"INTLITERAL",
	"FLOATLITERAL",
	"DURATIONLITERAL",
	"INC",
	"DEC",
	"DIV",
	"MOD",
	"MUL",
	"MINUS",
	"PLUS",
	"POW",
	"SHL",
	"SHR",
	"LT",
	"GT",
	"LE",
	"GE",
	"EQ",
	"NE",
	"BITAND",
	"XOR",
	"BITOR",
	"NOT",
	"AND",
	"OR",
	"ADD_ASSIGN",
	"ASSIGN",
	"CONCAT",
	"MATCH",
	"NOT_MATCH",
	"LCURLY",
	"RCURLY",
	"LPAREN",
	"RPAREN",
	"LSQUARE",
	"RSQUARE",
	"COMMA",
	"NL",
}
var mtailStatenames = [...]string{}

const mtailEofCode = 1
const mtailErrCode = 2
const mtailInitialStackSize = 16

//line parser.y:595

//  tokenpos returns the position of the current token.
func tokenpos(mtaillex mtailLexer) position {
	return mtaillex.(*parser).t.pos
}

// markedpos returns the position recorded from the most recent mark_pos
// production.
func markedpos(mtaillex mtailLexer) position {
	return mtaillex.(*parser).pos
}

//line yacctab:1
var mtailExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	1, 1,
	14, 107,
	25, 107,
	31, 107,
	-2, 87,
	-1, 104,
	14, 107,
	25, 107,
	31, 107,
	-2, 87,
}

const mtailPrivate = 57344

const mtailLast = 219

var mtailAct = [...]int{

	20, 121, 62, 43, 27, 26, 42, 41, 25, 40,
	28, 46, 21, 103, 13, 102, 119, 87, 45, 24,
	18, 149, 147, 148, 148, 158, 51, 52, 31, 83,
	34, 32, 33, 44, 124, 36, 37, 27, 26, 84,
	48, 91, 2, 12, 75, 76, 49, 50, 49, 50,
	86, 11, 23, 48, 19, 10, 14, 39, 31, 82,
	34, 32, 33, 44, 29, 36, 37, 35, 122, 78,
	77, 156, 110, 31, 58, 34, 32, 33, 44, 112,
	36, 37, 155, 113, 120, 120, 137, 39, 80, 81,
	114, 104, 44, 115, 116, 117, 130, 35, 118, 94,
	93, 123, 15, 128, 109, 26, 27, 26, 125, 101,
	100, 126, 35, 127, 135, 129, 108, 141, 26, 26,
	1, 136, 18, 140, 139, 146, 145, 144, 151, 150,
	142, 143, 138, 12, 64, 66, 65, 16, 97, 98,
	96, 11, 23, 99, 19, 10, 14, 88, 31, 157,
	34, 32, 33, 44, 74, 36, 37, 31, 95, 34,
	32, 33, 44, 59, 36, 37, 68, 69, 70, 71,
	72, 73, 92, 85, 60, 89, 90, 39, 111, 160,
	58, 154, 159, 38, 153, 107, 39, 35, 106, 134,
	133, 47, 15, 63, 79, 67, 35, 17, 89, 90,
	152, 131, 132, 61, 54, 55, 56, 57, 53, 9,
	8, 7, 105, 6, 30, 22, 5, 4, 3,
}
var mtailPact = [...]int{

	-1000, -1000, 129, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 68, -1000, -3, -16, -1000, -36, 199, 149, 54,
	89, -1000, -1000, -1000, 127, -1000, -10, 18, 51, 24,
	-31, -19, -1000, -1000, -1000, 138, -1000, -1000, 146, 138,
	65, -1000, -1000, 107, -1000, -1000, 91, -50, -1000, -1000,
	-1000, -1000, -1000, 164, -1000, -1000, -1000, -1000, -1000, 80,
	-16, 169, -1000, -50, -1000, -1000, -1000, -50, -1000, -1000,
	-1000, -1000, -1000, -1000, -50, -1000, -1000, -50, -50, -50,
	-1000, -1000, -50, 138, 9, -25, -1, 43, -1000, -1000,
	-1000, -1000, -50, -1000, -1000, -50, -1000, -1000, -1000, -1000,
	24, -16, 138, -1000, 39, 179, -1000, -1000, 94, -16,
	-1000, 58, 138, 138, 54, 138, 138, 138, 68, -39,
	89, -1000, -1000, -38, -1000, 138, 138, -1000, 89, -1000,
	-1000, -1000, -1000, 160, 61, 40, -1000, -1000, 127, 51,
	-1000, -1000, -1, -1, 65, -1000, -1000, -1000, 138, -1000,
	107, -1000, -37, -1000, -1000, -1000, -1000, 89, 158, -1000,
	-1000,
}
var mtailPgo = [...]int{

	0, 42, 218, 16, 11, 217, 216, 137, 2, 3,
	9, 183, 1, 215, 19, 10, 0, 14, 214, 6,
	64, 8, 213, 212, 211, 210, 7, 12, 209, 208,
	202, 201, 200, 197, 195, 194, 193, 191, 172, 158,
	154, 147, 120, 15, 17, 116,
}
var mtailR1 = [...]int{

	0, 42, 1, 1, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 5, 5, 5, 6, 6, 4, 7,
	13, 13, 13, 17, 17, 17, 17, 37, 37, 16,
	16, 36, 36, 36, 14, 14, 34, 34, 34, 34,
	34, 34, 15, 15, 35, 35, 10, 10, 27, 27,
	27, 40, 40, 21, 20, 20, 20, 38, 38, 9,
	9, 39, 39, 39, 39, 12, 12, 11, 11, 41,
	41, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	18, 18, 19, 3, 3, 26, 22, 33, 33, 23,
	23, 23, 23, 29, 29, 29, 29, 31, 32, 32,
	32, 32, 30, 24, 25, 28, 28, 44, 45, 43,
	43,
}
var mtailR2 = [...]int{

	0, 1, 0, 2, 1, 1, 1, 1, 1, 1,
	1, 3, 1, 4, 2, 2, 1, 2, 3, 1,
	1, 4, 4, 1, 1, 4, 4, 1, 1, 1,
	4, 1, 1, 1, 1, 4, 1, 1, 1, 1,
	1, 1, 1, 4, 1, 1, 1, 4, 1, 4,
	4, 1, 1, 1, 1, 4, 4, 1, 1, 1,
	4, 1, 1, 1, 1, 1, 2, 1, 2, 1,
	1, 1, 3, 4, 1, 1, 1, 3, 1, 1,
	1, 4, 1, 1, 3, 5, 3, 0, 1, 2,
	2, 1, 1, 1, 1, 1, 1, 2, 1, 1,
	3, 3, 2, 4, 3, 4, 2, 0, 0, 0,
	1,
}
var mtailChk = [...]int{

	-1000, -42, -1, -2, -5, -6, -22, -24, -25, -28,
	16, 12, 4, -17, 17, 63, -7, -33, -44, 15,
	-16, -27, -13, 13, -14, -21, -8, -12, -15, -20,
	-18, 19, 22, 23, 21, 58, 26, 27, -11, 48,
	-10, -26, -19, -9, 24, -19, -4, -37, 56, 49,
	50, -4, 63, -29, 5, 6, 7, 8, 31, 14,
	25, -11, -8, -36, 45, 47, 46, -34, 39, 40,
	41, 42, 43, 44, -40, 54, 55, 52, 51, -35,
	37, 38, 35, 60, 58, -7, -17, -44, -41, 29,
	30, -12, -38, 35, 34, -39, 33, 31, 32, 36,
	-20, 18, -43, 63, -1, -23, 24, 21, -45, 24,
	-4, 9, -43, -43, -43, -43, -43, -43, -43, -3,
	-16, -12, 59, -3, 59, -43, -43, -4, -16, -27,
	57, -31, -30, 11, 10, 20, -4, 28, -14, -15,
	-21, -8, -17, -17, -10, -26, -19, 61, 62, 59,
	-9, -12, -32, 24, 21, 21, 31, -16, 62, 24,
	21,
}
var mtailDef = [...]int{

	2, -2, -2, 3, 4, 5, 6, 7, 8, 9,
	10, 0, 12, 20, 0, 16, 0, 0, 0, 0,
	23, 24, 19, 88, 29, 48, 67, 59, 34, 53,
	71, 0, 74, 75, 76, 107, 78, 79, 65, 0,
	42, 54, 80, 46, 82, 107, 14, 109, 2, 27,
	28, 15, 17, 0, 93, 94, 95, 96, 108, 0,
	0, 106, 67, 109, 31, 32, 33, 109, 36, 37,
	38, 39, 40, 41, 109, 51, 52, 109, 109, 109,
	44, 45, 109, 0, 0, 0, 20, 0, 68, 69,
	70, 66, 109, 57, 58, 109, 61, 62, 63, 64,
	11, 0, 107, 110, -2, 86, 91, 92, 0, 0,
	104, 0, 0, 0, 107, 107, 107, 0, 107, 0,
	83, 59, 72, 0, 77, 0, 0, 13, 25, 26,
	18, 89, 90, 0, 0, 0, 103, 105, 30, 35,
	49, 50, 21, 22, 43, 55, 56, 81, 0, 73,
	47, 60, 97, 98, 99, 102, 85, 84, 0, 100,
	101,
}
var mtailTok1 = [...]int{

	1,
}
var mtailTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63,
}
var mtailTok3 = [...]int{
	0,
}

var mtailErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{
	{108, 4, "unexpected end of file"},
}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	mtailDebug        = 0
	mtailErrorVerbose = false
)

type mtailLexer interface {
	Lex(lval *mtailSymType) int
	Error(s string)
}

type mtailParser interface {
	Parse(mtailLexer) int
	Lookahead() int
}

type mtailParserImpl struct {
	lval  mtailSymType
	stack [mtailInitialStackSize]mtailSymType
	char  int
}

func (p *mtailParserImpl) Lookahead() int {
	return p.char
}

func mtailNewParser() mtailParser {
	return &mtailParserImpl{}
}

const mtailFlag = -1000

func mtailTokname(c int) string {
	if c >= 1 && c-1 < len(mtailToknames) {
		if mtailToknames[c-1] != "" {
			return mtailToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func mtailStatname(s int) string {
	if s >= 0 && s < len(mtailStatenames) {
		if mtailStatenames[s] != "" {
			return mtailStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func mtailErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !mtailErrorVerbose {
		return "syntax error"
	}

	for _, e := range mtailErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + mtailTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := mtailPact[state]
	for tok := TOKSTART; tok-1 < len(mtailToknames); tok++ {
		if n := base + tok; n >= 0 && n < mtailLast && mtailChk[mtailAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if mtailDef[state] == -2 {
		i := 0
		for mtailExca[i] != -1 || mtailExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; mtailExca[i] >= 0; i += 2 {
			tok := mtailExca[i]
			if tok < TOKSTART || mtailExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if mtailExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += mtailTokname(tok)
	}
	return res
}

func mtaillex1(lex mtailLexer, lval *mtailSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = mtailTok1[0]
		goto out
	}
	if char < len(mtailTok1) {
		token = mtailTok1[char]
		goto out
	}
	if char >= mtailPrivate {
		if char < mtailPrivate+len(mtailTok2) {
			token = mtailTok2[char-mtailPrivate]
			goto out
		}
	}
	for i := 0; i < len(mtailTok3); i += 2 {
		token = mtailTok3[i+0]
		if token == char {
			token = mtailTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = mtailTok2[1] /* unknown char */
	}
	if mtailDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", mtailTokname(token), uint(char))
	}
	return char, token
}

func mtailParse(mtaillex mtailLexer) int {
	return mtailNewParser().Parse(mtaillex)
}

func (mtailrcvr *mtailParserImpl) Parse(mtaillex mtailLexer) int {
	var mtailn int
	var mtailVAL mtailSymType
	var mtailDollar []mtailSymType
	_ = mtailDollar // silence set and not used
	mtailS := mtailrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	mtailstate := 0
	mtailrcvr.char = -1
	mtailtoken := -1 // mtailrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		mtailstate = -1
		mtailrcvr.char = -1
		mtailtoken = -1
	}()
	mtailp := -1
	goto mtailstack

ret0:
	return 0

ret1:
	return 1

mtailstack:
	/* put a state and value onto the stack */
	if mtailDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", mtailTokname(mtailtoken), mtailStatname(mtailstate))
	}

	mtailp++
	if mtailp >= len(mtailS) {
		nyys := make([]mtailSymType, len(mtailS)*2)
		copy(nyys, mtailS)
		mtailS = nyys
	}
	mtailS[mtailp] = mtailVAL
	mtailS[mtailp].yys = mtailstate

mtailnewstate:
	mtailn = mtailPact[mtailstate]
	if mtailn <= mtailFlag {
		goto mtaildefault /* simple state */
	}
	if mtailrcvr.char < 0 {
		mtailrcvr.char, mtailtoken = mtaillex1(mtaillex, &mtailrcvr.lval)
	}
	mtailn += mtailtoken
	if mtailn < 0 || mtailn >= mtailLast {
		goto mtaildefault
	}
	mtailn = mtailAct[mtailn]
	if mtailChk[mtailn] == mtailtoken { /* valid shift */
		mtailrcvr.char = -1
		mtailtoken = -1
		mtailVAL = mtailrcvr.lval
		mtailstate = mtailn
		if Errflag > 0 {
			Errflag--
		}
		goto mtailstack
	}

mtaildefault:
	/* default state action */
	mtailn = mtailDef[mtailstate]
	if mtailn == -2 {
		if mtailrcvr.char < 0 {
			mtailrcvr.char, mtailtoken = mtaillex1(mtaillex, &mtailrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if mtailExca[xi+0] == -1 && mtailExca[xi+1] == mtailstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			mtailn = mtailExca[xi+0]
			if mtailn < 0 || mtailn == mtailtoken {
				break
			}
		}
		mtailn = mtailExca[xi+1]
		if mtailn < 0 {
			goto ret0
		}
	}
	if mtailn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			mtaillex.Error(mtailErrorMessage(mtailstate, mtailtoken))
			Nerrs++
			if mtailDebug >= 1 {
				__yyfmt__.Printf("%s", mtailStatname(mtailstate))
				__yyfmt__.Printf(" saw %s\n", mtailTokname(mtailtoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for mtailp >= 0 {
				mtailn = mtailPact[mtailS[mtailp].yys] + mtailErrCode
				if mtailn >= 0 && mtailn < mtailLast {
					mtailstate = mtailAct[mtailn] /* simulate a shift of "error" */
					if mtailChk[mtailstate] == mtailErrCode {
						goto mtailstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if mtailDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", mtailS[mtailp].yys)
				}
				mtailp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if mtailDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", mtailTokname(mtailtoken))
			}
			if mtailtoken == mtailEofCode {
				goto ret1
			}
			mtailrcvr.char = -1
			mtailtoken = -1
			goto mtailnewstate /* try again in the same state */
		}
	}

	/* reduction by production mtailn */
	if mtailDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", mtailn, mtailStatname(mtailstate))
	}

	mtailnt := mtailn
	mtailpt := mtailp
	_ = mtailpt // guard against "declared and not used"

	mtailp -= mtailR2[mtailn]
	// mtailp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if mtailp+1 >= len(mtailS) {
		nyys := make([]mtailSymType, len(mtailS)*2)
		copy(nyys, mtailS)
		mtailS = nyys
	}
	mtailVAL = mtailS[mtailp+1]

	/* consult goto table to find next state */
	mtailn = mtailR1[mtailn]
	mtailg := mtailPgo[mtailn]
	mtailj := mtailg + mtailS[mtailp].yys + 1

	if mtailj >= mtailLast {
		mtailstate = mtailAct[mtailg]
	} else {
		mtailstate = mtailAct[mtailj]
		if mtailChk[mtailstate] != -mtailn {
			mtailstate = mtailAct[mtailg]
		}
	}
	// dummy call; replaced with literal code
	switch mtailnt {

	case 1:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:81
		{
			mtaillex.(*parser).root = mtailDollar[1].n
		}
	case 2:
		mtailDollar = mtailS[mtailpt-0 : mtailpt+1]
//line parser.y:88
		{
			mtailVAL.n = &stmtlistNode{}
		}
	case 3:
		mtailDollar = mtailS[mtailpt-2 : mtailpt+1]
//line parser.y:92
		{
			mtailVAL.n = mtailDollar[1].n
			if mtailDollar[2].n != nil {
				mtailVAL.n.(*stmtlistNode).children = append(mtailVAL.n.(*stmtlistNode).children, mtailDollar[2].n)
			}
		}
	case 4:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:102
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 5:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:104
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 6:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:106
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 7:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:108
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 8:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:110
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 9:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:112
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 10:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:114
		{
			mtailVAL.n = &nextNode{tokenpos(mtaillex)}
		}
	case 11:
		mtailDollar = mtailS[mtailpt-3 : mtailpt+1]
//line parser.y:118
		{
			mtailVAL.n = &patternFragmentDefNode{id: mtailDollar[2].n, expr: mtailDollar[3].n}
		}
	case 12:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:122
		{
			mtailVAL.n = &errorNode{tokenpos(mtaillex), mtailDollar[1].text}
		}
	case 13:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:129
		{
			mtailVAL.n = &condNode{mtailDollar[1].n, mtailDollar[2].n, mtailDollar[4].n, nil}
		}
	case 14:
		mtailDollar = mtailS[mtailpt-2 : mtailpt+1]
//line parser.y:133
		{
			if mtailDollar[1].n != nil {
				mtailVAL.n = &condNode{mtailDollar[1].n, mtailDollar[2].n, nil, nil}
			} else {
				mtailVAL.n = mtailDollar[2].n
			}
		}
	case 15:
		mtailDollar = mtailS[mtailpt-2 : mtailpt+1]
//line parser.y:141
		{
			o := &otherwiseNode{tokenpos(mtaillex)}
			mtailVAL.n = &condNode{o, mtailDollar[2].n, nil, nil}
		}
	case 16:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:149
		{
			mtailVAL.n = nil
		}
	case 17:
		mtailDollar = mtailS[mtailpt-2 : mtailpt+1]
//line parser.y:151
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 18:
		mtailDollar = mtailS[mtailpt-3 : mtailpt+1]
//line parser.y:156
		{
			mtailVAL.n = mtailDollar[2].n
		}
	case 19:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:163
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 20:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:168
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 21:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:172
		{
			mtailVAL.n = &binaryExprNode{lhs: mtailDollar[1].n, rhs: mtailDollar[4].n, op: mtailDollar[2].op}
		}
	case 22:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:176
		{
			mtailVAL.n = &binaryExprNode{lhs: mtailDollar[1].n, rhs: mtailDollar[4].n, op: mtailDollar[2].op}
		}
	case 23:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:183
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 24:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:185
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 25:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:187
		{
			mtailVAL.n = &binaryExprNode{lhs: mtailDollar[1].n, rhs: mtailDollar[4].n, op: mtailDollar[2].op}
		}
	case 26:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:191
		{
			mtailVAL.n = &binaryExprNode{lhs: mtailDollar[1].n, rhs: mtailDollar[4].n, op: mtailDollar[2].op}
		}
	case 27:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:198
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 28:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:200
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 29:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:205
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 30:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:207
		{
			mtailVAL.n = &binaryExprNode{lhs: mtailDollar[1].n, rhs: mtailDollar[4].n, op: mtailDollar[2].op}
		}
	case 31:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:214
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 32:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:216
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 33:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:218
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 34:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:223
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 35:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:225
		{
			mtailVAL.n = &binaryExprNode{lhs: mtailDollar[1].n, rhs: mtailDollar[4].n, op: mtailDollar[2].op}
		}
	case 36:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:232
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 37:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:234
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 38:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:236
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 39:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:238
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 40:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:240
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 41:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:242
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 42:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:247
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 43:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:249
		{
			mtailVAL.n = &binaryExprNode{lhs: mtailDollar[1].n, rhs: mtailDollar[4].n, op: mtailDollar[2].op}
		}
	case 44:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:256
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 45:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:258
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 46:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:263
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 47:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:265
		{
			mtailVAL.n = &binaryExprNode{lhs: mtailDollar[1].n, rhs: mtailDollar[4].n, op: mtailDollar[2].op}
		}
	case 48:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:272
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 49:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:274
		{
			mtailVAL.n = &binaryExprNode{lhs: mtailDollar[1].n, rhs: mtailDollar[4].n, op: mtailDollar[2].op}
		}
	case 50:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:278
		{
			mtailVAL.n = &binaryExprNode{lhs: mtailDollar[1].n, rhs: mtailDollar[4].n, op: mtailDollar[2].op}
		}
	case 51:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:285
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 52:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:287
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 53:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:292
		{
			mtailVAL.n = &patternExprNode{expr: mtailDollar[1].n}
		}
	case 54:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:299
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 55:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:301
		{
			mtailVAL.n = &binaryExprNode{lhs: mtailDollar[1].n, rhs: mtailDollar[4].n, op: CONCAT}
		}
	case 56:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:305
		{
			mtailVAL.n = &binaryExprNode{lhs: mtailDollar[1].n, rhs: mtailDollar[4].n, op: CONCAT}
		}
	case 57:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:312
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 58:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:314
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 59:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:319
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 60:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:321
		{
			mtailVAL.n = &binaryExprNode{lhs: mtailDollar[1].n, rhs: mtailDollar[4].n, op: mtailDollar[2].op}
		}
	case 61:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:328
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 62:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:330
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 63:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:332
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 64:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:334
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 65:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:339
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 66:
		mtailDollar = mtailS[mtailpt-2 : mtailpt+1]
//line parser.y:341
		{
			mtailVAL.n = &unaryExprNode{pos: tokenpos(mtaillex), expr: mtailDollar[2].n, op: mtailDollar[1].op}
		}
	case 67:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:348
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 68:
		mtailDollar = mtailS[mtailpt-2 : mtailpt+1]
//line parser.y:350
		{
			mtailVAL.n = &unaryExprNode{pos: tokenpos(mtaillex), expr: mtailDollar[1].n, op: mtailDollar[2].op}
		}
	case 69:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:357
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 70:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:359
		{
			mtailVAL.op = mtailDollar[1].op
		}
	case 71:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:364
		{
			mtailVAL.n = mtailDollar[1].n
		}
	case 72:
		mtailDollar = mtailS[mtailpt-3 : mtailpt+1]
//line parser.y:366
		{
			mtailVAL.n = &builtinNode{pos: tokenpos(mtaillex), name: mtailDollar[1].text, args: nil}
		}
	case 73:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:370
		{
			mtailVAL.n = &builtinNode{pos: tokenpos(mtaillex), name: mtailDollar[1].text, args: mtailDollar[3].n}
		}
	case 74:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:374
		{
			mtailVAL.n = &caprefNode{tokenpos(mtaillex), mtailDollar[1].text, false, nil}
		}
	case 75:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:378
		{
			mtailVAL.n = &caprefNode{tokenpos(mtaillex), mtailDollar[1].text, true, nil}
		}
	case 76:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:382
		{
			mtailVAL.n = &stringConstNode{tokenpos(mtaillex), mtailDollar[1].text}
		}
	case 77:
		mtailDollar = mtailS[mtailpt-3 : mtailpt+1]
//line parser.y:386
		{
			mtailVAL.n = mtailDollar[2].n
		}
	case 78:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:390
		{
			mtailVAL.n = &intConstNode{tokenpos(mtaillex), mtailDollar[1].intVal}
		}
	case 79:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:394
		{
			mtailVAL.n = &floatConstNode{tokenpos(mtaillex), mtailDollar[1].floatVal}
		}
	case 80:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:401
		{
			mtailVAL.n = &indexedExprNode{lhs: mtailDollar[1].n, index: &exprlistNode{}}
		}
	case 81:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:405
		{
			mtailVAL.n = mtailDollar[1].n
			mtailVAL.n.(*indexedExprNode).index.(*exprlistNode).children = append(
				mtailVAL.n.(*indexedExprNode).index.(*exprlistNode).children,
				mtailDollar[3].n.(*exprlistNode).children...)
		}
	case 82:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:415
		{
			mtailVAL.n = &idNode{tokenpos(mtaillex), mtailDollar[1].text, nil, false}
		}
	case 83:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:422
		{
			mtailVAL.n = &exprlistNode{}
			mtailVAL.n.(*exprlistNode).children = append(mtailVAL.n.(*exprlistNode).children, mtailDollar[1].n)
		}
	case 84:
		mtailDollar = mtailS[mtailpt-3 : mtailpt+1]
//line parser.y:427
		{
			mtailVAL.n = mtailDollar[1].n
			mtailVAL.n.(*exprlistNode).children = append(mtailVAL.n.(*exprlistNode).children, mtailDollar[3].n)
		}
	case 85:
		mtailDollar = mtailS[mtailpt-5 : mtailpt+1]
//line parser.y:435
		{
			mp := markedpos(mtaillex)
			tp := tokenpos(mtaillex)
			pos := MergePosition(&mp, &tp)
			mtailVAL.n = &patternConstNode{pos: *pos, pattern: mtailDollar[4].text}
		}
	case 86:
		mtailDollar = mtailS[mtailpt-3 : mtailpt+1]
//line parser.y:445
		{
			mtailVAL.n = mtailDollar[3].n
			d := mtailVAL.n.(*declNode)
			d.kind = mtailDollar[2].kind
			d.hidden = mtailDollar[1].flag
		}
	case 87:
		mtailDollar = mtailS[mtailpt-0 : mtailpt+1]
//line parser.y:455
		{
			mtailVAL.flag = false
		}
	case 88:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:459
		{
			mtailVAL.flag = true
		}
	case 89:
		mtailDollar = mtailS[mtailpt-2 : mtailpt+1]
//line parser.y:466
		{
			mtailVAL.n = mtailDollar[1].n
			mtailVAL.n.(*declNode).keys = mtailDollar[2].texts
		}
	case 90:
		mtailDollar = mtailS[mtailpt-2 : mtailpt+1]
//line parser.y:471
		{
			mtailVAL.n = mtailDollar[1].n
			mtailVAL.n.(*declNode).exportedName = mtailDollar[2].text
		}
	case 91:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:476
		{
			mtailVAL.n = &declNode{pos: tokenpos(mtaillex), name: mtailDollar[1].text}
		}
	case 92:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:480
		{
			mtailVAL.n = &declNode{pos: tokenpos(mtaillex), name: mtailDollar[1].text}
		}
	case 93:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:487
		{
			mtailVAL.kind = metrics.Counter
		}
	case 94:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:491
		{
			mtailVAL.kind = metrics.Gauge
		}
	case 95:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:495
		{
			mtailVAL.kind = metrics.Timer
		}
	case 96:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:499
		{
			mtailVAL.kind = metrics.Text
		}
	case 97:
		mtailDollar = mtailS[mtailpt-2 : mtailpt+1]
//line parser.y:506
		{
			mtailVAL.texts = mtailDollar[2].texts
		}
	case 98:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:513
		{
			mtailVAL.texts = make([]string, 0)
			mtailVAL.texts = append(mtailVAL.texts, mtailDollar[1].text)
		}
	case 99:
		mtailDollar = mtailS[mtailpt-1 : mtailpt+1]
//line parser.y:518
		{
			mtailVAL.texts = make([]string, 0)
			mtailVAL.texts = append(mtailVAL.texts, mtailDollar[1].text)
		}
	case 100:
		mtailDollar = mtailS[mtailpt-3 : mtailpt+1]
//line parser.y:523
		{
			mtailVAL.texts = mtailDollar[1].texts
			mtailVAL.texts = append(mtailVAL.texts, mtailDollar[3].text)
		}
	case 101:
		mtailDollar = mtailS[mtailpt-3 : mtailpt+1]
//line parser.y:528
		{
			mtailVAL.texts = mtailDollar[1].texts
			mtailVAL.texts = append(mtailVAL.texts, mtailDollar[3].text)
		}
	case 102:
		mtailDollar = mtailS[mtailpt-2 : mtailpt+1]
//line parser.y:536
		{
			mtailVAL.text = mtailDollar[2].text
		}
	case 103:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:543
		{
			mtailVAL.n = &decoDefNode{pos: markedpos(mtaillex), name: mtailDollar[3].text, block: mtailDollar[4].n}
		}
	case 104:
		mtailDollar = mtailS[mtailpt-3 : mtailpt+1]
//line parser.y:550
		{
			mtailVAL.n = &decoNode{markedpos(mtaillex), mtailDollar[2].text, mtailDollar[3].n, nil, nil}
		}
	case 105:
		mtailDollar = mtailS[mtailpt-4 : mtailpt+1]
//line parser.y:557
		{
			mtailVAL.n = &delNode{pos: tokenpos(mtaillex), n: mtailDollar[2].n, expiry: mtailDollar[4].duration}
		}
	case 106:
		mtailDollar = mtailS[mtailpt-2 : mtailpt+1]
//line parser.y:561
		{
			mtailVAL.n = &delNode{pos: tokenpos(mtaillex), n: mtailDollar[2].n}
		}
	case 107:
		mtailDollar = mtailS[mtailpt-0 : mtailpt+1]
//line parser.y:571
		{
			glog.V(2).Infof("position marked at %v", tokenpos(mtaillex))
			mtaillex.(*parser).pos = tokenpos(mtaillex)
		}
	case 108:
		mtailDollar = mtailS[mtailpt-0 : mtailpt+1]
//line parser.y:581
		{
			mtaillex.(*parser).inRegex()
		}
	}
	goto mtailstack /* stack new state and value */
}