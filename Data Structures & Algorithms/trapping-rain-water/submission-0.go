func trap(height []int) int {
    n := len(height)
    leftMax, rightMax := make([]int, n), make([]int, n)

    for i := 1; i < n; i++{
        leftMax[i] = max(leftMax[i-1], height[i-1])
    }
    for i := n - 2; i >= 0; i--{
        rightMax[i] = max(rightMax[i+1], height[i+1])
    }

    area := 0
    for i := 0; i < n; i++{
        minHeight := min(leftMax[i], rightMax[i])
        if height[i] < minHeight{
            area += minHeight - height[i]
        }
    }
    return area
}
