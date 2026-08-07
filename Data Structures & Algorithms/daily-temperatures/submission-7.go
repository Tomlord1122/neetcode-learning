func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	stk := []int{}
	res := make([]int, n)
	for idx, val := range temperatures{
		for len(stk) != 0 && temperatures[stk[len(stk)-1]] < val{
			res[stk[len(stk)-1]] = idx - stk[len(stk)-1]
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, idx)
	}
	return res
}


// monotonic decreasing stack
// when the current temp is greater than the top stack
// we should keep poping stack and record the waiting day