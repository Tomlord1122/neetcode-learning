func carFleet(target int, position []int, speed []int) int {
    n := len(position)
    cars := make([]car, n)
    for i := 0; i < n; i++{
        cars[i].position = position[i]
        cars[i].speed = speed[i]
    }

    // sort by position in descending order
    sort.Slice(cars, func(i, j int) bool{
        return cars[i].position > cars[j].position
    })

    stack := []float64{}
    for i := 0; i < n; i++{
        time := float64(target - cars[i].position) / float64(cars[i].speed)
        if len(stack) == 0{
            stack = append(stack, time)
        } else if stack[len(stack)-1] >= time{
            continue
        } else {
            stack = append(stack, time)
        }
    }
    return len(stack)
}

type car struct{
    position int
    speed int
}

