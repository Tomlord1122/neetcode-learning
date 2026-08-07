func merge(intervals [][]int) [][]int {
    res := [][]int{}
	sort.Slice(intervals, func(i, j int) bool{
		return intervals[i][0] < intervals[j][0]
	})
	for i := 0; i < len(intervals); i++{
		if len(res) != 0 && intervals[i][0] <= res[len(res)-1][1]{
			end := max(intervals[i][1], res[len(res)-1][1])
			res[len(res)-1][1] = end
		} else {
			res = append(res, []int{intervals[i][0], intervals[i][1]})
		}
	}
	return res
}
