func generateParenthesis(n int) []string {
	res := []string{}
	cur := []byte{}
	var bt func(open, close int)
	bt = func(open, close int){
		if open == n && close == n{
			res = append(res, string(cur))
			return
		}
		if open < n{
			cur = append(cur, '(')
			bt(open+1, close)
			cur = cur[:len(cur)-1]
		}
		if close < open{
			cur = append(cur, ')')
			bt(open, close+1)
			cur = cur[:len(cur)-1]
		}
	}
	bt(0, 0)
	return res
}
