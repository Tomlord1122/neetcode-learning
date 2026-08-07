func generateParenthesis(n int) []string {
	res := []string{}
	cur := []byte{}
	var backtrack func(open, close int)
	backtrack = func(open, close int){
		if open == n && close == n{
			res = append(res, string(cur))
			return
		}
		if open < n{
			cur = append(cur, '(')
			backtrack(open+1, close)
			cur = cur[:len(cur)-1]
		}
		if close < open{
			cur = append(cur, ')')
			backtrack(open, close+1)
			cur = cur[:len(cur)-1]
		}
	}
	backtrack(0, 0)
	return res
}
