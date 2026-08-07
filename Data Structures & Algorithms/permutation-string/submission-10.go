func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2){
		return false
	}
	s1Map := [26]int{}
	s2Map := [26]int{}
	for i := 0; i < len(s1); i++{
		s1Map[s1[i]-'a']++
		s2Map[s2[i]-'a']++
	}
	// Calculate the current matches
	matches := 0
	for i := 0; i < 26; i++{
		if s1Map[i] == s2Map[i]{
			matches++
		}
	}
	if matches == 26{
		return true
	}
	start := 0
	for i := len(s1); i < len(s2); i++{
		// remove the first letter and add the new letter 
		if matches == 26{
			return true
		}
		leftVal := s2[start] - 'a'
		if s1Map[leftVal] == s2Map[leftVal]{
			matches--
		}
		s2Map[leftVal]--
		if s1Map[leftVal] == s2Map[leftVal]{
			matches++
		}
		start++
		// add the new letter
		rightVal := s2[i] - 'a'
		if s1Map[rightVal] == s2Map[rightVal]{
			matches--
		}
		s2Map[rightVal]++
		if s1Map[rightVal] == s2Map[rightVal]{
			matches++
		}
	}
	return matches == 26
}
