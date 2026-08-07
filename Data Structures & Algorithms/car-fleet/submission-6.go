func carFleet(target int, position []int, speed []int) int {
	pairs := make([]pair, len(position))
	for i := 0; i < len(position); i++{
		pairs[i] = pair{position:position[i], speed:speed[i]}
	}
	// sort by position in descending order
	sort.Slice(pairs, func(i, j int) bool{
		return pairs[i].position > pairs[j].position
	})

	stk := []float64{}
	for _, car := range pairs{
		time := float64(target - car.position) / float64(car.speed)
		if len(stk) != 0 && stk[len(stk)-1] >= time{
			continue
		}
		stk = append(stk, time)
	}
	return len(stk)
}

type pair struct{
	position int
	speed int
}