func generateParenthesis(n int) []string {
    res := []string{}
    stack := []byte{}

    var backtracking func(open, close int)
    backtracking = func(open, close int){
        if open == n && close == n{
            res = append(res, string(stack))
            return
        }
        if open < n{
            stack = append(stack, '(')
            backtracking(open+1, close)
            stack = stack[:len(stack)-1]
        }
        if close < open{
            stack = append(stack, ')')
            backtracking(open, close+1)
            stack = stack[:len(stack)-1]
        }
    }
    backtracking(0, 0)
    return res
}
