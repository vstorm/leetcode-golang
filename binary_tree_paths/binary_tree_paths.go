package binary_tree_paths

import "strconv"

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

var paths []string

func binaryTreePaths(root *TreeNode) []string {
	paths = []string{}
	if root != nil {
		constructPath(root, "")
	}
	return paths
}

func constructPath(root *TreeNode, path string) {
	path += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		paths = append(paths, path)
	} else {
		path += "->"
		if root.Left != nil {
			constructPath(root.Left, path)
		}
		if root.Right != nil {
			constructPath(root.Right, path)
		}
	}
}