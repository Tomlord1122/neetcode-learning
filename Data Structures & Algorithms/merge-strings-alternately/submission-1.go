func mergeAlternately(word1 string, word2 string) string {
	res := []byte{}
	i, j := 0, 0

	for i < len(word1) && j < len(word2){
		res = append(res, word1[i])
		i++
		res = append(res, word2[j])
		j++
	}

	if i < len(word1){
		res = append(res, word1[i:]...)
	} else {
		res = append(res, word2[j:]...)
	}
	return string(res)
}