func carFleet(target int, position []int, speed []int) int {
    cars := make([]pair, len(position))

    for i := 0; i < len(position); i++{
        cars[i] = pair{position: position[i], speed: speed[i]}
    }

    sort.Slice(cars, func(i, j int) bool{
        return cars[i].position > cars[j].position
    })

    stack := []float64{}
    for i := 0; i < len(cars); i++{
        time := float64(target - cars[i].position) / float64(cars[i].speed)
        if len(stack) == 0 || time > stack[len(stack)-1]{
            stack = append(stack, time)
        }
    }
    return len(stack)
}

type pair struct{
    position int
    speed int
}



