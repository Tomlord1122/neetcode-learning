func search(nums []int, target int) int {
    left, right := 0, len(nums)-1

    // [1]
    for left <= right{
        mid := left + (right - left) / 2
        if nums[mid] == target{
            return mid
        }

        if nums[left] <= nums[mid]{
            if target > nums[mid] || target < nums[left]{
                left = mid + 1
            // target < nums[mid] && target > nums[left]
            } else {
                right = mid - 1
            }
        // right sorted portion
        } else {
            if target < nums[mid] || target > nums[right]{
                right = mid - 1
            // target > nums[mid] && target < nums[right]
            } else {
                left = mid + 1
            }
        }
       
    }
    return -1
}
