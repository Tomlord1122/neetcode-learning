func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	stack := []int{} // store the index
	for i, v := range temperatures{
		for len(stack) != 0 && temperatures[stack[len(stack)-1]] < v{
			top := len(stack)-1
			res[stack[top]] = i - stack[top]
			stack = stack[:top]
		}
		stack = append(stack, i)
	}
	return res
}

// maintain a monotonic decreasing stack but store the index 
// rather than the value
