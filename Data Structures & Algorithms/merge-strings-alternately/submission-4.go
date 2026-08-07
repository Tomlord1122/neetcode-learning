func mergeAlternately(word1 string, word2 string) string {
	res := []byte{}
	str1 := []byte(word1)
	str2 := []byte(word2)
	i, j := 0, 0
	for i < len(str1) && j < len(str2){
		res = append(res, str1[i])
		i++
		res = append(res, str2[j])
		j++
	}
	if i != len(str1){
		res = append(res, str1[i:]...)
	} 
	if j != len(str2){
		res = append(res, str2[j:]...)
	}
	return string(res)
}
