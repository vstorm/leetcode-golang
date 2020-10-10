package binary_tree_level_order_traversal

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

func levelOrder(root *TreeNode) [][]int {
	var r [][]int
	if root == nil {
		return r
	}

	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		r = append(r, []int{})
		var p []*TreeNode
		for j := 0; j < len(q); j++ {
			t := q[j]
			r[i] = append(r[i], t.Val)
			if t.Left != nil {
				p = append(p, t.Left)
			}
			if t.Right != nil {
				p = append(p, t.Right)
			}
		}
		q = p
	}
	return r
}
