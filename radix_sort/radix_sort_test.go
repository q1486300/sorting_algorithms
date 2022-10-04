package radix_sort

import (
	"github.com/stretchr/testify/assert"
	"sorting_algorithms/comparator"
	"sorting_algorithms/sorting_strategy"
	"testing"
)

func TestRadixSort(t *testing.T) {
	testTime := 500000
	maxSize := 100
	maxValue := 100000

	for i := 0; i < testTime; i++ {
		arr1 := comparator.GenerateNoNegativeRandomArray(maxSize, maxValue)
		arr2 := comparator.CopyArray(arr1)

		data := sorting_strategy.NewIntegerData(arr1, &RadixSort{})
		data.Sort()

		comparator.SortingComparator(arr2)

		if !assert.Equal(t, arr2, arr1) {
			break
		}
	}

	arr := comparator.GenerateNoNegativeRandomArray(maxSize, maxValue)
	data := sorting_strategy.NewIntegerData(arr, &RadixSort{})
	data.PrintArray()
	data.Sort()
	data.PrintArray()
}
