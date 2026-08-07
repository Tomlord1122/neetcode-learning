func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool{
		if intervals[i][1] == intervals[j][1]{
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})

	stk := [][]int{}
	count := 0
	for i := 0; i < len(intervals); i++{
		if len(stk) != 0 && stk[len(stk)-1][1] > intervals[i][0]{
			count++
		} else {
			stk = append(stk, intervals[i])
		}
	}
	return count
}