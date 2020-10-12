package path_sum_II

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

var paths [][]int

func pathSum(root *TreeNode, sum int) [][]int {
	paths = [][]int{}
	if root != nil {
		constructPath(root, []int{}, sum)
	}
	return paths
}

func constructPath(root *TreeNode, path []int, sum int) {
	path = append(path, root.Val)
	sum -= root.Val
	if root.Left == nil && root.Right == nil {
		if sum == 0 {
			paths = append(paths, append([]int{}, path...))
		}
	} else {
		if root.Left != nil {
			constructPath(root.Left, path, sum)
		}
		if root.Right != nil {
			constructPath(root.Right, path, sum)
		}
	}
}