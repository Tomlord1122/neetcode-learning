func evalRPN(tokens []string) int {
	stk := []int{}
	for _, token := range tokens{
		if token == "+" || token == "-" || token == "*" || token == "/"{
			n := len(stk)
			b, a := stk[n-1], stk[n-2]
			stk = stk[:n-2]
			switch token{
				case "+":
					stk = append(stk, a + b)
				case "-":
					stk = append(stk, a - b)
				case "*":
					stk = append(stk, a * b)
				case "/":
					stk = append(stk, a / b)
			}
		} else {
			v, _ := strconv.Atoi(token)
			stk = append(stk, v)
		}
	}
	return stk[0]
}
