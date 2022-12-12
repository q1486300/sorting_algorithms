package heap_sort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestMyMaxHeap(t *testing.T) {
	testTime := 500000
	limit := 100
	value := 1000

testLabel:
	for i := 0; i < testTime; i++ {
		curLimit := rand.Intn(limit) + 1
		my := NewMyMaxHeap(curLimit)
		test := NewRightMaxHeap(curLimit)

		curOpTime := rand.Intn(limit) + 1
		for j := 0; j < curOpTime; j++ {
			if !assert.Equal(t, test.IsEmpty(), my.IsEmpty()) {
				break testLabel
			}
			if !assert.Equal(t, test.IsFull(), my.IsFull()) {
				break testLabel
			}
			if my.IsEmpty() {
				curValue := rand.Intn(value)
				my.Push(curValue)
				test.Push(curValue)
			} else if my.IsFull() {
				if !assert.Equal(t, test.Pop(), my.Pop()) {
					break testLabel
				}
			} else {
				if rand.Float64() < 0.5 {
					curValue := rand.Intn(value)
					my.Push(curValue)
					test.Push(curValue)
				} else {
					if !assert.Equal(t, test.Pop(), my.Pop()) {
						break testLabel
					}
				}
			}
		}
	}
}

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
