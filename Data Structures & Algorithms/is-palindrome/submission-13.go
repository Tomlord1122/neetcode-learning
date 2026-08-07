func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j{
		for i < j && !isValid(s[i]){
			i++
		}
		for i < j && !isValid(s[j]){
			j--
		}
		if s[i] != s[j]{
			return false
		}
		i++
		j--
	}
	return true
}


func isValid(c byte) bool{
	if (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z'){
		return true
	}
	return false
}