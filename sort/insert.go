package sort

// 将待排序的记录按其大小插入到前面已排好序的子序列中
// 从后向前比较再移动
// 空间复杂度 O( 1 )
// 时间复杂度 与待排序的初始状态有关
// 有序 O( n )
// 逆序 O( n^2 )
// 稳定
// 顺序存储 链式存储
func Insert(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		preIndex := i - 1
		cur := arr[i]
		for ; preIndex >= 0 && arr[preIndex] > cur; preIndex-- {
			arr[preIndex+1] = arr[preIndex]
		}
		arr[preIndex+1] = cur
	}
	return arr
}

// 折半插入排序
// 顺序存储改进
func BinaryInsert(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		low, hight, mid := 0, i-1, 0
		cur := arr[i]

		for low <= hight {
			mid = low + (hight-low)/2
			if arr[mid] > cur {
				hight = mid - 1
			} else {
				low = mid + 1
			}
		}

		for preIndex := i - 1; preIndex > low; preIndex-- {
			arr[preIndex+1] = arr[preIndex]
		}
		arr[low] = cur
	}
	return arr
}
