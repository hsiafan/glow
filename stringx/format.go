package stringx

import (
	"fmt"
	"github.com/hsiafan/glow/reflectx"
	"github.com/hsiafan/glow/stringx/ascii"
	"strconv"
	"strings"
)

const (
	statePlain                = 0
	stateFirstCurseBegin      = 1
	stateFirstCurseEnd        = 2
	stateReadIndex            = 3
	stateReadName             = 3
	stateReadFormat           = 4
	stateReadFormatPad        = 41
	stateReadFormatFraction   = 42
	stateReadFormatNumberType = 43
	stateParamEnd             = 5
)

// Format string.
// Common Usage:
//  stringx.Format("{},{}", 1, 2, 3) // format using automatic field numbering
//	stringx.Format("{0},{1}", 1, 2) // format using manual field specification
// Escape: using "{{" for {, and "}}" for }
//  stringx.Format("{{") // producers "{"
// Padding: < for padding and align to left, > for padding and align to left, ^ for padding and align to center.
// A padding char can be specified before the padding sign, and a len can be specified after the sign.
//  stringx.Format("{:<2}", 1) // producers "1 "
//	stringx.Format("{:^3}", 1) // producers " 1 "
//  stringx.Format("{:0>3}", 1) // producers "001"
// Number Format. b for binary, o for octal, d for decimal, x for hex with lower case, X for hex with upper case.
// f for float, and can specify a fraction num.
//  stringx.Format("{:X}", 160) // producers "A0"
//  stringx.Format("{:x}", 160) // producers "a0"
//  Format("{:.2f}", 1.0) // producer "1.00"
func Format(pattern string, params ...interface{}) string {
	t := tokenizer{runes: []rune(pattern)}
	var state = 0
	var sb Builder
	var count = 0
	var paramIndex = 0 // the index read from format string
	var hasParam = false
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
			} else if c == '}' {
				if hasIndexParam {
					panic("cannot switch from automatic field numbering to manual field specification")
				}
				hasParam = true
				paramIndex = count
				count++
				state = stateParamEnd
				t.putBack()
			} else if ascii.IsDigit(byte(c)) {
				if hasParam {
					panic("cannot switch from automatic field numbering to manual field specification")
				}
				hasIndexParam = true
				state = stateReadIndex
				paramIndex = 0
				t.putBack()
			} else if c == ':' {
				state = stateReadFormat
			} else {
				panic("invalid format at: " + pattern + ", position:" + strconv.Itoa(t.index()))
			}
		case stateReadIndex:
			t.putBack()
			paramIndex = t.nextInt()
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
			param := params[paramIndex]
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
					panic("value and format do not has leading 0x/0b/0o")
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
	return sb.String()
}

type NamedParams map[string]interface{}

// Format string with named params. This func works like Format, while the format string can set a name.
// The name must begin with alphabet char, or '_', and contains only alphabet/numbers/'_' chars.
//
// Common Usage:
//  stringx.FormatNamed("{my_name},{your_name}", stringx.NamedParams {
//		"my_name": 1,
//		"your_name": "2",
//  })
// More :
//  stringx.FormatNamed("{my_name:<2},{your_name:X>3}", stringx.NamedParams {
//		"my_name": 1,
//		"your_name": "2",
//  })
//
func FormatNamed(pattern string, params NamedParams) string {
	t := tokenizer{runes: []rune(pattern)}
	var state = 0
	var sb Builder
	var paramName string // the index read from format string
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
			} else if ascii.IsAlphaBet(byte(c)) || c == '_' {
				state = stateReadName
				t.putBack()
			} else {
				panic("invalid format at: " + pattern + ", position:" + strconv.Itoa(t.index()))
			}
		case stateReadName:
			t.putBack()
			paramName = t.nextName()
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
			param, ok := params[paramName]
			if !ok {
				panic("param with name: " + paramName + "not found")
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
					panic("value and format do not has leading 0x/0b/0o")
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
	return sb.String()
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
