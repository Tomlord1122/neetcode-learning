func evalRPN(tokens []string) int {
	stk := []int{}
	for _, token := range tokens{
		if token != "+" && token != "-" && token != "*" && token != "/"{
			val, _ := strconv.Atoi(token)
			stk = append(stk, val)
		} else {
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
		}
	}
	return stk[0]
}
