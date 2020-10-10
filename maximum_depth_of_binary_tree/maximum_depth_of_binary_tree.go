package maximum_depth_of_binary_tree

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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	mld, mrd := maxDepth(root.Left), maxDepth(root.Right)
	if mld > mrd {
		return mld + 1
	}
	return mrd + 1
}
