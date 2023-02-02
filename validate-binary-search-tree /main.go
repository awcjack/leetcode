/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return isValidBSTwithBoundary(root, math.MinInt32, math.MaxInt32)
}

func isValidBSTwithBoundary(root *TreeNode, lowerBoundary int, upperBoundary int) bool {
	if root.Val > upperBoundary || root.Val < lowerBoundary {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return true
	}

	var left, right bool
	if root.Left != nil {
		if root.Left.Val >= root.Val {
			return false
		}
		left = isValidBSTwithBoundary(root.Left, lowerBoundary, root.Val-1)
	} else {
		left = true
	}

	if root.Right != nil {
		if root.Right.Val <= root.Val {
			return false
		}
		right = isValidBSTwithBoundary(root.Right, root.Val+1, upperBoundary)
	} else {
		right = true
	}

	if left && right {
		return true
	}
	return false
}
