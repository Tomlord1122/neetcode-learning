func countBits(n int) []int {
	res := make([]int, n+1)
	res[0] = 0
	for i := 1; i < n+1; i++{
		res[i] = res[i>>1] + i & 1
	}
	return res
}


// 0 --> 0
// 1 --> 1
// 2 --> 10
// 3 --> 11
// 4 --> 100
// 5 -> 101