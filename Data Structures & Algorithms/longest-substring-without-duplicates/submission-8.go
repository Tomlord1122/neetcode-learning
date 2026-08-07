func lengthOfLongestSubstring(s string) int {
	longest, l := 0, 0
	charMap := make(map[byte]bool)
	for r := 0; r < len(s); r++{
		for charMap[s[r]]{
			// remove left most char
			charMap[s[l]] = false
			l++	
		} 
		charMap[s[r]] = true
		longest = max(longest, r-l+1)
	}
	return longest
}
