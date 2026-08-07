func topKFrequent(nums []int, k int) []int {
    freq := make(map[int]int)
    for _, num := range nums{
        freq[num]++
    }

    buckets := make([][]int, len(nums)+1)
    for val, f := range freq{
        buckets[f] = append(buckets[f], val)
    }

    res := make([]int, 0, k)
    for f := len(buckets)-1; f >= 1 && len(res) < k; f--{
        for _, val := range buckets[f]{
            res = append(res, val)
            if len(res) == k{
                return res
            }
        }
    }
    return res
}
