func longestCommonPrefix(strs []string) string {
	for i := 0; i < len(strs[0]); i++{
		for _, s := range strs{
			if i == len(s) || s[i] != strs[0][i]{
				return s[:i]
			}
		}
	}
	return strs[0]
}


// we can choose strs[0] as longest at first
// then compare every str in strs character by character
