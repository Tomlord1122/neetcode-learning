func singleNumber(nums []int) int {
    a := 0
    for _ , num := range nums{
        a = a ^ num 
    }
    return a 
}
