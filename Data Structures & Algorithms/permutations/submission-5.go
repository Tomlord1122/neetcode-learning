func permute(nums []int) [][]int {
	res := [][]int{}
	cur := []int{}
	pick := make([]bool, len(nums))

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
				cur = cur[:len(cur)-1]
				pick[i] = false
			}
		}
	}
	backtrack()
	return res
}
