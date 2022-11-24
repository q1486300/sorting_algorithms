package radix_sort

import (
	"math"
	"sorting_algorithms/sorting_strategy"
)

// only for no-negative value
// 時間複雜度 O(N)
// 空間複雜度 O(M)，輔助陣列的長度和(0~9)基數的桶
// 應用範圍有限，需要樣本的數據狀況滿足桶的劃分
type RadixSort struct {
}

var _ sorting_strategy.ISortingAlgorithm = &RadixSort{}

func (r *RadixSort) Sort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	r.process(arr, 0, len(arr)-1, r.maxBits(arr))
}

// arr[L..R] 排序，digit 表示最大值總共有幾位十進制數
func (r *RadixSort) process(arr []int, L, R, digit int) {
	const radix = 10
	// 有多少個數就準備多少個輔助空間
	help := make([]int, R-L+1)
	for d := 1; d <= digit; d++ { // 有多少位數就進出桶幾次
		// 10 個空間
		// count[0] 目前位(d位)是0的數字有幾個
		// count[1] 目前位(d位)是(0和1)的數字有幾個
		// count[2] 目前位(d位)是(0、1和2)的數字有幾個
		// count[i] 目前位(d位)是(0~i)的數字有幾個
		count := make([]int, radix)
		for i := L; i <= R; i++ {
			digitValue := r.getDigit(arr[i], d)
			count[digitValue]++
		}
		for i := 1; i < radix; i++ {
			count[i] = count[i] + count[i-1]
		}
		for i := R; i >= L; i-- {
			digitValue := r.getDigit(arr[i], d)
			help[count[digitValue]-1] = arr[i]
			count[digitValue]--
		}

		for i, value := range help {
			arr[L+i] = value
		}
	}
}

// 返回陣列中的最大值總共有幾位十進制數
func (r *RadixSort) maxBits(arr []int) int {
	max := math.MinInt
	for _, cur := range arr {
		max = getMax(max, cur)
	}
	res := 0
	for max != 0 {
		res++
		max /= 10
	}
	return res
}

func (r *RadixSort) getDigit(value int, digit int) int {
	return (value / int(math.Pow10(digit-1))) % 10
}

func getMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
