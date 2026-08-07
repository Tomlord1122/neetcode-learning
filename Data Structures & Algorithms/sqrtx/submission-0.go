func mySqrt(x int) int {
	if x == 0{
		return 0
	}

	res := 1

	for i := 1; i <= x; i++{
		if i*i > x{
			return res
		}
		res = i 
	}

	return res
}
