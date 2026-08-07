func topKFrequent(nums []int, k int) []int {
    numFreq := make(map[int]int)

    for _, num := range nums{
        numFreq[num]++
    }

    pairs := []pair{}
    for num, freq := range numFreq{
        pairs = append(pairs, pair{val:num, freq:freq})
    }

    sort.Slice(pairs, func(i, j int) bool{
        return pairs[i].freq > pairs[j].freq
    })

    res := []int{}

    for i := 0; i < k; i++{
        res = append(res, pairs[i].val)
    }
    return res
}


type pair struct{
    val int
    freq int
}