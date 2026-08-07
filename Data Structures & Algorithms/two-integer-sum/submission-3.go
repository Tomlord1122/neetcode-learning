func twoSum(nums []int, target int) []int {
    numIdx := make(map[int]int) // num, idx
    for i2, v := range nums{
        if i1, exist := numIdx[target-v]; exist{
            return []int{i1, i2}
        }
        numIdx[v] = i2
    }
    return []int{}
}
