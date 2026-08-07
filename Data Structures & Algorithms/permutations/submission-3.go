func permute(nums []int) [][]int {
	if len(nums) == 0{
		return [][]int{{}}
	}
	perms := permute(nums[1:])
	res := [][]int{}
	for _, p := range perms{
		for i := 0; i <= len(p); i++{
			// prepare left, right
			left := append([]int{}, p[:i]...)
			right := append([]int{nums[0]}, p[i:]...)
			merge := append(left, right...)
			res = append(res, merge)
		}
	}
	return res
}
