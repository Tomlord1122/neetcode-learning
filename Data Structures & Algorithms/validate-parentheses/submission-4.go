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
            // Check the parentheses condition and pop the stack
            if len(stack) != 0 && closeToOpen[s[i]] == stack[len(stack)-1]{
                stack = stack[:len(stack)-1]
            } else {
                return false
            }
        }
    }
    return len(stack) == 0
}
