package binary_search

func Exist(sortedArr []int, num int) bool {
	if sortedArr == nil || len(sortedArr) == 0 {
		return false
	}
	L := 0
	R := len(sortedArr) - 1
	mid := 0
	// L..R
	for L < R { // L..R 至少兩個數時
		mid = L + ((R - L) >> 1)
		if sortedArr[mid] == num {
			return true
		} else if sortedArr[mid] > num {
			R = mid - 1
		} else {
			L = mid + 1
		}
	}
	return sortedArr[L] == num
}

// 在 arr 上，找滿足 >= value 的最左位置
func NearestLeftIndex(arr []int, value int) int {
	if arr == nil {
		return -1
	}
	L := 0
	R := len(arr) - 1
	index := -1  // 紀錄最左的索引
	for L <= R { // 至少一個數的時候
		mid := L + ((R - L) >> 1)
		if arr[mid] >= value {
			index = mid
			R = mid - 1
		} else {
			L = mid + 1
		}
	}
	return index
}

// 在 arr 上，找滿足 <= value 的最右位置
func NearestRightIndex(arr []int, value int) int {
	if arr == nil {
		return -1
	}
	L := 0
	R := len(arr) - 1
	index := -1  //紀錄最右的索引
	for L <= R { // 至少一個數的時候
		mid := L + ((R - L) >> 1)
		if arr[mid] <= value {
			index = mid
			L = mid + 1
		} else {
			R = mid - 1
		}
	}
	return index
}

// 返回陣列中的一個區域最小值
func GetLessIndex(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 || arr[0] < arr[1] {
		return 0
	}
	if arr[len(arr)-1] < arr[len(arr)-2] {
		return len(arr) - 1
	}
	left := 1
	right := len(arr) - 2
	mid := 0
	for left < right {
		mid = left + ((right - left) >> 1)
		if arr[mid] > arr[mid-1] {
			right = mid - 1
		} else if arr[mid] > arr[mid+1] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}
