func topKFrequent(nums []int, k int) []int {
    numCount := make(map[int]int) // val, freq

    for _, num := range nums{
        numCount[num]++
    }

    var numPair []pair
    for val, freq := range numCount{
        numPair = append(numPair, pair{val:val, freq:freq})
    }

    // Sort the numPair in descending order
    sort.Slice(numPair, func(i, j int) bool{
        return numPair[i].freq > numPair[j].freq
    })

    res := []int{}

    for i := 0; i < k && i < len(numPair); i++{
        res = append(res, numPair[i].val)
    }
    return res
}


type pair struct{
    val int
    freq int
}
// Make a new struct (val, freq)
// Sort this by freq (descending order)
// Append to res array
