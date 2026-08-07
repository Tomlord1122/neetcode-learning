func minEatingSpeed(piles []int, h int) int {
	maxPile := 0
	for _, v := range piles{
		maxPile = max(maxPile, v)
	}

	var canEat func(speed int) bool
	canEat = func(speed int) bool {
		time := 0
		for _, pile := range piles{
			time += (pile + speed - 1) / speed
		}
		return time <= h
	}

	l, r := 1, maxPile
	for l <= r{
		speed := (l+r) / 2
		if canEat(speed){
			r = speed -  1
		} else {
			l = speed + 1
		}
	}
	return l
}
