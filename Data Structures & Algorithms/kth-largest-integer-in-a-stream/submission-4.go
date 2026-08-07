type KthLargest struct {
    k int
	nums []int
}


func Constructor(k int, nums []int) KthLargest {
    return KthLargest{
		k: k,
		nums: nums,
	}
}


func (this *KthLargest) Add(val int) int {
    this.nums = append(this.nums, val)
	sort.Slice(this.nums, func(i, j int) bool{
		return this.nums[i] > this.nums[j]
	})
	return this.nums[this.k - 1]
}
