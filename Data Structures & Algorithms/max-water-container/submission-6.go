func maxArea(heights []int) int {
	n := len(heights)
	l, r := 0, n-1
	res := 0
	for l < r{
		if heights[l] < heights[r]{
			res = max(res, (r-l)*heights[l])
			l++
		} else{
			res = max(res, (r-l)*heights[r])
			r--
		}
	}
	return res
}
