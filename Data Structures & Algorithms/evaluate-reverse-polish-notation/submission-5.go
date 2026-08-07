func evalRPN(tokens []string) int {
	stk := []int{}
	for _, token := range tokens{
		if isOperator(token){
			b := stk[len(stk)-1]
			a := stk[len(stk)-2]
			stk = stk[:len(stk)-2]
			switch token{
				case "+":
					stk = append(stk, a+b)
				case "-":
					stk = append(stk, a-b)
				case "*":
					stk = append(stk, a*b)
				case "/":
					stk = append(stk, a/b)
			}
		} else {
			val, _ := strconv.Atoi(token)
			stk = append(stk, val)
		}
	}
	return stk[0]
}

func isOperator (char string) bool{
	if char == "+" || char == "-" || char == "*" || char == "/"{
		return true
	}
	return false
}
