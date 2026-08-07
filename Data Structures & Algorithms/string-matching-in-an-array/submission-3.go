func stringMatching(words []string) []string {
    res := []string{}
	for i := 0; i < len(words); i++{
		for j := 0; j < len(words); j++{
			if i == j{
				continue
			}
			// check if words[i] is substring of words[j]
			if strings.Contains(words[j], words[i]){
				res = append(res, words[i])
				break
			}
		}
	}
	return res
}