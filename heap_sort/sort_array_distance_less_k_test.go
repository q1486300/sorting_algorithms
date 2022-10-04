package heap_sort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sorting_algorithms/comparator"
	"testing"
)

func TestSortedArrDistanceLessK(t *testing.T) {
	testTime := 500000
	maxSize := 100
	maxValue := 100

	for i := 0; i < testTime; i++ {
		k := rand.Intn(maxSize-1) + 1
		arr := comparator.RandomArrayNoMoveMoreK(maxSize, maxValue, k)
		arr1 := comparator.CopyArray(arr)
		arr2 := comparator.CopyArray(arr)

		SortedArrDistanceLessK(arr1, k)
		comparator.SortingComparator(arr2)

		if !assert.Equal(t, arr2, arr1) {
			fmt.Printf("k = %d\n", k)
			comparator.PrintArray(arr)
			break
		}
	}
}
