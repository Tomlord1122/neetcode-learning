func generateParenthesis(n int) []string {
	res := []string{}
	stk := []byte{}
	var dfs func(open, close int)
	dfs = func(open, close int){
		if open == n && close == n{
			res = append(res, string(stk))
			return
		}

		if open < n{
			stk = append(stk, '(')
			dfs(open+1, close)
			stk = stk[:len(stk)-1]
		}
		if close < open{
			stk = append(stk, ')')
			dfs(open, close+1)
			stk = stk[:len(stk)-1]
		}
	}
	dfs(0, 0)
	return res
}
