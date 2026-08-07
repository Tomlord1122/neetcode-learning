func stringMatching(words []string) []string {
    res := []string{}
	for i := 0; i < len(words); i++{
		for j := 0; j < len(words); j++{
			// check if words[i] is substring of another string
			if i == j{
				continue
			}
			if strings.Contains(words[j], words[i]){
				res = append(res, words[i])
				break
			}
		}
	}
	return res
}