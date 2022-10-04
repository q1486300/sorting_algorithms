package merge_sort

// 使用到合併排序流程的題目
// 在一個陣列中，每一個數左邊比目前數小的數累加起來，叫做這個陣列的小和
func SmallSum(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	return process(arr, 0, len(arr)-1)
}

// arr[L..R] 既要排好序，也要求小和返回
// 所有 merge 時，產生的小和累加
// 左 排序 merge
// 右 排序 merge
// merge
func process(arr []int, L, R int) int {
	if L == R {
		return 0
	}
	mid := L + ((R - L) >> 1)
	return process(arr, L, mid) + process(arr, mid+1, R) + merge(arr, L, mid, R)
}

func merge(arr []int, L, M, R int) int {
	help := make([]int, R-L+1)
	i := 0
	p1 := L
	p2 := M + 1
	res := 0
	for p1 <= M && p2 <= R {
		if arr[p1] < arr[p2] {
			res += (R - p2 + 1) * arr[p1]
			help[i] = arr[p1]
			p1++
		} else {
			help[i] = arr[p2]
			p2++
		}
		i++
	}
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
	return res
}
