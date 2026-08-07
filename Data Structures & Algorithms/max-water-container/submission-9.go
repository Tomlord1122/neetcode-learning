func maxArea(heights []int) int {
	l, r := 0, len(heights)-1
	res := 0
	for l < r{
		minH := min(heights[l], heights[r])
		res = max(res, minH * (r-l))
		if minH == heights[l]{
			l++ 
		} else {
			r--
		}
	}
	return res
}
