package heap_sort

import "container/heap"

// 已知一個幾乎有序的陣列，幾乎有序是指: 如果把陣列排好順序的話，每個元素移動的距離不可以超過 K，並且 K 小於陣列的長度
// 時間複雜度 O(N * logK); heap 內最多只有 K 個數，調整時時間複雜度 O(logK)
func SortedArrDistanceLessK(arr []int, k int) {
	if k == 0 {
		return
	}
	// 最小堆積樹
	h := NewMinHeap()

	index := 0
	// 0..K-1
	for ; index <= min(len(arr)-1, k-1); index++ {
		heap.Push(h, arr[index])
	}

	i := 0
	// 放入第 K 個數，並從 heap 中彈出最小值依序放到 i 位置
	for ; index < len(arr); index++ {
		heap.Push(h, arr[index])
		arr[i] = heap.Pop(h).(int)
		i++
	}

	// index 遍歷完整個陣列後，把 heap 中剩下的數彈出依序放到 i 位置
	for h.Len() != 0 {
		arr[i] = heap.Pop(h).(int)
		i++
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

type IntHeap []int

func NewMinHeap() *IntHeap {
	h := &IntHeap{}
	heap.Init(h)
	return h
}

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}
