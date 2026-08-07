func mySqrt(x int) int {
	res := 0
	for i := 1; i <= x; i++{
		if i * i <= x{
			res = i
		} else {
			break
		}
	}
	return res
}
