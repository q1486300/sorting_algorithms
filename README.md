## 目錄結構

sorting_algorithms/<br/>
├── binary_search &nbsp;&nbsp;&nbsp;&nbsp;二分搜尋相關內容<br/>
├── bubble_sort &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;氣泡排序相關內容<br/>
├── comparator &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;對數器，產生隨機樣本提供測試使用<br/>
├── heap_sort &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;堆排序相關內容<br/>
├── insertion_sort &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;插入排序相關內容<br/>
├── merge_sort &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;合併排序相關內容<br/>
├── quick_sort &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;快速排序相關內容<br/>
├── radix_sort &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;基數排序相關內容<br/>
├── selection_sort &nbsp;&nbsp;&nbsp;&nbsp;選擇排序相關內容<br/>
└── sorting_strategy &nbsp;封裝排序資料和策略

## 排序演算法比較

| 演算法               | 時間複雜度       | 空間複雜度     | 穩定性 |
|-------------------|-------------|-----------|-----|
| 選擇排序              | $O(N^2)$    | $O(1)$    | 無   |
| 氣泡排序              | $O(N^2)$    | $O(1)$    | 有   |
| 插入排序              | $O(N^2)$    | $O(1)$    | 有   |
| 合併排序              | $O(N*logN)$ | $O(N)$    | 有   |
| 快速排序 <br/>(隨機劃分值) | $O(N*logN)$ | $O(logN)$ | 無   |
| 堆排序               | $O(N*logN)$ | $O(1)$    | 無   |

> 穩定性: 相同值的物件之間，如果不因為排序而改變相對次序，那這個排序就有穩定性