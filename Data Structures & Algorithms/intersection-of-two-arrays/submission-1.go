func intersection(nums1 []int, nums2 []int) []int {
	numSet := make(map[int]bool)
	res := []int{}
	for _, val := range nums1{
		numSet[val] = true
	}
	for _, val := range nums2{
		if numSet[val]{
			res = append(res, val)
			numSet[val] = false
		}
	}
	return res
}
