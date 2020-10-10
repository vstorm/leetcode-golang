package convert_sorted_array_to_binary_search_tree

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

func sortedArrayToBST(nums []int) *TreeNode {
	return dfs(nums, 0, len(nums) - 1)
}

func dfs(nums []int, left, right int) *TreeNode {
	if left < right {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{
		Val: nums[mid],
	}
	root.Left = dfs(nums, left, mid-1)
	root.Right = dfs(nums, mid+1, right)
	return root
}
