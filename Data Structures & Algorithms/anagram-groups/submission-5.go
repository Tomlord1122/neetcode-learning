func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[[26]int][]string)
	for _, s := range strs{
		cur := [26]int{}
		for i := 0; i < len(s); i++{
			cur[s[i]-'a']++
		}
		// add it to anagrams
		anagrams[cur] = append(anagrams[cur], s)
	}
	res := [][]string{}
	for _, group := range anagrams{
		res = append(res, group)
	}
	return res
}
