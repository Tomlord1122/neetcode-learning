func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int) // value -> index
    
    for idx1, val := range nums{
        if idx2, exist := numMap[target-val]; exist{
            return []int{idx2, idx1}
        }
        numMap[val] = idx1
    }
    return []int{}
}
