func topKFrequent(nums []int, k int) []int {
	numFreq := make(map[int]int) // val -> fre
	for _, val := range nums{
		numFreq[val]++
	}
	pairs := []pair{}
	for val, freq := range numFreq{
		pairs = append(pairs, pair{val:val, freq:freq})
	}

	sort.Slice(pairs, func(i, j int) bool{
		return pairs[i].freq > pairs[j].freq
	})

	res := []int{}
	for i := 0; i < k && i < len(pairs); i++{
		res = append(res, pairs[i].val)
	}
	return res
}

type pair struct{
	val int
	freq int
}



// return the k most frequent elements within the array
// maintain an array that sort by frequency
