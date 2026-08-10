func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool{
		return intervals[i][0] < intervals[j][0]
	})
	res := 0
	prevEnd := intervals[0][1]
	for _, val := range intervals[1:]{
		start, end := val[0], val[1]
		if start >= prevEnd{
			prevEnd = end
		} else {
			res++
			prevEnd = min(prevEnd, end)
		}
	}
	return res
}


