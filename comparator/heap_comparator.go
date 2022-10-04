package comparator

import "fmt"

type RightMaxHeap struct {
	arr   []int
	limit int
	size  int
}

func NewRightMaxHeap(limit int) *RightMaxHeap {
	return &RightMaxHeap{
		arr:   make([]int, limit),
		limit: limit,
		size:  0,
	}
}

func (r *RightMaxHeap) IsEmpty() bool {
	return r.size == 0
}

func (r *RightMaxHeap) IsFull() bool {
	return r.size == r.limit
}

func (r *RightMaxHeap) Push(value int) {
	if r.size == r.limit {
		fmt.Printf("heap is full, value:%d can't push into heap.", value)
	}
	r.arr[r.size] = value
	r.size++
}

func (r *RightMaxHeap) Pop() int {
	maxIndex := 0
	for i := 1; i < r.size; i++ {
		if r.arr[i] > r.arr[maxIndex] {
			maxIndex = i
		}
	}
	ans := r.arr[maxIndex]
	r.size--
	r.arr[maxIndex] = r.arr[r.size]
	return ans
}
