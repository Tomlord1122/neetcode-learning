/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
    indices := make(map[int]int) // value to index from inorder
	for i, v := range inorder{
		indices[v] = i
	}

	preIdx := 0
	var dfs func(l, r int) *TreeNode
	dfs = func(l, r int) *TreeNode{
		if l > r{
			return nil
		}
		val := preorder[preIdx]
		preIdx++
		node := &TreeNode{Val:val}
		midIdx := indices[val]
		node.Left = dfs(l, midIdx-1)
		node.Right = dfs(midIdx+1, r)
		return node
	}
	root := dfs(0, len(preorder)-1)
	return root
}
