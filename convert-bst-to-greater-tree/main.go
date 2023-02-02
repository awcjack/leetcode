/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func greaterCalculation(root *TreeNode, sum int) int {
	if root.Right == nil && root.Left == nil {
		root.Val += sum
		return root.Val
	}
	if root.Right != nil {
		sum = greaterCalculation(root.Right, sum)
	}
	root.Val += sum
	sum = root.Val
	if root.Left != nil {
		sum = greaterCalculation(root.Left, sum)
	}
	return sum
}

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	greaterCalculation(root, 0)
	return root
}
