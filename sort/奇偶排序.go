package sort

func OddEven(arr []int) []int {

	isSorted := false
	for !isSorted {
		isSorted = true
		for i := 1; i < len(arr)-1; i += 2 { // 奇数排序
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false // 发生了交换
			}
		}
		for i := 0; i < len(arr)-1; i += 2 { // 偶数排序
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false // 发生了交换
			}
		}
	}
	return arr
}
