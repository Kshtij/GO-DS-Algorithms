package main

func main() {

}

type SnapshotArray struct {
	currentSnapShotID int
	index             map[string]int
}

var keyfmtStr = "%d-%d"

func Constructor(length int) SnapshotArray {
	return SnapshotArray{
		currentSnapShotID: 0,
		index:             make(map[string]int),
	}
}

func (this *SnapshotArray) Set(index int, val int) {

}

func (this *SnapshotArray) Snap() int {

}

func (this *SnapshotArray) Get(index int, snap_id int) int {

}
