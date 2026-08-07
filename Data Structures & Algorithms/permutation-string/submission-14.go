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
		leftIdx := s2[l] - 'a'
		if s1Map[leftIdx] == s2Map[leftIdx]{
			matches--
		}
		s2Map[leftIdx]--
		l++
		if s1Map[leftIdx] == s2Map[leftIdx]{
			matches++
		}
		rightIdx := s2[r] - 'a'
		if s1Map[rightIdx] == s2Map[rightIdx]{
			matches--
		}
		s2Map[rightIdx]++
		if s1Map[rightIdx] == s2Map[rightIdx]{
			matches++
		}
	}
	return matches == 26
}
