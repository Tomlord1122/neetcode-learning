func maxArea(heights []int) int {
    res := 0
    left, right := 0, len(heights)-1
    for left < right{
        leftH, rightH := heights[left], heights[right]
        minHeight := min(leftH, rightH)
        res = max(res, (right-left) * minHeight)
        if leftH < rightH{
            left++
        } else {
            right--
        }
    }
    return res
}
