func isValid(s string) bool {
	bracketMap := map[byte]byte{
		')':'(',
		']':'[',
		'}':'{',
	}
	stk := []byte{}
	for i := 0; i < len(s); i++{
		if s[i] == '(' || s[i] == '[' || s[i] == '{'{
			stk = append(stk, s[i])
		} else if len(stk) != 0 && stk[len(stk)-1] == bracketMap[s[i]]{
			stk = stk[:len(stk)-1]
		} else {
			return false
		}
	}
	return len(stk)==0
}
