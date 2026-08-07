func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	anagram := make(map[[26]int][]string)

	for _, str := range strs{
		charCount := [26]int{}
		for i := 0; i < len(str); i++{
			charCount[str[i]-'a']++
		}
		anagram[charCount] = append(anagram[charCount], str)
	}
	
	for _, group := range anagram{
		res = append(res, group)
	}
	return res
}
