func countSeniors(details []string) int {
	res := 0
	for _, passenger := range details{
		n := len(passenger)
		age, _ := strconv.Atoi(passenger[n-4:n-2])
		if age > 60{
			res++
		}
	}
	return res
}
