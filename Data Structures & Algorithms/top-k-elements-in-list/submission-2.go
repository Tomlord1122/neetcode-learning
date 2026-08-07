type pair struct{
    val int
    freq int
}

func topKFrequent(nums []int, k int) []int {
    numFreq := make(map[int]int) // val, freq
    for _, num := range nums{
        numFreq[num]++
    }

    var numPair []pair
    for val, freq := range numFreq{
        numPair = append(numPair, pair{val:val, freq:freq})
    }

    sort.Slice(numPair, func(i, j int) bool{
        return numPair[i].freq > numPair[j].freq
    })

    res := []int{}
    for i := 0; i < k; i++{
        res = append(res, numPair[i].val)
    }
    return res
}

