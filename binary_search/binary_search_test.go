package binary_search

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"sorting_algorithms/comparator"
	"testing"
)

func TestBinarySearchExist(t *testing.T) {
	testTime := 500000
	maxSize := 10
	maxValue := 100

	for i := 0; i < testTime; i++ {
		arr := comparator.GenerateRandomArray(maxSize, maxValue)
		sort.Ints(arr)

		value := rand.Intn(maxValue+1) - rand.Intn(maxValue+1)

		b1 := Exist(arr, value)
		b2 := comparator.SearchingExistComparator(arr, value)

		if !assert.Equal(t, b2, b1) {
			break
		}
	}
}

func TestNearestLeftIndex(t *testing.T) {
	testTime := 500000
	maxSize := 10
	maxValue := 100

	for i := 0; i < testTime; i++ {
		arr := comparator.GenerateRandomArray(maxSize, maxValue)
		sort.Ints(arr)

		value := rand.Intn(maxValue+1) - rand.Intn(maxValue+1)

		i1 := NearestLeftIndex(arr, value)
		i2 := comparator.SearchingNearestLeftIndexComparator(arr, value)

		if !assert.Equal(t, i2, i1) {
			for _, cur := range arr {
				fmt.Printf("%d ", cur)
			}
			fmt.Println()
			fmt.Println(value)
			break
		}
	}
}

func TestNearestRightIndex(t *testing.T) {
	testTime := 500000
	maxSize := 10
	maxValue := 100

	for i := 0; i < testTime; i++ {
		arr := comparator.GenerateRandomArray(maxSize, maxValue)
		sort.Ints(arr)

		value := rand.Intn(maxValue+1) - rand.Intn(maxValue+1)

		i1 := NearestRightIndex(arr, value)
		i2 := comparator.SearchingNearestRightIndexComparator(arr, value)

		if !assert.Equal(t, i2, i1) {
			for _, cur := range arr {
				fmt.Printf("%d ", cur)
			}
			fmt.Println()
			fmt.Println(value)
			break
		}
	}
}

func TestGetLessIndex(t *testing.T) {
	testTime := 500000
	maxSize := 30
	maxValue := 100

	for i := 0; i < testTime; i++ {
		arr := comparator.GenerateNotEqualRandomArray(maxSize, maxValue)
		ans := GetLessIndex(arr)

		if !assert.True(t, comparator.LessIndexComparator(arr, ans)) {
			break
		}
	}
}
