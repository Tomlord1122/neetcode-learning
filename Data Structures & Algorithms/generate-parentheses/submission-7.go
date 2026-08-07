func generateParenthesis(n int) []string {
	res := []string{}

	var backtrack func(open, close int, cur string)
	backtrack = func(open, close int, cur string){
		if open == n && close == n{
			res = append(res, cur)
		}
		if open < n{
			backtrack(open+1, close, cur+"(")
		}
		if close < open{
			backtrack(open, close+1, cur+")")
		}
	}

	backtrack(0, 0, "")
	return res
}
