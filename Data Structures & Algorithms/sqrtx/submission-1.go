func mySqrt(x int) int {
	l, r := 1, x
	res := 0

	for l <= r{
		m := l + (r-l) / 2
		if m * m > x{
			r = m - 1
		} else {
			l = m + 1
			res = m 
		}
	}
	return res
}
