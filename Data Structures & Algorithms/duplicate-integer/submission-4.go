func hasDuplicate(nums []int) bool {
    seen := make(map[int]bool)

    for _, num := range nums{
        if exist, _ := seen[num]; exist{
            return true
        }
        seen[num] = true
    }
    return false
}
