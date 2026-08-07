func longestCommonPrefix(strs []string) string {
	var res string
	for i := 0; i < len(strs[0]); i++{
		for _, s := range strs{
			// s may be out of bound
			if i == len(s) || s[i] != strs[0][i]{
				return res
			}
		}
		res += string(strs[0][i])
	}
	return strs[0]
}