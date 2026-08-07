func isAnagram(s string, t string) bool {
	var wordCount [26]int
	if len(s) != len(t){
		return false
	}

	for i := 0; i < len(s); i++{
		wordCount[s[i]-'a']++
	}
	for i := 0; i < len(t); i++{
		wordCount[t[i]-'a']--
		if wordCount[t[i]-'a'] < 0{
			return false
		}
	}

	for i := 0; i < 26; i++{
		if wordCount[i] != 0{
			return false
		}
	}
	return true
}

