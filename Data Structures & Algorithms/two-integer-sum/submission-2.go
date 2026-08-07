func twoSum(nums []int, target int) []int {
    // return index pair (0-index)
    // Return the answer with the smaller index first
    numIdx := make(map[int]int) // value, index
    for i, num := range nums{
        if idx, exist := numIdx[target-num]; exist{
            return []int{idx, i}
        }
        numIdx[num] = i
    }
    return []int{}
}
