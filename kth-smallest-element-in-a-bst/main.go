/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type Result struct {
	Index int
	Val   int
}

func inOrderTraversal(current *TreeNode, k int, result *Result) {
	if current.Left != nil {
		inOrderTraversal(current.Left, k, result)
	}

	if result.Index == k {
		return
	}

	result.Index++
	result.Val = current.Val

	if result.Index == k {
		return
	}

	if current.Right != nil {
		inOrderTraversal(current.Right, k, result)
	}
}

func kthSmallest(root *TreeNode, k int) int {
	result := Result{0, 0}
	inOrderTraversal(root, k, &result)
	return result.Val
}
