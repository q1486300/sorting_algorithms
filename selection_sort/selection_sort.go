package selection_sort

import "sorting_algorithms/sorting_strategy"

// 時間複雜度 O(N^2)
// 空間複雜度 O(1)
type SelectionSort struct {
}

var _ sorting_strategy.ISortingAlgorithm = &SelectionSort{}

func (s *SelectionSort) Sort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	// 0 ~ N-1  找到最小值，之後跟原本在 0 位置上的值交換
	// 1 ~ N-1  找到最小值，之後跟原本在 1 位置上的值交換
	// 2 ~ N-1  找到最小值，之後跟原本在 2 位置上的值交換
	for i := 0; i < len(arr)-1; i++ { // 排到剩最後一個數值就不用比較了，不需要到 (N-1)
		minIndex := i
		for j := i + 1; j < len(arr); j++ { // i + 1 ~ N-1 上找最小值的索引
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		s.swap(arr, i, minIndex)
	}
}

func (s *SelectionSort) swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
