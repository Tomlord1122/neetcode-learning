func generateParenthesis(n int) []string {
	res := []string{}
	stk := []byte{}

	var bt func(open, close int)
	bt = func(open, close int){
		if open == n && close == n{
			res = append(res, string(stk))
		}
		if open < n{
			stk = append(stk, '(')
			bt(open+1, close)
			stk = stk[:len(stk)-1]
		}
		if close < open{
			stk = append(stk, ')')
			bt(open, close+1)
			stk = stk[:len(stk)-1]
		}
	}
	bt(0, 0)
	return res
}
