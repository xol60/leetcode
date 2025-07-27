package main

type Data struct {
	diameter int
	line     int
}

func recursion_3(head *TreeNode) Data {
	if head == nil {
		return Data{
			diameter: 0,
			line:     0,
		}
	}
	a := recursion_3(head.Left)
	b := recursion_3(head.Right)
	h := maxInt(a.line+b.line+1, maxInt(a.diameter, b.diameter))
	return Data{
		diameter: h,
		line:     maxInt(a.line, b.line) + 1,
	}
}
func diameterOfBinaryTree(root *TreeNode) int {
	left := recursion_3(root.Left)
	right := recursion_3(root.Right)
	return maxInt(left.line+right.line+1, maxInt(left.diameter, right.diameter)) - 1
}
