func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	cars := make([]car, n)
	for i := 0; i < n; i++{
		cars[i].position = position[i]
		cars[i].speed = speed[i]
	}

	sort.Slice(cars, func(i, j int) bool{
		return cars[i].position > cars[j].position
	})

	stk := []float64{}
	for i := 0; i < n; i++{
		time := float64(target - cars[i].position) / float64(cars[i].speed)
		if len(stk) != 0 && time <= stk[len(stk)-1]{
			continue
		}
		stk = append(stk, time)
	}
	return len(stk)
}

type car struct{
	position int
	speed int
}