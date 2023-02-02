/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	val := root.Val

	if p.Val < val && q.Val < val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > val && q.Val > val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
