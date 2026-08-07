func isAnagram(s string, t string) bool {
	var charCount [26]int
	if len(s) != len(t){
		return false
	}
	for i := 0; i < len(s); i++{
		charCount[s[i]-'a']++
		charCount[t[i]-'a']--
	}

	for i := 0; i < 26; i++{
		if charCount[i] != 0{
			return false
		}
	}
	return true
}
