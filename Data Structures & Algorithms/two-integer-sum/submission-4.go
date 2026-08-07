func twoSum(nums []int, target int) []int {
    numIdx := make(map[int]int) // value index
    for i, v := range nums{
        if id1, exist := numIdx[target-v]; exist{
            return []int{id1, i}
        }
        numIdx[v] = i
    }
    return []int{}
}
