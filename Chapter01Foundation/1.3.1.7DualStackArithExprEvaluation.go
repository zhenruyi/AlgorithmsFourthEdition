package Chapter01Foundation

import (
	"container/list"
	"strconv"
)

/*func main() {
	expressions := "( 1.123 + ( ( 2 + 3 ) * ( 4 * 5 ) ) )"
	res := Evaluate(expressions)
	fmt.Println(res)
}*/


// Dijkstra的双栈算术表达式求值算法
func Evaluate(exp string) float64 {
	// 双栈
	ops := list.New()
	vals := list.New()
	// 处理数字
	val := ""
	for i := 0; i < len(exp); i++ {
		// 不处理空格和(
		if exp[i] == ' ' || exp[i] == '('{
			continue
		}
		// 如果是数字，则拼接
		val += string(exp[i])
		if i+1 < len(exp) && isNum(exp[i]) && isNum(exp[i+1]) {
			continue
		}
		// 算术表达式处理
		switch {
		// 数字的处理
		case len(val) > 1 || (exp[i] >= '0' && exp[i] <= '9'):
			v, _ := strconv.ParseFloat(val, 64)
			vals.PushBack(v)
		// 运算符的处理
		case exp[i] == '+' || exp[i] == '-' || exp[i] == '*' || exp[i] == '/':
			ops.PushBack(exp[i])
		// 遇到反括号，进行一次运算
		case exp[i] == ')':
			// 运算符出栈
			curOp := ops.Remove(ops.Back()).(uint8)
			// 操作数出栈
			curVal2 := vals.Remove(vals.Back()).(float64)
			curVal1 := vals.Remove(vals.Back()).(float64)
			switch curOp {
			case '+': vals.PushBack(curVal1 + curVal2)
			case '-': vals.PushBack(curVal1 - curVal2)
			case '*': vals.PushBack(curVal1 * curVal2)
			case '/': vals.PushBack(curVal1 / curVal2)
			}
		default:
		}
		val = ""
	}
	return vals.Front().Value.(float64)
}
// 判断是不是数字
func isNum(val uint8) bool {
	if val >= '0' && val <= '9' {
		return true
	}
	// 可能是浮点数
	if val == '.' {
		return true
	}
	return false
}
