/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inOrderTraversal(current *TreeNode, allVal *[]int) {
	if current.Left != nil {
		inOrderTraversal(current.Left, allVal)
	}

	*allVal = append(*allVal, current.Val)

	if current.Right != nil {
		inOrderTraversal(current.Right, allVal)
	}
}

func kthSmallest(root *TreeNode, k int) int {
	var allVal []int
	inOrderTraversal(root, &allVal)
	sort.Ints(allVal)
	return allVal[k-1]
}
