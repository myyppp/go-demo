package sort

// 选择排序
// 思想：在未排序序列中找到最小的元素，存放在有序序列的起始位置；以此类推
// 平均时间复杂度：O(N^2)
// 最佳时间复杂度：O(N^2)
// 最差时间复杂度：O(N^2)
// 空间复杂度：O(1)
// 排序方式：In-place
// 稳定性：不稳定
// 比较次数与关键字的初始状态无关
func Select(arr []int) []int {
	length := len(arr)

	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

// 改进
// 1. 二元选择排序
// 每趟循环确定两个元素，最大和最小
// 对n个数据进行排序，只需要n/2趟即可

// 2.堆排序
