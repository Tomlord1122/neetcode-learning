func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, val := range nums{
		freqMap[val]++
	}
	pairs := []pair{}
	for val, freq := range freqMap{
		pairs = append(pairs, pair{val:val, freq:freq})
	}

	sort.Slice(pairs, func(i, j int)bool{
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
