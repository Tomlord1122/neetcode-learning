func isValid(s string) bool {
    stk := []byte{}
	closeToOpen := map[byte]byte{
		']': '[',
		'}': '{',
		')':'(',
	}

	for i := 0; i < len(s); i++{
		if s[i] == '(' || s[i] == '[' || s[i] == '{'{
			stk = append(stk, s[i])
		} else {
			if len(stk) == 0 || closeToOpen[s[i]] != stk[len(stk)-1]{
				return false
			}
			// pop the stk
			stk = stk[:len(stk)-1]
		}
	}
	return len(stk) == 0
}
