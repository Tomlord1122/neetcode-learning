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
		slow = nums[slow]
		head = nums[head]
		if slow == head{
			return head
		}
	}
	return -1
}
