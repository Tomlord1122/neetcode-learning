func isValid(s string) bool {
    closeToOpen := map[byte]byte{
		')':'(',
		']':'[',
		'}':'{',
	}

	stack := []byte{}
	for i := 0; i < len(s); i++{
		if s[i] != ')' && s[i] != ']' && s[i] != '}'{
			stack = append(stack, s[i])
		} else {
			
			if len(stack) == 0 || stack[len(stack)-1] != closeToOpen[s[i]]{
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}
