func isIsomorphic(s string, t string) bool {
	sToT := make(map[byte]byte)
	tToT := make(map[byte]byte)
	if len(s) != len(t){
		return false
	}

	for i := 0; i < len(s); i++{
		sChar := s[i]
		tChar := t[i]
		if val, exist := sToT[sChar]; exist{
			if val != tChar{
				return false
			}
		} else {
			// add this to current map
			sToT[sChar] = tChar
		}

		if val, exist := tToT[tChar]; exist{
			if val != sChar{
				return false
			}
		} else {
			tToT[tChar] = sChar
		}
	}

	return true
}
