package balanced_binary_tree

import "math"

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


func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}

func height(root *TreeNode) int {
	// postorder
	if root == nil {
		return 0
	}
	lh := height(root.Left)
	rh := height(root.Right)
	if lh == -1 || rh == -1 || math.Abs(float64(lh - rh)) > 1 {
		return -1
	} else {
		return max(height(root.Left), height(root.Right)) + 1
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}