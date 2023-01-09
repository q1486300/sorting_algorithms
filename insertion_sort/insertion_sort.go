package insertion_sort

import "sorting_algorithms/sorting_strategy"

// 時間複雜度 O(N^2); 資料狀況好的話(絕大部分已經有序)，時間複雜度O(N)
// 空間複雜度 O(1)
type InsertionSort struct {
}

var _ sorting_strategy.ISortingAlgorithm = &InsertionSort{}

func (data *InsertionSort) Sort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	// 0~0 有序的
	// 0~i 想有序
	for i := 1; i < len(arr); i++ { // 0 ~ i 做到有序
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			data.swap(arr, j, j+1)
		}
	}
}

func (data *InsertionSort) swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
