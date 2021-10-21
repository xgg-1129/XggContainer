package Heap

import (
	"fmt"
	"testing"
)

type MinHeap []int

func (h MinHeap) Swap(i, j int) {
	h[i],h[j] = h[j],h[i]
}

func (h MinHeap) Less(i, j int) bool {
	return h[i]>h[j]
}

func (receiver MinHeap) Len() int {
	return len(receiver)
}

func (receiver *MinHeap) Push(i interface{}) {
	*receiver = append(*receiver,i.(int))
}
func (receiver *MinHeap) Pop() interface{} {
	n:=len(*receiver)
	res:=(*receiver)[n-1]
	*receiver=(*receiver)[0:n-1]
	return res
}

func TestName(t *testing.T) {
	heap:=&MinHeap{2,1,5}
	Init(heap)
	Push(heap,3)
	fmt.Printf("minimum: %d\n", (*heap)[0])
	for heap.Len()>0{
		fmt.Println(heap.Pop())
	}
}