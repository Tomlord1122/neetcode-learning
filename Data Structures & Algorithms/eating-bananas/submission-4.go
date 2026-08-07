func minEatingSpeed(piles []int, h int) int {
    maxPile := 0
    for _, pile := range piles{
        maxPile = max(maxPile, pile)
    }

    l, r := 1, maxPile
    for l <= r{
        m := (l+r)/2
        if canEat(piles, m, h){
            r = m - 1
        } else {
            l = m + 1
        }
    }
    return l
}

func canEat(piles []int, m int, h int) bool{
    time := 0
    for _, pile := range piles{
        time += (pile + m - 1) / m
    }
    return time <= h
}