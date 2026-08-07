/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    
	if root == nil{
		return true
	}
	var helper func(node *TreeNode) result
	helper = func(node *TreeNode) result{
		if node == nil{
			return result{
				height: 0,
				balanced: true,
			}
		}
		left, right := helper(node.Left), helper(node.Right)
		return result{
			height: 1 + max(left.height, right.height),
			balanced: left.balanced && right.balanced && abs(left.height-right.height) < 2,
		}
	}
	return helper(root).balanced
}

func abs(x int) int{
	if x > 0{
		return x
	}
	return -x
}
type result struct{
	height int
	balanced bool
}

// how to check a tree is balanced?
// abs(Left sub tree - right subtree) < 2 && left sub tree and right sub tree are balanced too
