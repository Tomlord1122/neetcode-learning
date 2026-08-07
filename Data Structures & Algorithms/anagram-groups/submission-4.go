func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[[26]int][]string)
	for _, s := range strs{
		wordCount := [26]int{}
		for i := 0; i < len(s); i++{
			wordCount[s[i]-'a']++
		}
		anagramMap[wordCount] = append(anagramMap[wordCount], s)
	}

	res := [][]string{}
	for _, group := range anagramMap{
		res = append(res, group)
	}
	return res
}
