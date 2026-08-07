func isHappy(n int) bool {
    seen := make(map[int]bool)
	for n != 1{
		tmp := n
		cur := 0
		for tmp != 0{
			digit := tmp % 10
			cur += digit * digit
			tmp = tmp / 10
		}
		if seen[cur]{
			return false
		}
		seen[cur] = true
		n = cur
	}
	return true
}
