func prefixCount(words []string, pref string) int {
	n := len(pref)
	res := 0
	for _, w := range words{
		if len(w) >= n && w[:n] == pref{
			res += 1
		}
	}
	return res
}
