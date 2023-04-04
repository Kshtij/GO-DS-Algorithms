package main

func main() {

}

type AllOne struct {
	StrMap   map[string]int
	CountMap map[int]map[string]interface{}
	
}

func Constructor() AllOne {

}

func (this *AllOne) Inc(key string) {
	node, strExists := this.StrMap[key]
	if !strExists {
		node = Node{
			Data : key
			Count: 1
		}
		_,countExists := this.CountMap[1]
		if !countExists{
			this.CountMap[1] = &node
		}

		if  this.ListHead == nil { 
			this.ListHead = &node
		}else {
			tailNode := this.ListHead.Prev

			tailNode.Next = &node
			node.Prev = tailNode

			node.Next = this.ListHead
		}
		
	}else{
		node.Count++
		node.Prev.Next = node.Next

		tempNode,countExists := this.CountMap[node.Count]
		if countExists{
			tempNode.Prev.Next = node
			node.next = tempNode
		}else{
			this.
		}

	}
}

func (this *AllOne) Dec(key string) {

}

func (this *AllOne) GetMaxKey() string {

}

func (this *AllOne) GetMinKey() string {

}

type Node struct {
	Data string
	Count int
	Next *Node
	Prev *Node
}

/*
node, strExists := this.StrMap[key]
	if !strExists {
		node = Node{
			Data : key
			Count: 1
		}
		_,countExists := this.CountMap[1]
		if !countExists{
			this.CountMap[1] = &node
		}

		if  this.ListHead == nil { 
			this.ListHead = &node
		}else {
			tailNode := this.ListHead.Prev

			tailNode.Next = &node
			node.Prev = tailNode

			node.Next = this.ListHead
		}
		
	}else{
		node.Count++
		node.Prev.Next = node.Next

		tempNode,countExists := this.CountMap[node.Count]
		if countExists{
			tempNode.Prev.Next = node
			node.next = tempNode
		}else{
			this.
		}

	}
*/