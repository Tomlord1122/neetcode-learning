func lengthOfLongestSubstring(s string) int {
	l := 0
	longest := 0
	charMap := make(map[byte]bool)
	for r := 0; r < len(s); r++{
		for charMap[s[r]]{
			charMap[s[l]] = false
			l++
		}
		charMap[s[r]] = true
		longest = max(longest, r - l + 1)
	}
	return longest
}
