/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root.Val >= low && root.Val <= high {
		// root doesn't change
		if root.Left != nil {
			root.Left = trimBST(root.Left, low, high)
		}
		if root.Right != nil {
			root.Right = trimBST(root.Right, low, high)
		}
	} else if root.Val < low {
		if root.Right != nil {
			root = trimBST(root.Right, low, high)
		} else {
			root = nil
		}
	} else if root.Val > high {
		if root.Left != nil {
			root = trimBST(root.Left, low, high)
		} else {
			root = nil
		}
	}
	return root
}
