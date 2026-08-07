func combinationSum(nums []int, target int) [][]int {
    cur := []int{}
	res := [][]int{}

	var bt func(i, sum int)
	bt = func(i, sum int){
		if sum == 0{
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}
		if i == len(nums) || sum < 0{
			return
		}
		cur = append(cur, nums[i])
		bt(i, sum-nums[i])
		cur = cur[:len(cur)-1]
		bt(i+1, sum)
	}
	bt(0, target)
	return res
}
