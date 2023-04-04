package main

import "log"

func main() {
	//log.Println(generateTrees(5))
	arr := generateTrees(3)
	for _, node := range arr {
		PrintTree(node)
		log.Println()
	}
}

func PrintTree(node *TreeNode) {
	if node == nil {

		return
	}
	PrintTree(node.Left)
	log.Print(node.Val)
	PrintTree(node.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}
	retList := make([]*TreeNode, 0, 1)
	for i := 0; i < n; i++ {
		retList = append(retList, generate(arr[i], arr[:i], arr[i+1:])...)
	}
	return retList
}

func generate(element int, leftElements, rightElements []int) []*TreeNode {
	retList := make([]*TreeNode, 0, 1)

	if len(leftElements) == 0 && len(rightElements) == 0 {
		node := &TreeNode{
			Val: element,
		}
		retList = append(retList, node)
		return retList
	}

	var leftSubRootNodeList, rightSubRootNodeList []*TreeNode
	for index := range leftElements {
		leftSubRootNodeList = append(leftSubRootNodeList, generate(leftElements[index], leftElements[:index], leftElements[index+1:])...)
	}

	for index := range rightElements {
		rightSubRootNodeList = append(rightSubRootNodeList, generate(rightElements[index], rightElements[:index], rightElements[index+1:])...)
	}

	if len(rightSubRootNodeList) == 0 {
		for i := 0; i < len(leftSubRootNodeList); i++ {
			newRoot := &TreeNode{
				Val:   element,
				Left:  leftSubRootNodeList[i],
				Right: nil,
			}
			retList = append(retList, newRoot)
		}
	} else if len(leftSubRootNodeList) == 0 {
		for j := 0; j < len(rightSubRootNodeList); j++ {
			newRoot := &TreeNode{
				Val:   element,
				Left:  nil,
				Right: rightSubRootNodeList[j],
			}
			retList = append(retList, newRoot)
		}
	} else {
		for i := 0; i < len(leftSubRootNodeList); i++ {
			for j := 0; j < len(rightSubRootNodeList); j++ {
				newRoot := &TreeNode{
					Val:   element,
					Left:  leftSubRootNodeList[i],
					Right: rightSubRootNodeList[j],
				}
				retList = append(retList, newRoot)
			}
		}
	}
	return retList
}
