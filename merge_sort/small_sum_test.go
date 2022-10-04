package merge_sort

import (
	"github.com/stretchr/testify/assert"
	"sorting_algorithms/comparator"
	"testing"
)

func TestSmallSum(t *testing.T) {
	testTime := 500000
	maxSize := 100
	maxValue := 100

	for i := 0; i < testTime; i++ {
		arr1 := comparator.GenerateRandomArray(maxSize, maxValue)
		arr2 := comparator.CopyArray(arr1)

		s1 := SmallSum(arr1)
		s2 := comparator.SmallSumComparator(arr2)

		if !assert.Equal(t, s2, s1) {
			break
		}
	}
}
