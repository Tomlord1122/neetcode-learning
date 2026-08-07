func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool{
		return intervals[i][0] < intervals[j][0]
	})

	stk := [][]int{}

	for _, interval := range intervals{
		if len(stk) != 0 && stk[len(stk)-1][1] >= interval[0]{
			cur := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			stk = append(stk, []int{cur[0], max(cur[1],interval[1])})
		} else {
			stk = append(stk, []int{interval[0], interval[1]})
		}
	}
	return stk
}
