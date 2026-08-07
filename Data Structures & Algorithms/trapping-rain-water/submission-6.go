func trap(height []int) int {
    n := len(height)
    prefixH := make([]int, n)
    postfixH := make([]int, n)

    prefixH[0] = 0
    postfixH[n-1] = 0
    for i := 1; i < len(height); i++{
        prefixH[i] = max(prefixH[i-1], height[i-1])
    }
    for i := n - 2; i >= 0; i--{
        postfixH[i] = max(postfixH[i+1], height[i+1])
    }
    area := 0
    for i, val := range height{
        minH := min(prefixH[i], postfixH[i])
        if minH > val{
            area += minH - val
        }
    }
    return area
}
