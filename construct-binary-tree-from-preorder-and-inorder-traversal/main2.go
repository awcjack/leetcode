/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func indexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{preorder[0], nil, nil}
	}

	root := &TreeNode{preorder[0], nil, nil}
	rootPos := indexOf(preorder[0], inorder)

	if rootPos != -1 {
		root.Left = buildTree(preorder[1:rootPos+1], inorder[:rootPos])
		root.Right = buildTree(preorder[rootPos+1:], inorder[rootPos+1:])
	}
	return root
}
