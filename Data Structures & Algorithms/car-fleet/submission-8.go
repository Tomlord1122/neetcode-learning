func carFleet(target int, position []int, speed []int) int {
	pairs := []pair{}
	for i := 0; i < len(position); i++{
		pairs = append(pairs, pair{position:position[i], speed:speed[i]})
	}

	sort.Slice(pairs, func(i, j int) bool{
		return pairs[i].position > pairs[j].position
	})

	stk := []float64{}
	for _, pair := range pairs{
		time := float64(target - pair.position) / float64(pair.speed)
		if len(stk) > 0 && time <= stk[len(stk)-1]{
			continue
		} else {
			stk = append(stk, time)
		}
	}
	return len(stk)
}


type pair struct{
	position int
	speed int
}