func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1
	for l < r{
		for l < r && !isValid(s[l]){
			l++
		}
		for l < r && !isValid(s[r]){
			r--
		}
		if s[l] != s[r]{
			return false
		}
		l++
		r--
	}
	return true
}

func isValid(c byte) bool{
	if c >= '0' && c <= '9' || c >= 'a' && c <= 'z'{
		return true
	}
	return false
}