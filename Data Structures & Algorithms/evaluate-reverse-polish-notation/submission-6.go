func evalRPN(tokens []string) int {
	stk := []int{}
	for i := 0; i < len(tokens); i++{
		if isOperator(tokens[i]){
			b := stk[len(stk)-1]
			a := stk[len(stk)-2]
			stk = stk[:len(stk)-2]
			switch tokens[i]{
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
			val, _ := strconv.Atoi(tokens[i])
			stk = append(stk, val)
		}
	}
	return stk[0]
}

func isOperator(c string) bool{
	if c == "+" || c == "-" || c == "*" || c == "/"{
		return true
	}
	return false
}