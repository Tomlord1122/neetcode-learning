func minEatingSpeed(piles []int, h int) int {
	maxPile := 0
	for _, v := range piles{
		maxPile = max(maxPile, v)
	}
	l, r := 1, maxPile
	res := math.MaxInt
	for l <= r{
		m := (l+r) / 2
		if canEat(piles, h, m){
			r = m - 1
			res = min(res, m)
		} else {
			l = m + 1
		}
	}
	return res
}


func canEat (piles []int, h int, m int) bool{
	time := 0
	for _, pile := range piles{
		time += (pile + m - 1) / m
	}
	return time <= h
}