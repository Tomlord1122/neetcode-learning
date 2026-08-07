func maxArea(heights []int) int {
	res := 0
	l, r := 0, len(heights)-1

	for l < r{
		if heights[l] < heights[r]{
			res = max(res, (r-l) * heights[l])
			l++
		} else {
			res = max(res, (r-l) * heights[r])
			r--
		}
	}
	return res
}
