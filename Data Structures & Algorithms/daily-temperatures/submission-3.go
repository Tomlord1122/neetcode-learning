func dailyTemperatures(temperatures []int) []int {
    res := make([]int, len(temperatures))
    stack := []int{}

    for idx, temp := range temperatures{
        for len(stack) != 0 && temperatures[stack[len(stack)-1]] < temp{
            res[stack[len(stack)-1]] = idx - stack[len(stack)-1]
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, idx)
    }
    return res
}

// monotonic decreasing stack
