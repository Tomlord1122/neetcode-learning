/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
    indices := make(map[int]int)
	for i, v := range inorder{
		indices[v] = i
	}
	preIdx := 0
	var dfs func(l, r int) *TreeNode
	dfs = func(l, r int) *TreeNode{
		if l > r{
			return nil
		}
		root := &TreeNode{Val: preorder[preIdx]}
		mid := indices[preorder[preIdx]]
		preIdx++
		root.Left = dfs(l, mid-1)
		root.Right = dfs(mid+1, r)
		return root
	}
	return dfs(0, len(preorder)-1)
}
