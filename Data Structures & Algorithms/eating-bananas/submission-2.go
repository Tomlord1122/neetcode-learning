func minEatingSpeed(piles []int, h int) int {
    maxPile := 0
    for _, pile := range piles{
        maxPile = max(maxPile, pile)
    }
    left, right := 1, maxPile

    var canEatAll func(speed int) bool
    canEatAll = func(speed int) bool{
        totalTime := 0
        for _, pile := range piles{
            totalTime += (pile + speed - 1) / speed
        }
        return totalTime <= h
    }

    for left < right{
        mid := left + (right - left) / 2
        if canEatAll(mid){
            right = mid
        } else {
            left = mid + 1
        }
    }
    return right
}

