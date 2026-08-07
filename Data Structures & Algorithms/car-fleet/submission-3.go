func carFleet(target int, position []int, speed []int) int {
    cars := []car{}
    for i := 0; i < len(position); i++{
        cars = append(cars, car{position:position[i], speed:speed[i]})
    }
    sort.Slice(cars, func(i, j int) bool{
        return cars[i].position > cars[j].position
    })

    stack := []float64{}
    for i := 0; i < len(cars); i++{
        time := float64(target - cars[i].position) / float64(cars[i].speed)
        if len(stack) != 0 && time <= stack[len(stack)-1]{
            continue
        }
        stack = append(stack, time)
    }
    return len(stack)
}

type car struct{
    position int
    speed int
}