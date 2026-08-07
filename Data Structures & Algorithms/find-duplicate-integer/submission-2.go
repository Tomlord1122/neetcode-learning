func findDuplicate(nums []int) int {
    slow, fast := 0, 0
	for{
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast{
			break
		}
	}
	head := 0
	for{
		head = nums[head]
		slow = nums[slow]
		if slow == head{
			break
		}
	}
	return head
}
