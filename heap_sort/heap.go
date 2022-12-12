package heap_sort

import "fmt"

type MyMaxHeap struct {
	heap     []int
	limit    int
	heapSize int
}

func NewMyMaxHeap(limit int) *MyMaxHeap {
	return &MyMaxHeap{
		heap:     make([]int, limit),
		limit:    limit,
		heapSize: 0,
	}
}

func (m *MyMaxHeap) IsEmpty() bool {
	return m.heapSize == 0
}

func (m *MyMaxHeap) IsFull() bool {
	return m.heapSize == m.limit
}

func (m *MyMaxHeap) Push(value int) {
	if m.IsFull() {
		fmt.Printf("heap is full, value:%d can't be pushed into heap.", value)
		return
	}
	m.heap[m.heapSize] = value
	m.heapInsert()
	m.heapSize++
}

func (m *MyMaxHeap) Pop() int {
	ans := m.heap[0]
	m.heapSize--
	m.swap(0, m.heapSize)
	m.heapify()
	return ans
}

func (m *MyMaxHeap) heapInsert() {
	index := m.heapSize
	for m.heap[index] > m.heap[(index-1)/2] {
		m.swap(index, (index-1)/2)
		index = (index - 1) / 2
	}
}

func (m *MyMaxHeap) heapify() {
	index := 0
	left := index*2 + 1
	for left < m.heapSize {
		largest := left
		if left+1 < m.heapSize && m.heap[left+1] > m.heap[left] {
			largest = left + 1
		}
		if m.heap[index] > m.heap[largest] {
			largest = index
		}
		if largest == index {
			break
		}
		m.swap(largest, index)
		index = largest
		left = index*2 + 1
	}
}

func (m *MyMaxHeap) swap(i, j int) {
	m.heap[i], m.heap[j] = m.heap[j], m.heap[i]
}
