package comparator

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

// 對數器
func SortingComparator(arr []int) {
	sort.Ints(arr)
}

func SearchingExistComparator(sortedArr []int, num int) bool {
	for _, cur := range sortedArr {
		if cur == num {
			return true
		}
	}
	return false
}

func SearchingNearestLeftIndexComparator(arr []int, value int) int {
	for i, cur := range arr {
		if cur >= value {
			return i
		}
	}
	return -1
}

func SearchingNearestRightIndexComparator(arr []int, value int) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] <= value {
			return i
		}
	}
	return -1
}

func LessIndexComparator(arr []int, index int) bool {
	if len(arr) <= 1 {
		return true
	}
	if index == 0 {
		return arr[index] < arr[index+1]
	}
	if index == len(arr)-1 {
		return arr[index] < arr[index-1]
	}
	return arr[index] < arr[index-1] && arr[index] < arr[index+1]
}

func SmallSumComparator(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	res := 0
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				res += arr[j]
			}
		}
	}
	return res
}

func ReversePairComparator(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	res := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				res++
			}
		}
	}
	return res
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 隨機樣本產生器
func GenerateRandomArray(maxSize, maxValue int) []int {
	// rand.Intn(maxSize)		[0, N)
	// rand.Intn(maxSize+1)		[0, N]
	arr := make([]int, rand.Intn(maxSize)+1)
	for i := 0; i < len(arr); i++ {
		// [-?, +?]
		arr[i] = (rand.Intn(maxValue) + 1) - (rand.Intn(maxValue) + 1)
	}
	return arr
}

// 產生只有正數的陣列
func GenerateNoNegativeRandomArray(maxSize, maxValue int) []int {
	// rand.Intn(maxSize)		[0, N)
	// rand.Intn(maxSize+1)		[0, N]
	arr := make([]int, rand.Intn(maxSize)+1)
	for i := 0; i < len(arr); i++ {
		// [-?, +?]
		arr[i] = rand.Intn(maxValue) + 1
	}
	return arr
}

// 產生相鄰不相等的陣列
func GenerateNotEqualRandomArray(maxSize, maxValue int) []int {
	// rand.Intn(maxSize)		[0, N)
	// rand.Intn(maxSize+1)		[0, N]
	arr := make([]int, rand.Intn(maxSize)+1)
	arr[0] = (rand.Intn(maxValue) + 1) - (rand.Intn(maxValue) + 1)
	for i := 1; i < len(arr); i++ {
		// [-?, +?]
		for {
			arr[i] = (rand.Intn(maxValue) + 1) - (rand.Intn(maxValue) + 1)
			if arr[i] != arr[i-1] {
				break
			}
		}
	}
	return arr
}

func RandomArrayNoMoveMoreK(maxSize, maxValue, K int) []int {
	arr := make([]int, rand.Intn(maxSize)+1)
	for i := 0; i < len(arr); i++ {
		arr[i] = (rand.Intn(maxValue) + 1) - (rand.Intn(maxValue) + 1)
	}
	// 先排序
	sort.Ints(arr)
	// 然後開始隨意交換，但保證每個數的距離不超過 K
	// isSwap[i] == true, 表示 i 位置已經參與過交換
	// isSwap[i] == false, 表示 i 位置沒有參與過交換
	isSwap := make([]bool, len(arr))
	for i := 0; i < len(arr); i++ {
		j := int(math.Min(float64(i+(rand.Intn(K)+1)), float64(len(arr)-1)))
		if !isSwap[i] && !isSwap[j] {
			isSwap[i] = true
			isSwap[j] = true
			tmp := arr[i]
			arr[i] = arr[j]
			arr[j] = tmp
		}
	}
	return arr
}

func CopyArray(arr []int) []int {
	if arr == nil {
		return nil
	}
	res := make([]int, len(arr))
	copy(res, arr)
	return res
}

func PrintArray(arr []int) {
	if arr == nil {
		return
	}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}
