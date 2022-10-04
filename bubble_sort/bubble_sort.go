package bubble_sort

import "sorting_algorithms/sorting_strategy"

// 時間複雜度 O(N^2)
// 空間複雜度 O(1)
type BubbleSort struct {
}

var _ sorting_strategy.ISortingAlgorithm = &BubbleSort{}

func (b *BubbleSort) Sort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	// 0 ~ N-1
	// 0 ~ N-2
	// 0 ~ N-3
	for e := len(arr) - 1; e > 0; e-- { // 0 ~ e
		for i := 0; i < e; i++ {
			if arr[i] > arr[i+1] {
				b.swap(arr, i, i+1)
			}
		}
	}
}

func (b *BubbleSort) swap(arr []int, i, j int) {
	// 能使用 ^ 交換兩個數的前提是: i、j 索引內的記憶體位址不同
	// 否則會把數值交換成 0
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]
}
