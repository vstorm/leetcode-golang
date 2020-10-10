package binary_tree_inorder_traversal

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

func inorderTraversal(root *TreeNode) []int {
	var r []int
	if root == nil {
		return r
	}
	r = append(r, inorderTraversal(root.Left)...)
	r = append(r, root.Val)
	r = append(r, inorderTraversal(root.Right)...)
	return r
}
