func missingNumber(nums []int) int {
    miss := 0
    for _, num := range nums{
        miss = miss ^ num
    }
    n := len(nums)
    for i := 0; i <= n; i++{
        miss = miss ^ i
    }
    return miss
}
