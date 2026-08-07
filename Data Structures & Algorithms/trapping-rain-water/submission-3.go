func trap(height []int) int {
    area := 0
    n := len(height)
    prefixH := make([]int, n)
    postfixH := make([]int, n)
    prefixH[0] = 0
    postfixH[n-1] = 0

    for i := 1; i < n; i++{
        prefixH[i] = max(prefixH[i-1], height[i-1]) 
    }
    for i := n-2; i >= 0; i--{
        postfixH[i] = max(postfixH[i+1], height[i+1])
    }

    // Get the area
    for i := 0; i < len(height); i++{
        minHeight := min(prefixH[i], postfixH[i])
        if minHeight > height[i]{
            area += (minHeight - height[i])
        }
    }
    return area
}
