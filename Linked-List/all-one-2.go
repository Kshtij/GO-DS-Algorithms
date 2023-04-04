package main

func main() {

}

type AllOne struct {
	StrMap   map[string]int
	CountMap map[int]map[string]interface{}

	ListHead *Node
}

func Constructor() AllOne {

}

func (this *AllOne) Inc(key string) {

}

func (this *AllOne) Dec(key string) {

}

func (this *AllOne) GetMaxKey() string {

}

func (this *AllOne) GetMinKey() string {

}

type Node struct {
	Count int
	Next  *Node
	Prev  *Node
}
