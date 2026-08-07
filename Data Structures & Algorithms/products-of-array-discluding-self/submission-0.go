func productExceptSelf(nums []int) []int {
    n := len(nums)
    prefixProduct := make([]int, n)
    postfixProduct := make([]int, n)
    prefixProduct[0], postfixProduct[n-1] = 1, 1

    for i := 1; i < n; i++{
        prefixProduct[i] = nums[i-1] * prefixProduct[i-1]
    }
    for i := n - 2; i >= 0; i--{
        postfixProduct[i] = nums[i+1] * postfixProduct[i+1]
    }

    res := make([]int, n)
    for i := 0; i < n; i++{
        res[i] = prefixProduct[i] * postfixProduct[i]
    }
    return res
}
