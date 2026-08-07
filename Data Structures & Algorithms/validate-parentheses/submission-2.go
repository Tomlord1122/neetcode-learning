func isValid(s string) bool {
    // When it is a open brackets, we should push it into stack
    stack := []byte{}
    closeToOpen := map[byte]byte{
        '}':'{',
        ']':'[',
        ')':'(',
    }

    for i := 0; i < len(s); i++{
        if s[i] != '}' && s[i] != ']' && s[i] != ')'{
            stack = append(stack, s[i])
        } else if len(stack) != 0{
            top := stack[len(stack)-1]
            if top != closeToOpen[s[i]]{
                return false
            }
            stack = stack[:len(stack)-1]
        } else {
            return false
        }
    }
    return len(stack) == 0
}
