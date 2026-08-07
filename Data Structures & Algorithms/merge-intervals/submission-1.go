func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool{
		return intervals[i][0] < intervals[j][0]
	})
	stk := [][]int{}
	for i := 0; i < len(intervals); i++{
		if len(stk) != 0 && stk[len(stk)-1][1] >= intervals[i][0]{
			top := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			stk = append(stk, []int{top[0], max(top[1], intervals[i][1])})
		} else {
			stk = append(stk, []int{intervals[i][0], intervals[i][1]})
		}
	}
	return stk
}

// overlapping condition
// previous end time is >= current start time
// the new interval should be (previous start time, and min(pEndTime, cStartTime))