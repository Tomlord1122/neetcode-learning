func isHappy(n int) bool {
	seen := make(map[int]bool)

	for n != 1{
		tmp := n
		sum := 0
		for tmp != 0{
			sum += (tmp % 10) * (tmp % 10)
			tmp /= 10
		}
		if seen[sum]{
			return false
		}
		seen[sum] = true
		n = sum
	}
	return true
}



// Given a positive integer, replace it with the sum of the squares of its digits.


// Repeat the above step until the number equals 1, or it loops infinitely in a cycle which does not include 1.
// If it stops at 1, then the number is a non-cyclical number.
