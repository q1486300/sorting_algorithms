package selection_sort

import (
	"github.com/stretchr/testify/assert"
	"sorting_algorithms/comparator"
	"sorting_algorithms/sorting_strategy"
	"testing"
)

func TestSelectSort(t *testing.T) {
	testTime := 500000
	maxSize := 100
	maxValue := 100

	for i := 0; i < testTime; i++ {
		arr1 := comparator.GenerateRandomArray(maxSize, maxValue)
		arr2 := comparator.CopyArray(arr1)

		data := sorting_strategy.NewIntegerData(arr1, &SelectionSort{})
		data.Sort()

		comparator.SortingComparator(arr2)

		if !assert.Equal(t, arr2, arr1) {
			break
		}
	}

	arr := comparator.GenerateRandomArray(maxSize, maxValue)
	data := sorting_strategy.NewIntegerData(arr, &SelectionSort{})
	data.PrintArray()
	data.Sort()
	data.PrintArray()
}
