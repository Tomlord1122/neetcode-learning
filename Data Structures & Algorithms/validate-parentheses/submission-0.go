func isValid(s string) bool {
    closeToOpen := map[byte]byte{
        ')':'(',
        ']':'[',
        '}':'{',
    }

    stack := []byte{}

    for i := 0; i < len(s); i++{
        if s[i] == '(' || s[i] == '[' || s[i] == '{'{
            stack = append(stack, s[i])
        } else {
            if len(stack) == 0{
                return false
            }
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if closeToOpen[s[i]] != top{
                return false
            }
        }
    }
    return len(stack) == 0
}



// When we get a open bracket -> Push it into a stack
// When we get a close bracket -> pop the stack and check if it is valid
// Maintain a hashMap for verification

// In what situation. We can know that s is valid?
// stack is empty

