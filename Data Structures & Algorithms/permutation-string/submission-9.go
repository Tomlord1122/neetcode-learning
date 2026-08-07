func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	s1Map := make([]int, 26)
	s2Map := make([]int, 26)
	for i := 0; i < len(s1); i++{
		s1Map[s1[i]-'a']++
		s2Map[s2[i]-'a']++
	}

	matches := 0 
	for i := 0; i < 26; i++{
		if s1Map[i] == s2Map[i]{
			matches++
		}
	}

	l := 0
	for r := len(s1); r < len(s2); r++{
		if matches == 26{
			return true
		}
		// update left
		lIdx := s2[l] - 'a'
		if s1Map[lIdx] == s2Map[lIdx]{
			matches--
		}
		s2Map[lIdx]--
		l++
		if s1Map[lIdx] == s2Map[lIdx]{
			matches++
		}
		// update right
		rIdx := s2[r] - 'a'
		if s1Map[rIdx] == s2Map[rIdx]{
			matches--
		}
		s2Map[rIdx]++
		if s1Map[rIdx] == s2Map[rIdx]{
			matches++
		}
	}
	return matches == 26
}
