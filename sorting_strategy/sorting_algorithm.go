package sorting_strategy

import "fmt"

type ISortingAlgorithm interface {
	Sort([]int)
}

type IntegerData struct {
	arr             []int
	sortingBehavior ISortingAlgorithm
}

func (data *IntegerData) Sort() {
	if data.sortingBehavior == nil {
		panic("SetSortingAlgorithm before Sorting")
	}
	data.sortingBehavior.Sort(data.arr)
}

func (data *IntegerData) SetArr(arr []int) {
	data.arr = arr
}

func (data *IntegerData) SetSortingAlgorithm(sortingBehavior ISortingAlgorithm) {
	data.sortingBehavior = sortingBehavior
}

func (data *IntegerData) PrintArray() {
	fmt.Println(data.arr)
}

func NewIntegerData(arr []int, sortingAlgorithm ISortingAlgorithm) *IntegerData {
	return &IntegerData{
		arr:             arr,
		sortingBehavior: sortingAlgorithm,
	}
}
