func minEatingSpeed(piles []int, h int) int {
	maxPile := 0
	for _, val := range piles{
		maxPile = max(maxPile, val)
	}
	l, r := 1, maxPile
	for l < r{
		m := l + (r - l) / 2
		if canEat(piles, h, m){
		r = m 
		} else {
			l = m + 1
		}
	}
	return r
}


func canEat(piles []int, h int, speed int) bool{
	time := 0
	for _, pile := range piles{
		time += (pile + speed - 1) / speed
	}
	return time <= h
}