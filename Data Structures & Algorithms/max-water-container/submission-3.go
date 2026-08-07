func maxArea(heights []int) int {
    l, r := 0, len(heights)-1
    area := 0
    for l <= r{
       leftH, rightH := heights[l], heights[r]
       if leftH < rightH{
        area = max(area, (r - l) * leftH)
        l++
       } else {
        area = max(area, (r - l) * rightH)
        r--
       }
    }
    return area
}
