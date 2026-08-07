func largestRectangleArea(heights []int) int {
    maxArea := 0
    stack := make([][2]int, len(heights)) // (index, height)

    for idx, height := range heights{
        start := idx
        for len(stack) != 0 && stack[len(stack)-1][1] > height{
            i := stack[len(stack)-1][0]
            h := stack[len(stack)-1][1]
            maxArea = max(maxArea, h * (idx - i))
            start = i
            stack = stack[:len(stack)-1]
        }
        // We need to update the index with start
        // This can make sure we can find the potential area
        stack = append(stack, [2]int{start, height})
    }

    n := len(heights)
    for i := 0; i < len(stack); i++{
        index := stack[i][0]
        height := stack[i][1]
        maxArea = max(maxArea, height * (n - index))
    }
    return maxArea
}
