package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recursion(node *TreeNode, max int) int {
	if node == nil {
		return 0
	}
	n := 0
	if node.Val >= max {
		max = node.Val
		n = 1
	}
	return n + recursion(node.Left, max) + recursion(node.Right, max)
}
func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return recursion(root, math.MinInt32)
}
