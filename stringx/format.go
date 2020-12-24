package stringx

import (
	"fmt"
	"github.com/hsiafan/glow/reflectx"
	"github.com/hsiafan/glow/stringx/ascii"
	"reflect"
	"strconv"
	"strings"
)

const (
	statePlain                = 0
	stateFirstCurseBegin      = 1
	stateFirstCurseEnd        = 2
	stateReadIndexOrName      = 3
	stateReadFormat           = 4
	stateReadFormatPad        = 41
	stateReadFormatFraction   = 42
	stateReadFormatNumberType = 43
	stateParamEnd             = 5
)

// Format string.
//
// Common Usage:
//  stringx.Format("{},{}", 1, 2, 3) // format using automatic field numbering
//	stringx.Format("{0},{1}", 1, 2) // format using manual field specification
// Escape: using "{{" for {, and "}}" for }
//  stringx.Format("{{") // producers "{"
// Padding: < for padding and align to left, > for padding and align to left, ^ for padding and align to center.
// A padding char can be specified before the padding sign, and a len can be specified after the sign.
//  stringx.Format("{:<2}", 1) // produces "1 "
//	stringx.Format("{:^3}", 1) // produces " 1 "
//  stringx.Format("{:0>3}", 1) // produces "001"
// Number format. b for binary, o for octal, d for decimal, x for hex with lower case, X for hex with upper case.
// f for float, and can specify a fraction num for float values.
//  stringx.Format("{:X}", 160) // produces "A0"
//  stringx.Format("{:x}", 160) // produces "a0"
//  stringx.Format("{:.2f}", 1.0) // produces "1.00", only for float values.
// Number prefix. a # char to indicate adding 0x/0o/ob prefix to number string.
//  stringx.Format("{:#x}", 160) // produces "0xa0"
func Format(pattern string, params ...interface{}) string {
	var sb Builder
	format(&sb, pattern, false, params, nil)
	return sb.String()
}

type namedParams interface {
	Get(name string) (value interface{}, exists bool)
}

// NamedParams provide named params to FormatNamed, by a map
type NamedParams map[string]interface{}

func (p NamedParams) Get(name string) (value interface{}, exists bool) {
	value, exists = p[name]
	return
}

type structParams struct {
	value reflect.Value
}

func newStructParams(value interface{}) *structParams {
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		panic("need struct params")
	}
	return &structParams{v}
}

func (s *structParams) Get(name string) (value interface{}, exists bool) {
	vv := s.value.FieldByName(name)
	if vv.IsZero() || !vv.CanInterface() {
		return nil, false
	}
	return vv.Interface(), true
}

// FormatNamed format string with named params. This func works like Format, while the format string can set a name.
// The name must begin with alphabet char, or '_', and contains only alphabet/numbers/'_' chars.
//
// Common Usage:
//  stringx.FormatNamed("{my_name},{your_name}", stringx.NamedParams {
//		"my_name": 1,
//		"your_name": "2",
//  })
// More :
//  stringx.FormatNamed("{my_name:<2},{your_name:>3#x}", stringx.NamedParams {
//		"my_name": "1",
//		"your_name": 2,
//  })
//
func FormatNamed(pattern string, params NamedParams) string {
	var sb Builder
	format(&sb, pattern, true, nil, params)
	return sb.String()
}

// FormatNamed2 format string with named params provided by struct value.
// This func works like Format, while the format string can set a name.
// The name must begin with alphabet char, or '_', and contains only alphabet/numbers/'_' chars.
//
// Common Usage:
//  myStructValue = MyStruct {
//  	Field1: "1",
//  	Field2: 2,
//  }
//  stringx.FormatNamed("{Field1},{Field2}", &myStructValue)
// More :
//  stringx.FormatNamed("{Field1:<2},{Field2::>3#x}", &myStructValue)
//
func FormatNamed2(pattern string, params interface{}) string {
	var sb Builder
	format(&sb, pattern, true, nil, newStructParams(params))
	return sb.String()
}

func format(sb *Builder, pattern string, named bool, params []interface{}, namedParams namedParams) {
	t := tokenizer{runes: []rune(pattern)}
	var state = 0
	var count = 0
	var paramIndex = 0 // the index read from format string
	var paramName = "" // the param name, used when named is true
	var hasNonIndexParam = false
	var hasIndexParam = false

	var paddingChar = ' '
	var paddingDirection rune = 0
	var paddingLen = 0

	var fractionCount = -1
	var hasNumberPrefix = false
	var numberType rune = 0

	for t.hasNext() {
		c := t.nextRune()
		switch state {
		case statePlain:
			if c == '{' {
				state = stateFirstCurseBegin
			} else if c == '}' {
				state = stateFirstCurseEnd
			} else {
				sb.WriteRune(c)
			}
		case stateFirstCurseBegin:
			if c == '{' {
				sb.WriteByte('{')
				state = statePlain
			} else if named {
				if ascii.IsAlphaBet(byte(c)) || c == '_' {
					state = stateReadIndexOrName
					t.putBack()
				} else {
					panic("invalid format at: " + pattern + ", position:" + strconv.Itoa(t.index()))
				}
			} else {
				if c == '}' {
					if hasIndexParam {
						panic("cannot switch from automatic field numbering to manual field specification")
					}
					hasNonIndexParam = true
					paramIndex = count
					count++
					state = stateParamEnd
					t.putBack()
				} else if ascii.IsDigit(byte(c)) {
					if hasNonIndexParam {
						panic("cannot switch from automatic field numbering to manual field specification")
					}
					hasIndexParam = true
					state = stateReadIndexOrName
					paramIndex = 0
					t.putBack()
				} else if c == ':' {
					state = stateReadFormat
				} else {
					panic("invalid format at: " + pattern + ", position:" + strconv.Itoa(t.index()))
				}
			}
		case stateReadIndexOrName:
			t.putBack()
			if named {
				paramName = t.nextName()
			} else {
				paramIndex = t.nextInt()
			}
			c = t.nextRune()
			if c == '}' {
				state = stateParamEnd
				t.putBack()
			} else if c == ':' {
				state = stateReadFormat
			} else {
				panic("invalid format at: " + pattern + ", position:" + strconv.Itoa(t.index()))
			}
		case stateReadFormat:
			if c == '>' || c == '<' || c == '^' {
				state = stateReadFormatPad
				t.putBack()
			} else {
				cn := t.nextRune()
				if cn == '>' || c == '<' || c == '^' {
					paddingChar = c
					state = stateReadFormatPad
					t.putBack()
				} else {
					t.putBack()
					t.putBack()
					state = stateReadFormatFraction
				}
			}
		case stateReadFormatPad:
			if c == '>' || c == '<' || c == '^' {
				paddingDirection = c
				paddingLen = t.nextInt()
				state = stateReadFormatFraction
			} else {
				panic("not padding")
			}
		case stateReadFormatFraction:
			if c == '.' {
				fractionCount = t.nextInt()
			} else {
				t.putBack()
			}
			state = stateReadFormatNumberType
		case stateReadFormatNumberType:
			if c == '#' {
				hasNumberPrefix = true
				c = t.nextRune()
			}
			if c == 'b' || c == 'd' || c == 'o' || c == 'x' || c == 'X' || c == 'f' {
				numberType = c
			} else if c == '}' {
				t.putBack()
				state = stateParamEnd
			} else {
				panic("invalid format at: " + pattern + ", position:" + strconv.Itoa(t.index()))
			}
		case stateParamEnd:
			if c != '}' {
				panic("should be }")
			}
			var str string
			var param interface{}
			if named {
				var exists bool
				param, exists = namedParams.Get(paramName)
				if !exists {
					panic("param with name " + paramName + " not exists")
				}
			} else {
				param = params[paramIndex]
			}
			var numberPrefix string
			if numberType == 'd' {
				if !reflectx.IsInt(param) {
					panic(fmt.Sprintf("non-int value use int format: %T", param))
				}
				str = fmt.Sprintf("%d", param)
			} else if numberType == 'b' {
				if !reflectx.IsInt(param) {
					panic(fmt.Sprintf("non-int value use int format: %T", param))
				}
				numberPrefix = "0b"
				str = fmt.Sprintf("%b", param)
			} else if numberType == 'o' {
				if !reflectx.IsInt(param) {
					panic(fmt.Sprintf("non-int value use int format: %T", param))
				}
				numberPrefix = "0o"
				str = fmt.Sprintf("%o", param)
			} else if numberType == 'x' {
				if !reflectx.IsInt(param) {
					panic(fmt.Sprintf("non-int value use int format: %T", param))
				}
				numberPrefix = "0x"
				str = fmt.Sprintf("%x", param)
			} else if numberType == 'X' {
				if !reflectx.IsInt(param) {
					panic(fmt.Sprintf("non-int value use int format: %T", param))
				}
				numberPrefix = "0x"
				str = fmt.Sprintf("%X", param)
			} else if numberType == 'f' {
				if !reflectx.IsFloat(param) {
					panic(fmt.Sprintf("non-float value use float format: %T", param))
				}
				if fractionCount >= 0 {
					str = fmt.Sprintf("%."+strconv.Itoa(fractionCount)+"f", param)
				} else {
					str = fmt.Sprintf("%f", param)
				}
			} else {
				str = fmt.Sprintf("%v", param)
			}

			if hasNumberPrefix {
				if numberPrefix == "" {
					panic("only int value with hex/octal/binary format can have leading 0x/0b/0o prefix")
				} else {
					sb.WriteString(numberPrefix)
					paddingLen = paddingLen - len(numberPrefix)
				}
			}
			// padding
			if paddingDirection == '>' {
				for i := len(str); i < paddingLen; i++ {
					sb.WriteRune(paddingChar)
				}
				sb.WriteString(str)
			} else if paddingDirection == '<' {
				sb.WriteString(str)
				for i := len(str); i < paddingLen; i++ {
					sb.WriteRune(paddingChar)
				}
			} else if paddingDirection == '^' {
				toPad := paddingLen - len(str)
				for i := 0; i < toPad/2; i++ {
					sb.WriteRune(paddingChar)
				}
				sb.WriteString(str)
				for i := 0; i < toPad-toPad/2; i++ {
					sb.WriteRune(paddingChar)
				}
			} else {
				sb.WriteString(str)
			}
			paddingChar = ' '
			paddingDirection = 0
			paddingLen = 0

			fractionCount = -1
			hasNumberPrefix = false
			numberType = 0

			state = statePlain
		case stateFirstCurseEnd:
			if c == '}' {
				sb.WriteByte('}')
				state = statePlain
			} else {
				panic("single '}' is not allowed")
			}
		}
	}
	if state != statePlain {
		panic("invalid format pattern: " + pattern)
	}
}

type tokenizer struct {
	runes []rune
	idx   int
}

func (t *tokenizer) hasNext() bool {
	return t.idx < len(t.runes)
}

func (t *tokenizer) nextRune() rune {
	r := t.runes[t.idx]
	t.idx++
	return r
}

func (t *tokenizer) putBack() {
	t.idx--
}

func (t *tokenizer) index() int {
	return t.idx
}

func (t *tokenizer) nextInt() int {
	number := 0
	for t.hasNext() {
		c := t.nextRune()
		if !ascii.IsDigit(byte(c)) {
			t.putBack()
			break
		}
		number = number*10 + int(c-'0')
	}
	return number
}

func (t *tokenizer) nextName() string {
	var sb strings.Builder
	for t.hasNext() {
		c := t.nextRune()
		if !ascii.IsDigit(byte(c)) && !ascii.IsAlphaBet(byte(c)) && c != '_' {
			t.putBack()
			break
		}
		sb.WriteRune(c)
	}
	return sb.String()
}
