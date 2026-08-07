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
		left := helper(node.Left)
		right := helper(node.Right)

		return result{
			height: 1 + max(left.height, right.height),
			balanced: left.balanced && right.balanced && abs(left.height - right.height) < 2,
		}
	}
	return helper(root).balanced
}

type result struct{
	height int
	balanced bool
}


func abs(x int) int{
	if x < 0{
		return -x
	}
	return x
}