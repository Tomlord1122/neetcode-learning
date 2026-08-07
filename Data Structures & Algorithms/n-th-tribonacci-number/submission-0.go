func tribonacci(n int) int {
	if n == 0{
		return 0
	}
	if n == 1 || n == 2{
		return 1
	}
	tArray := make([]int, n+1)
	tArray[0], tArray[1], tArray[2] = 0, 1, 1

	for i := 3; i <= n; i++{
		tArray[i] = tArray[i-3] + tArray[i-2] + tArray[i-1]
	}
	return tArray[n]
}
