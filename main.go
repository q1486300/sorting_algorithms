package main

import (
	"fmt"
	"sorting_algorithms/bubble_sort"
	"sorting_algorithms/comparator"
	"sorting_algorithms/heap_sort"
	"sorting_algorithms/insertion_sort"
	"sorting_algorithms/merge_sort"
	"sorting_algorithms/quick_sort"
	"sorting_algorithms/selection_sort"
	"sorting_algorithms/sorting_strategy"
)

func main() {
	const maxSize = 100
	const maxValue = 100

	fmt.Println("------------------ Original array ------------------")
	arr := comparator.GenerateRandomArray(maxSize, maxValue)
	fmt.Println(arr)

	fmt.Println("------------------ SelectionSort ------------------")
	copyArr := comparator.CopyArray(arr)
	data := sorting_strategy.NewIntegerData(copyArr, &selection_sort.SelectionSort{})
	data.Sort()
	data.PrintArray()

	fmt.Println("------------------ BubbleSort ------------------")
	copyArr = comparator.CopyArray(arr)
	data.SetArr(copyArr)
	data.SetSortingAlgorithm(&bubble_sort.BubbleSort{})
	data.Sort()
	data.PrintArray()

	fmt.Println("------------------ InsertionSort ------------------")
	copyArr = comparator.CopyArray(arr)
	data.SetArr(copyArr)
	data.SetSortingAlgorithm(&insertion_sort.InsertionSort{})
	data.Sort()
	data.PrintArray()

	fmt.Println("------------------ MergeSort ------------------")
	copyArr = comparator.CopyArray(arr)
	data.SetArr(copyArr)
	data.SetSortingAlgorithm(&merge_sort.MergeSort{})
	data.Sort()
	data.PrintArray()

	fmt.Println("------------------ QuickSort ------------------")
	copyArr = comparator.CopyArray(arr)
	data.SetArr(copyArr)
	data.SetSortingAlgorithm(&quick_sort.QuickSort{})
	data.Sort()
	data.PrintArray()

	fmt.Println("------------------ HeapSort ------------------")
	copyArr = comparator.CopyArray(arr)
	data.SetArr(copyArr)
	data.SetSortingAlgorithm(&heap_sort.HeapSort{})
	data.Sort()
	data.PrintArray()
}
