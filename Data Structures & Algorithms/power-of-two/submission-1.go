func isPowerOfTwo(n int) bool {
	if n <= 0{
		return false
	}
	x := 1
	for x < n{
		x *= 2
	}
	return x == n
}
