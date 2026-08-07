func replaceElements(arr []int) []int {
	n := len(arr)
	postfix := make([]int, n)
	postfix[n-1] = -1
	for i := n-2; i >= 0; i--{
		postfix[i] = max(postfix[i+1], arr[i+1])
	}
	return postfix
}
