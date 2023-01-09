package quick_sort

import (
	"math/rand"
	"sorting_algorithms/sorting_strategy"
)

// 時間複雜度 O(N*logN)，前提是劃分值隨機取; 如果劃分值取得不好(數值大小太靠右或太靠左)，時間複雜度會變成 O(N^2)
// 空間複雜度 O(logN)，遞迴調用的系統堆疊，前提是劃分值隨機取; 如果劃分值取得不好(數值大小太靠右或太靠左)，空間複雜度會變成 O(N)
type QuickSort struct {
}

var _ sorting_strategy.ISortingAlgorithm = &QuickSort{}

func (q *QuickSort) Sort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	q.process(arr, 0, len(arr)-1)
}

// arr[L..R] 上，將隨機一個位置的數跟 arr[R] 交換，並以 arr[R] 位置的數做劃分值
func (q *QuickSort) process(arr []int, L, R int) {
	if L >= R {
		return
	}
	q.swap(arr, L+rand.Intn(R-L+1), R)
	equalLeft, equalRight := q.netherlandsFlag(arr, L, R)
	q.process(arr, L, equalLeft-1)
	q.process(arr, equalRight+1, R)
}

func (q *QuickSort) swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// 荷蘭國旗問題
// <arr[R] 小於區域放左邊， ==arr[R] 等於區域放中間， >arr[R] 大於區域放右邊
// 返回 ==arr[R] 左邊和右邊的索引
func (q *QuickSort) netherlandsFlag(arr []int, L, R int) (equalLeft, equalRight int) {
	if L > R {
		return -1, -1
	}
	if L == R {
		return L, R
	}
	less := L - 1
	more := R
	index := L
	for index < more {
		if arr[index] == arr[R] {
			index++
		} else if arr[index] < arr[R] {
			less++
			q.swap(arr, index, less)
			index++
		} else {
			more--
			q.swap(arr, index, more)
		}
	}
	q.swap(arr, more, R)
	return less + 1, more
}
