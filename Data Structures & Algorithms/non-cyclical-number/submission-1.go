func isHappy(n int) bool {
    seen := make(map[int]bool)
    for n != 1{
        tmp := n
        res := 0
        for tmp != 0{
            b := tmp % 10
            res = res + b * b
            tmp /= 10
        }
        n = res
        if seen[res]{
            return false
        }
        seen[res] = true
    }
    return true
}
