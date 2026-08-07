func dailyTemperatures(temperatures []int) []int {
	stk := []int{} // we push index in this rather than value
	res := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++{
		for len(stk) != 0 && temperatures[i] > temperatures[stk[len(stk)-1]]{
			res[stk[len(stk)-1]] = i - stk[len(stk)-1]
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, i)
	}
	return res
}
