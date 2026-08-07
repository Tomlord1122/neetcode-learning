func maxArea(heights []int) int {
	l, r := 0, len(heights)-1
	area := 0
	for l < r{
		lHeight, rHeight := heights[l], heights[r]
		area = max(area, (r-l) * min(lHeight, rHeight))
		if lHeight < rHeight{
			l++
		} else {
			r--
		}
	} 
	return area
}
