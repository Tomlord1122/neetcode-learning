func isHappy(n int) bool {
    seen := make(map[int]bool)
	for n != 1{
		cur := 0
		tmp := n
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
