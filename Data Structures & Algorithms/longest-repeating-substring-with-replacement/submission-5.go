func characterReplacement(s string, k int) int {
	count := make(map[byte]int)

	res, l, maxf := 0, 0, 0
	for r := 0; r < len(s); r++{
		count[s[r]]++
		maxf = max(maxf, count[s[r]])
		for r - l + 1 - maxf > k{
			count[s[l]]--
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}
