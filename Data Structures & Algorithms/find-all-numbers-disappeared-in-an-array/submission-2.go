func findDisappearedNumbers(nums []int) []int {
	numSet := make(map[int]bool)
	for _, num := range nums{
		numSet[num] = true
	}
	res := []int{}
	for i := 1; i <= len(nums); i++{
		if !numSet[i]{
			res = append(res, i)
		}
	}	
	return res
}
