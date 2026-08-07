func trap(height []int) int {
    prefixHeight := make([]int, len(height))
    postfixHeight := make([]int, len(height))

    prefixHeight[0] = 0
    postfixHeight[len(height)-1] = 0

    for i := 1; i < len(height); i++{
        prefixHeight[i] = max(prefixHeight[i-1], height[i-1])
    }
    for i := len(height)-2; i >= 0; i--{
        postfixHeight[i] = max(postfixHeight[i+1], height[i+1])
    }
    area := 0

    for i := 0; i < len(height); i++{
        minHeight := min(prefixHeight[i], postfixHeight[i])
        if height[i] < minHeight{
            area += minHeight - height[i]
        }
    }
    return area
}
