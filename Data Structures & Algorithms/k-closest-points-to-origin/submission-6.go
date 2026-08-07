func kClosest(points [][]int, k int) [][]int {
	pairs := make([]pair, 0, len(points))
	for _, p := range points{
		x, y := p[0], p[1]
		d := math.Sqrt(float64(x*x+y*y))
		pairs = append(pairs, pair{coord:[]int{x, y}, distance:d})
	}


	// O(nlogn)
	sort.Slice(pairs, func(i, j int) bool{
		return pairs[i].distance < pairs[j].distance
	})

	res := [][]int{}
	for i := 0; i < k; i++{
		res = append(res, pairs[i].coord)
	}
	return res
}

type pair struct{
	coord []int
	distance float64
}
