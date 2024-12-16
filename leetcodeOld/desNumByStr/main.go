package main

import "fmt"

func main() {
	fmt.Println(isNumber("-1E-16"))
}

type State int
type CharType int

const (
	StateInitial         State = iota
	StateIntSign               // 整数符号
	StateInteger               // 整数
	StatePoint                 // 小数点
	StatePointWithoutInt       // 小数点前无整数
	StateDecimal               // 小数部分
	StateExp                   // 指数，e
	StateExpSign               // 指数符号部分
	StateExpInt                // 指数数字部分
	StateEnd                   // 末尾空格，结束状态
)

const (
	CharNumber CharType = iota
	CharExp
	CharPoint
	CharSign
	CharSpace
	CharIllegal
)

func chatToType(char byte) CharType {
	switch char {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return CharNumber
	case 'e', 'E':
		return CharExp
	case '.':
		return CharPoint
	case ' ':
		return CharSpace
	case '+', '-':
		return CharSign
	default:
		return CharIllegal
	}
}

func isNumber(s string) bool {
	transfer := map[State]map[CharType]State{
		StateInitial: map[CharType]State{
			CharSpace:  StateInitial,
			CharNumber: StateInteger,
			CharPoint:  StatePointWithoutInt,
			CharSign:   StateIntSign,
		},
		StateIntSign: map[CharType]State{
			CharNumber: StateInteger,
			CharPoint:  StatePointWithoutInt,
		},
		StateInteger: map[CharType]State{
			CharNumber: StateInteger,
			CharExp:    StateExp,
			CharPoint:  StatePoint,
			CharSpace:  StateEnd,
		},
		StatePoint: map[CharType]State{
			CharNumber: StateDecimal,
			CharExp:    StateExp,
			CharSpace:  StateEnd,
		},
		StatePointWithoutInt: map[CharType]State{
			CharNumber: StateDecimal,
		},
		StateDecimal: map[CharType]State{
			CharNumber: StateDecimal,
			CharExp:    StateExp,
			CharSpace:  StateEnd,
		},
		StateExp: map[CharType]State{
			CharNumber: StateExpInt,
			CharSign:   StateExpSign,
		},
		StateExpSign: map[CharType]State{
			CharNumber: StateExpInt,
		},
		StateExpInt: map[CharType]State{
			CharNumber: StateExpInt,
			CharSpace:  StateEnd,
		},
		StateEnd: map[CharType]State{
			CharSpace: StateEnd,
		},
	}
	state := StateInitial
	for i := 0; i < len(s); i++ {
		typ := chatToType(s[i])
		if _, ok := transfer[state][typ]; !ok {
			return false
		} else {
			state = transfer[state][typ]
		}
	}
	return state == StateInteger || state == StatePoint || state == StateDecimal || state == StateExpInt || state == StateEnd
}
