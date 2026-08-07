func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	numIdx := make(map[int]int) // val -> index

	for idx, num := range nums2{
		numIdx[num] = idx
	}

	for i := 0; i < len(nums1); i++{
		res[i] = -1 
		for j := numIdx[nums1[i]]; j < len(nums2); j++{
			if nums2[j] > nums1[i]{
				res[i] = nums2[j]
				break
			}
		}
	}
	return res
}



