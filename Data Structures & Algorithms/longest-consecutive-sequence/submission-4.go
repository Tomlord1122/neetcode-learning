func longestConsecutive(nums []int) int {
    seen := make(map[int]bool)
    res := 0

    for _, num := range nums{
        seen[num] = true
    }

    for _, num := range nums{
        if seen[num-1]{
            continue
        }

        length := 1
        cur := num
        for seen[cur+1]{
            cur++
            length++
        }
        res = max(res, length)
    }
    return res
}
