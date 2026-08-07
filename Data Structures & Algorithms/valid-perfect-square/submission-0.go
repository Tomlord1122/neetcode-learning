func isPerfectSquare(num int) bool {
	l, r := 1, num
	for l <= r{
		m := l + (r-l) / 2
		if m*m == num{
			return true
		} else if m*m < num{
			// move l
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return false
}
