package heap

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// Time: O(n)
func buildHeapByInit(array []int) *MaxHeap {
	maxHeap := &MaxHeap{}
	*maxHeap = array
	heap.Init(maxHeap)
	fmt.Println("buildHeapByInit: ", *maxHeap)
	return maxHeap
}

// Time: O(nlogn)
func buildHeapByPush(array []int) *MaxHeap {
	maxHeap := &MaxHeap{}
	for _, num := range array {
		heap.Push(maxHeap, num)
	}
	fmt.Println("buildHeapByPush: ", *maxHeap)
	return maxHeap
}
