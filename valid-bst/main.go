package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return CheckBST(root)
}

func CheckBST(node *TreeNode) bool,int,int {
	left := false
	right := false

	if node.Left == nil {
		left = true
	} else if node.Val > node.Left.Val && CheckBST(node.Left) {
		left = true
	}

	if node.Right == nil {
		right = true
	} else if node.Val < node.Right.Val && CheckBST(node.Right) {
		right = true
	}
	return left && right
}
