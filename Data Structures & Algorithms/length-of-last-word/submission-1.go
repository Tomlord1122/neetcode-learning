func lengthOfLastWord(s string) int {
	// trim the empty space at the end
	i := len(s)-1
	for i >= 0{
		if s[i] != ' '{
			break
		}
		i--
	}
	s = s[:i+1]
	res := 0
	for i := len(s)-1; i >= 0; i--{
		if s[i] == ' '{
			return res
		}
		res++
	}
	return res
}
