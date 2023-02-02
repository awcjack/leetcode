/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Right != nil {
		root.Right = increasingBST(root.Right)
	}
	if root.Left != nil {
		temp := root
		root = increasingBST(root.Left)
		temp.Left = nil
		var max *TreeNode = root
		for {
			if max.Right != nil {
				max = max.Right
			} else {
				break
			}
		}
		max.Right = temp
	}
	return root
}
