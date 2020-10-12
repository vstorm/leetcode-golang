package minimum_depth_of_binary_tree

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

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left != nil && root.Right != nil {
		minLeftDepth := minDepth(root.Left) + 1
		minRightDepth := minDepth(root.Right) + 1
		if minLeftDepth < minRightDepth {
			return minLeftDepth
		}
		return minRightDepth
	} else if root.Left != nil {
		return minDepth(root.Left) + 1
	} else if root.Right != nil {
		return minDepth(root.Right) + 1
	} else {
		return 1
	}
}