/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)

	if l > r {
		return 1 + l
	}
	return 1 + r
}

func main() {
	// TODO
}
