func minEatingSpeed(piles []int, h int) int {
    maxPile := 0
    for _, val := range piles{
        maxPile = max(maxPile, val)
    }

    left, right := 1, maxPile

    for left < right{
        mid := left + (right - left) / 2
        if canEat(piles, h, mid){
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}

func canEat (piles []int, h int, speed int) bool{
    for _, pile := range piles{
        time := (pile + speed - 1) / speed
        if time == 0{
            h -= 1
        } else {
            h -= time   
        }
    }
    if h < 0{
        return false
    }
    return true
}