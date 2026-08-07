func characterReplacement(s string, k int) int {
	freqMap := make(map[byte]int)
	l := 0
	res := 0 
	maxFreq := 0
	for r := 0; r < len(s); r++{
		freqMap[s[r]]++
		maxFreq = max(maxFreq, freqMap[s[r]])
		for r - l + 1 - maxFreq > k{
			freqMap[s[l]]--
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}