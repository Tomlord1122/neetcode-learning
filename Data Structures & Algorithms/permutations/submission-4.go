func permute(nums []int) [][]int {
	res := [][]int{}
	perm := []int{}
	pick := make([]bool, len(nums))
	var backtrack func()
	backtrack = func(){
		if len(perm) == len(nums){
			tmp := make([]int, len(perm))
			copy(tmp, perm)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(nums); i++{
			if !pick[i]{
				perm = append(perm, nums[i])
				pick[i] = true
				backtrack()
				perm = perm[:len(perm)-1]
				pick[i] = false
			}
		}
	}
	backtrack()
	return res
}
