func findMin(nums []int) int {
    left, right := 0, len(nums)-1
    // if the question doesn't give us the target perimeter, it represents that there's no specific number we want to find. So in this case,
    // I think we can design a for loop condition just less, not less or equal.
    for left < right{
        mid := left + (right- left) / 2
        if nums[mid] < nums[right]{
            right = mid
        } else {
            left = mid + 1
        }
    }
    // it doesn't matter we return the left index or the right index 
    // at the final output because they are the same in this case.

    return nums[right]
}




// We should find a method to check which part the mid is
