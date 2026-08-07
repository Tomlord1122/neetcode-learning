func maxArea(heights []int) int {

    l, r := 0, len(heights)-1
    area := 0
    for l < r{
        if heights[l] <= heights[r]{
            minHeight := heights[l]
            area = max(area, minHeight * (r - l))
            l++
        } else {
            minHeight := heights[r]
            area = max(area, minHeight * (r - l))
            r--
        }
    }
    return area
}
