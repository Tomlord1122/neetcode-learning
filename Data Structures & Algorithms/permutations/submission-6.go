func permute(nums []int) [][]int {
	res := [][]int{}
	cur := []int{}
	pick := make(map[int]bool)

	var backtrack func()
	backtrack = func(){
		// WHat's the base case
		if len(cur) == len(nums){
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}
		for i := 0; i < len(nums); i++{
			if !pick[i]{
				pick[i] = true
				cur = append(cur, nums[i])
				backtrack()
				pick[i] = false
				cur = cur[:len(cur)-1]
			}
		}
	}
	backtrack()
	return res
}
