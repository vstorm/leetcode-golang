package sum_root_to_leaf_numbers

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pathSums []int

func sumNumbers(root *TreeNode) int {
	pathSums = []int{}
	if root != nil {
		constructPath(root, 0)
	}
	sum := 0
	for _, s := range pathSums {
		sum += s
	}
	return sum
}

func constructPath(root *TreeNode, pathSum int) {
	pathSum = 10 * pathSum + root.Val
	if root.Left == nil && root.Right == nil {
		pathSums = append(pathSums, pathSum)
	} else {
		if root.Left != nil {
			constructPath(root.Left, pathSum)
		}
		if root.Right != nil {
			constructPath(root.Right, pathSum)
		}
	}
}