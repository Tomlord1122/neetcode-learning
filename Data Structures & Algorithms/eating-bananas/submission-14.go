func minEatingSpeed(piles []int, h int) int {
	maxPile := 0
	for _, val := range piles{
		maxPile = max(maxPile, val)
	}
	l, r := 1, maxPile
	var canEat func(speed int) bool
	canEat = func(speed int) bool{
		time := 0
		for _, pile := range piles{
			time += (pile + speed - 1) / speed
		}
		return time <= h
	}

	for l < r{
		speed := l + (r - l) / 2
		if canEat(speed){
			r = speed
		} else {
			l = speed + 1
		}
	}
	return r
}
