package binary_tree_preorder_traversal

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

func preorderTraversal(root *TreeNode) []int {
	var r []int
	if root == nil {
		return r
	}
	r = append(r, root.Val)
	r = append(r, preorderTraversal(root.Left)...)
	r = append(r, preorderTraversal(root.Right)...)
	return r
}
