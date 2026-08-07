func findTheDifference(s string, t string) string {
	var numCount [26]int
	for i := 0; i < len(s); i++{
		numCount[s[i]-'a']++
	}
	for i := 0; i < len(t); i++{
		numCount[t[i]-'a']--
		if numCount[t[i]-'a'] < 0{
			return string(t[i])
		}
	}
	return ""
}


