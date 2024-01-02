package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	arr := make([]int, 0, 1)

	PreorderTraversal(root, arr, 0)

	return arr
}

func PreorderTraversal(node *TreeNode, arr []int, index int) {
	if node == nil {
		return
	}

	if len(arr) <= index {
		arr = append(arr, node.Val)
	} else {
		if arr[index] < node.Val {
			arr[index] = node.Val
		}
	}

	PreorderTraversal(node.Left, arr, index+1)
	PreorderTraversal(node.Right, arr, index+1)
}
