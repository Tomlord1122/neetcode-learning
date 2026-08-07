func isHappy(n int) bool {
    seen := make(map[int]bool)
	for n != 1{
		tmp := n
		res := 0
		for tmp != 0{
			v := tmp % 10
			tmp = tmp / 10
			res = res + v*v
		}
		if seen[res]{
			return false
		}
		seen[res] =true
		n = res
	}
	return true
}
