/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    arr := []int{}
	// LDR
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode){
		if node == nil{
			return
		}
		inorder(node.Left)
		arr = append(arr, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return arr[k-1]
}

// use inorder traversal -> we can get the sorted order
// return len(arr) -k index's value
