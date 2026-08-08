func permute(nums []int) [][]int {
	cur := []int{}
	res := [][]int{}
	pick := make(map[int]bool)
	

	var backtrack func()
	backtrack = func(){
		if len(cur) == len(nums){
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}
		for i := 0; i < len(nums); i++{
			if !pick[i]{
				cur = append(cur, nums[i])
				pick[i] = true
				backtrack()
				pick[i] = false
				cur = cur[:len(cur)-1]
			}
		}
	}
	backtrack()
	return res
}
