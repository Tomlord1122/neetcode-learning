func minEatingSpeed(piles []int, h int) int {
    maxPile := 0
    for _, val := range piles{
        maxPile = max(maxPile, val)
    }

    l, r := 1, maxPile
    for l < r{
        speed := (l + r) / 2
        if canEatAll(piles, speed, h){
            r = speed
        } else {
            l = speed + 1
        }
    }
    return r
}


func canEatAll(piles []int, speed int, h int) bool{
    time := 0
    for _, pile := range piles{
        time += (pile + speed - 1) / speed
    }
    return time <= h
}