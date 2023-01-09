package heap_sort

import "sorting_algorithms/sorting_strategy"

// 時間複雜度 O(N*logN)
// 空間複雜度 O(1)
type HeapSort struct {
}

var _ sorting_strategy.ISortingAlgorithm = &HeapSort{}

func (h *HeapSort) Sort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	// O(N*logN)
	//for i := 0; i < len(arr); i++ { // O(N)
	//	h.heapInsert(arr, i) // O(logN)
	//}

	heapSize := len(arr)

	// O(N)
	for i := len(arr) - 1; i >= 0; i-- {
		h.heapify(arr, i, heapSize)
	}

	for heapSize > 1 { // O(N)
		heapSize--
		h.swap(arr, 0, heapSize)
		h.heapify(arr, 0, heapSize) // O(logN)
	}
}

// arr[index] 剛來的數，往上調整
func (h *HeapSort) heapInsert(arr []int, index int) {
	for arr[index] > arr[(index-1)/2] {
		h.swap(arr, index, (index-1)/2)
		index = (index - 1) / 2
	}
}

// arr[index] 位置的數，能否往下移動(調整)
func (h *HeapSort) heapify(arr []int, index, heapSize int) {
	left := index*2 + 1   // 左孩子的索引
	for left < heapSize { // 下方還有孩子的時候
		// 一開始假設最大的是左孩子
		largest := left
		// 右孩子不越界，且大於左孩子，把右孩子的索引賦值給 largest
		if left+1 < heapSize && arr[left+1] > arr[left] {
			largest = left + 1
		}
		// 父節點的值較大，把父節點的索引賦值給 largest
		if arr[index] > arr[largest] {
			largest = index
		}
		// 最大值是父節點，不用移動(調整)
		if largest == index {
			break
		}
		h.swap(arr, largest, index)
		index = largest
		left = index*2 + 1
	}
}

func (h *HeapSort) swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
