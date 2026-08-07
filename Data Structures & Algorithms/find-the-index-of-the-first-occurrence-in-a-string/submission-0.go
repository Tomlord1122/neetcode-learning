func strStr(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	for i := 0; i <= n-m; i++{
		j := 0
		for j < m{
			if haystack[i+j] != needle[j]{
				break
			}
			j++
		}
		if j == m{
			return i
		}
	}
	return -1
}
