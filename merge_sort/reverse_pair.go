package merge_sort

// 使用到合併排序流程的題目
// 在一個陣列中，左邊的數如果比右邊的數大，則這兩個數組成唷個逆序對，返回所有逆序對數量
func ReversePairNumber(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	return processReverse(arr, 0, len(arr)-1)
}

// arr[L..R] 既要排好序，也要求逆序對數量返回
// 所有 merge 時，產生的逆序對數量累加，並返回
// 左 排序 merge 並產生逆序對數量
// 右 排序 merge 並產生逆序對數量
// merge
func processReverse(arr []int, L, R int) int {
	if L == R {
		return 0
	}
	mid := L + ((R - L) >> 1)
	return processReverse(arr, L, mid) + processReverse(arr, mid+1, R) + mergeReverse(arr, L, mid, R)
}

func mergeReverse(arr []int, L, M, R int) int {
	help := make([]int, R-L+1)
	i := len(help) - 1
	p1 := M
	p2 := R
	res := 0
	for p1 >= L && p2 > M {
		if arr[p1] > arr[p2] {
			res += p2 - M
			help[i] = arr[p1]
			p1--
		} else {
			help[i] = arr[p2]
			p2--
		}
		i--
	}
	for p1 >= L {
		help[i] = arr[p1]
		p1--
		i--
	}
	for p2 > M {
		help[i] = arr[p2]
		p2--
		i--
	}
	for i := 0; i < len(help); i++ {
		arr[L+i] = help[i]
	}
	return res
}
