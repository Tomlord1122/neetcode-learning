func productExceptSelf(nums []int) []int {
    res := make([]int, len(nums))
    prefixProd := make([]int, len(nums))
    postfixProd := make([]int, len(nums))

    prefixProd[0] = 1
    postfixProd[len(nums)-1] = 1
    for i := 1; i < len(nums); i++{
        prefixProd[i] = prefixProd[i-1] * nums[i-1]
    }
    for i := len(nums)-2; i >= 0; i--{
        postfixProd[i] = postfixProd[i+1] * nums[i+1]
    }

    // Get the res by prefixProd and postfixProd
    for i := 0; i < len(nums); i++{
        res[i] = prefixProd[i] * postfixProd[i]
    }
    return res
}
