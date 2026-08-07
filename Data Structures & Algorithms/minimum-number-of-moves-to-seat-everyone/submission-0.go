import "sort"
func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	res := 0
	for i := 0; i < len(students); i++{
		res += abs(seats[i] - students[i])
	}
	return res
}

func abs(x int) int{
	if x < 0{
		return -x
	}
	return x
}