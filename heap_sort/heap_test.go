package heap_sort

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sorting_algorithms/comparator"
	"testing"
)

func TestMyMaxHeap(t *testing.T) {
	testTime := 1000000
	limit := 100
	value := 1000

testLabel:
	for i := 0; i < testTime; i++ {
		curLimit := rand.Intn(limit) + 1
		my := NewMyMaxHeap(curLimit)
		test := comparator.NewRightMaxHeap(curLimit)

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
