func plusOne(digits []int) []int {
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1{
		digits[i], digits[j] = digits[j], digits[i]
	}
	car := 1
	i := 0
	for car != 0{
		if digits[i] < 9{
			digits[i]++
			car = 0
		} else if i == len(digits)-1{
			digits[i] = 0
			digits = append(digits, 1)
			car = 0
		} else {
			digits[i] = 0
		}
		i++
	}

	// swap the order back
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1{
		digits[i], digits[j] = digits[j], digits[i]
	}
	return digits
}
