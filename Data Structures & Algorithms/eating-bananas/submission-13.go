func minEatingSpeed(piles []int, h int) int {
	maxPile := 0
	for _, v := range piles{
		maxPile = max(v, maxPile)
	}

	var helper func(speed int) bool
	helper = func(speed int) bool{
		time := 0
		for _, p := range piles{
			time += (p + speed - 1) / speed
		}
		return time <= h
	}

	l , r := 1, maxPile
	for l < r{
		m := l + (r - l) / 2
		if helper(m){
			r = m 
		} else {
			l = m + 1
		}
	}
	return r
}
