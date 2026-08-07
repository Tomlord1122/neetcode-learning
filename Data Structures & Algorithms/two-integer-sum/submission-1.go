func twoSum(nums []int, target int) []int {
    numIdx := make(map[int]int) // value, index
    for idx2, num := range nums{
        if idx, exist := numIdx[target-num]; exist{
            return []int{idx, idx2}
        }
        numIdx[num] = idx2
    }
    return []int{}
}
