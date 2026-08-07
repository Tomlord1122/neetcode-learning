func lengthOfLastWord(s string) int {
	res := 0
	i := len(s)-1
	for i >= 0{
		if s[i] != ' '{
			break
		}
		i--
	}
	s = s[:i+1]
	for i := len(s)-1; i >= 0; i--{
		if s[i] == ' '{
			return res
		}
		res++
	}
	return res
}
