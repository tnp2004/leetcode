package binaryTreePreorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	order := new([]int)
	search(root, order)

	return *order
}

func search(root *TreeNode, order *[]int) {
	if root == nil {
		return
	}
	*order = append(*order, root.Val)
	if root.Left != nil {
		search(root.Left, order)
	}
	if root.Right != nil {
		search(root.Right, order)
	}
}
