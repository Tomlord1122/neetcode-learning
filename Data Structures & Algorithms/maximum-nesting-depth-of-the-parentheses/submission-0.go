func maxDepth(s string) int {
	stk := []byte{}
	res := 0
	for i := 0; i < len(s); i++{
		if s[i] == '('{
			stk = append(stk, '(')
			res = max(res, len(stk))
		} else if s[i] == ')'{
			stk = stk[:len(stk)-1]
		}
	}
	return res
}



// a paraentheses question may be related to maintain a stack
// I think we can keep tracking the length of the stack
// It means the current nesting depty of the paraentheses
// So just use a res var to track the maximum nesting depth