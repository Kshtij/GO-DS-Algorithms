package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	heap := LevelOrderTraversal(root, p, q)

	pIndex := -1
	qIndex := -1

	for i := 0; i < len(heap) && (pIndex == -1 || qIndex == -1); i++ {
		if heap[i] == p {
			pIndex = i
		}
		if heap[i] == q {
			qIndex = i
		}
	}

	pDepth := getHeight(pIndex)
	qDepth := getHeight(qIndex)

	for pDepth > qDepth {
		if pIndex%2 == 0 {
			pIndex = pIndex - 2
		} else {
			pIndex = pIndex - 1
		}
		pIndex = pIndex / 2
		pDepth = getHeight(pIndex)
	}

	for qDepth > pDepth {
		if qIndex%2 == 0 {
			qIndex = qIndex - 2
		} else {
			qIndex = qIndex - 1
		}
		qIndex = qIndex / 2
		qDepth = getHeight(qIndex)
	}

	for pIndex != qIndex {
		if pIndex%2 == 0 {
			pIndex = pIndex - 2
		} else {
			pIndex = pIndex - 1
		}
		pIndex = pIndex / 2

		if qIndex%2 == 0 {
			qIndex = qIndex - 2
		} else {
			qIndex = qIndex - 1
		}
		qIndex = qIndex / 2
	}
	return heap[pIndex]
}

func LevelOrderTraversal(node, node1, node2 *TreeNode) []*TreeNode {
	heap := make([]*TreeNode, 0, 1)

	q := Queue{
		Arr: make([]*TreeNode, 0, 1),
		End: -1,
	}
	q.Push(node)

	var node1Done, node2Done bool

	for len(q.Arr) > 0 && (!node1Done || !node2Done) {
		node := q.Pop()
		if node == node1 {
			node1Done = true
		}

		if node == node2 {
			node2Done = true
		}
		heap = append(heap, node)
		if node != nil {
			q.Push(node.Left)
			q.Push(node.Right)
		} else {
			q.Push(nil)
			q.Push(nil)
		}
	}

	return heap
}

func getHeight(index int) int {
	if index%2 == 0 {
		index = index - 2
	} else {
		index = index - 1
	}

	height := 0
	for index > 0 {
		index = index >> 1
	}
	return height
}

type Queue struct {
	Arr []*TreeNode
	End int
}

func (q *Queue) Push(node *TreeNode) {
	q.Arr = append(q.Arr, node)
	q.End = q.End + 1
}

func (q *Queue) Pop() *TreeNode {
	node := q.Arr[0]
	q.Arr = q.Arr[1:]
	q.End = q.End - 1
	return node
}
