func dailyTemperatures(temperatures []int) []int {
    stack := []int{}
    res := make([]int, len(temperatures))
    for idx, val := range temperatures{
        for len(stack) != 0 && temperatures[stack[len(stack)-1]] < val{
            res[stack[len(stack)-1]] = idx - stack[len(stack)-1]
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, idx)
    }
    return res
}