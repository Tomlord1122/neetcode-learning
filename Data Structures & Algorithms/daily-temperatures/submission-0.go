func dailyTemperatures(temperatures []int) []int {
    idxStack := []int{}
    res := make([]int, len(temperatures))

    for i := 0; i < len(temperatures); i++{
        for len(idxStack) != 0 && temperatures[idxStack[len(idxStack)-1]] < temperatures[i]{
            // calculate the difference between indices.
            topIdx := idxStack[len(idxStack)-1]
            // pop the idxStack
            idxStack = idxStack[:len(idxStack)-1]
            res[topIdx] = i - topIdx
        }
        idxStack = append(idxStack, i)
    }
    return res
}
