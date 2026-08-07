func replaceElements(arr []int) []int {
	postfix := make([]int, len(arr))
	postfix[len(arr)-1] = -1
	for i := len(arr)-2; i >= 0; i--{
		postfix[i] = max(postfix[i+1], arr[i+1])
	}

	return postfix
}
