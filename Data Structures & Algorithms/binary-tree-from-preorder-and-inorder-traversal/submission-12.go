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
		rootVal := preorder[preIdx]
		root := &TreeNode{Val:rootVal}
		rootIdx := indices[rootVal]
		preIdx++
		root.Left = dfs(l, rootIdx-1)
		root.Right = dfs(rootIdx+1, r)
		return root
	}

	return dfs(0, len(preorder)-1)
}
