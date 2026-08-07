func maxArea(heights []int) int {
    l, r := 0, len(heights)-1
    area := 0
    for l < r{
        if heights[l] <= heights[r]{
            area = max(area, (r-l) * heights[l])
            l++
        } else {
            area = max(area, (r-l) * heights[r])
            r--
        }
    }
    return area
}
