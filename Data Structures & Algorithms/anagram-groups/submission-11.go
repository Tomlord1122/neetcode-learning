func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)
	for _, s := range strs{
		var cur [26]int
		for i := 0; i < len(s); i++{
			cur[s[i]-'a']++
		}
		groups[cur] = append(groups[cur], s)
	}

	res := [][]string{}
	for _, group := range groups{
		res = append(res, group)
	}
	return res
}
