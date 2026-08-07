func findBuildings(heights []int) []int {
	res := []int{}

	for idx, height := range heights{
		for len(res) != 0 && heights[res[len(res)-1]] <= height{
			res = res[:len(res)-1]
		}
		res = append(res, idx)
	}
	return res
}

// push index to stack
// stk top value > current value
// maintain a monotonic increasing stack


// the return value shold be sorted in increasing order.

