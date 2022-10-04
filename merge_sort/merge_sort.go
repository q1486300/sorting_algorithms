package merge_sort

import "sorting_algorithms/sorting_strategy"

// 時間複雜度 O(N*logN)
// 空間複雜度 O(N)
type MergeSort struct {
}

var _ sorting_strategy.ISortingAlgorithm = &MergeSort{}

func (m *MergeSort) Sort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	m.process(arr, 0, len(arr)-1)
}

// 在 arr[L..R] 排好序
// L..R (N個數)
// T(N) = 2 * T(N / 2) + O(N)
// O(N * logN)
func (m *MergeSort) process(arr []int, L, R int) {
	if L == R { // 已有序，因為只剩一個數，base case
		return
	}
	mid := L + ((R - L) >> 1)
	m.process(arr, L, mid)
	m.process(arr, mid+1, R)
	m.merge(arr, L, mid, R)
}

func (m *MergeSort) merge(arr []int, L, M, R int) {
	help := make([]int, R-L+1)
	i := 0
	p1 := L
	p2 := M + 1
	for p1 <= M && p2 <= R {
		if arr[p1] <= arr[p2] {
			help[i] = arr[p1]
			p1++
		} else {
			help[i] = arr[p2]
			p2++
		}
		i++
	}
	// 不是 p1 越界，就是 p2 越界
	for p1 <= M {
		help[i] = arr[p1]
		p1++
		i++
	}
	for p2 <= R {
		help[i] = arr[p2]
		p2++
		i++
	}
	for i := 0; i < len(help); i++ {
		arr[L+i] = help[i]
	}
}
