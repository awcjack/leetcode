/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func translatePointerToVal(input [][]*TreeNode) [][]int {
	result := make([][]int, len(input))
	for i := 0; i < len(input); i++ {
		result[i] = make([]int, len(input[i]))
		for j := 0; j < len(input[i]); j++ {
			result[i][j] = input[i][j].Val
		}
	}
	return result
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	if root.Left == nil && root.Right == nil {
		return [][]int{[]int{root.Val}}
	}
	level := 0
	nodesOnAllLevels := [][]*TreeNode{[]*TreeNode{root}}
	for {
		//if len(nodesOnAllLevels) == level + 1 {
		nextLevel := []*TreeNode{}
		for i := 0; i < len(nodesOnAllLevels[level]); i++ {
			if nodesOnAllLevels[level][i].Left != nil {
				nextLevel = append(nextLevel, nodesOnAllLevels[level][i].Left)
			}
			if nodesOnAllLevels[level][i].Right != nil {
				nextLevel = append(nextLevel, nodesOnAllLevels[level][i].Right)
			}
		}
		if len(nextLevel) == 0 {
			break
		}
		nodesOnAllLevels = append(nodesOnAllLevels, nextLevel)
		level++
		//}
	}

	return translatePointerToVal(nodesOnAllLevels)
}
