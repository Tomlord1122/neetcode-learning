func productExceptSelf(nums []int) []int {
    res := make([]int, len(nums))

    val := 1
    res[0] = 1
    for i := 1; i < len(nums); i++{
        res[i] = val * nums[i-1]
        val = res[i]
    }

    val = 1
    for i := len(nums)-2; i >= 0; i--{
        val = val * nums[i+1]
        res[i] = res[i] * val
    }
    return res
}
