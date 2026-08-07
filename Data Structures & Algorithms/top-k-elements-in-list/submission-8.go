func topKFrequent(nums []int, k int) []int {
	numFreq := make(map[int]int) // num -> freq

	for _, num := range nums{
		numFreq[num]++
	}

	pairs := [][]int{}
	for num, freq := range numFreq{
		pairs = append(pairs, []int{num, freq})
	}

	// sory by freq
	sort.Slice(pairs, func(i, j int) bool{
		return pairs[i][1] > pairs[j][1]
	})

	res := make([]int, k)
	for i := 0; i < k; i++{
		res[i] = pairs[i][0]
	}
	return res
}
