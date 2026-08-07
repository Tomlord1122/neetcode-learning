func longestConsecutive(nums []int) int {
    exist := make(map[int]bool)
    for _, num := range nums{
        exist[num] = true
    }
    res := 0
    for _, num := range nums{
        if !exist[num-1]{
            count := 1
            tmp := num
            for exist[tmp+1]{
                count++
                tmp++
            }
            res = max(res, count)
        }
    }
    return res
}


// find the start index