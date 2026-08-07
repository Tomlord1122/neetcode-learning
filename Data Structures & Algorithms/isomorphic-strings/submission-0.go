func isIsomorphic(s string, t string) bool {
	sToT := make(map[byte]byte)
	tToS := make(map[byte]byte)
	if len(s) != len(t){
		return false
	}
	for i := 0; i < len(s); i++{
		sChar := s[i]
		tChar := t[i]
		if val, ok := sToT[sChar]; ok {
			if val != tChar {
				return false
			}
		} else {
			sToT[sChar] = tChar
		}
		if val, ok := tToS[tChar]; ok {
			if val != sChar {
				return false
			}
		} else {
			tToS[tChar] = sChar
		}
	}
	return true
}