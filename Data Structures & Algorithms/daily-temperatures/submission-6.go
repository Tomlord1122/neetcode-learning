func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	stk := []int{} // store the index of each temperatures
	for i, t := range temperatures{
		for len(stk) != 0 && t > temperatures[stk[len(stk)-1]]{
			res[stk[len(stk)-1]] = i - stk[len(stk)-1]
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, i)
	}
	return res
}
