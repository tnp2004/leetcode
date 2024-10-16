package binaryTreeInorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	order := new([]int)
	search(root, order)

	return *order
}

func search(root *TreeNode, order *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		search(root.Left, order)
	}
	*order = append(*order, root.Val)
	if root.Right != nil {
		search(root.Right, order)
	}
}
