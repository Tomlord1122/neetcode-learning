func longestConsecutive(nums []int) int {
    seen := make(map[int]bool)
    res := 0
    for _, num := range nums{
        seen[num] = true
    }
    
    for _, num := range nums{
        if _, exist := seen[num-1]; exist{
            continue
        }
        length := 1
        nxt := num + 1
        for seen[nxt]{
            length++
            nxt++
        }
        res = max(res, length)
    }
    return res
}


// Find the start index with a hash set