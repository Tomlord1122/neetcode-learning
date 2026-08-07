func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    left, right := 0, len(s)-1
    for left < right{
        for left < right && !isAlphanumeric(s[left]){
            left++
        }
        for left < right && !isAlphanumeric(s[right]){
            right--
        }
        if s[left] != s[right]{
            return false
        }
        left++
        right--
    }
    return true
}



func isAlphanumeric(char byte) bool{
    if char >= 'a' && char  <= 'z' || 
    char >= 'A' && char <= 'Z' || 
    char >= '0' && char <= '9'{
        return true
    }
    return false
}