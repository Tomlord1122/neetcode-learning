func topKFrequent(nums []int, k int) []int {
	// get the frequency.
	numFreq := make(map[int]int)
	for _, num := range nums{
		numFreq[num]++
	}
	// sort by frequency
	pairs := [][]int{}
	
	for val, freq := range numFreq{
		pairs = append(pairs, []int{freq, val})
	}

	sort.Slice(pairs, func(i, j int) bool{
		return pairs[i][0] > pairs[j][0]
	})
	// get the top k
	res := []int{}
	for i := 0; i < k; i++{
		res = append(res, pairs[i][1])
	}
	return res
}
