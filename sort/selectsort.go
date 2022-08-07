package sort

import "fmt"

func SelectSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	for i := 0; i < length-1; i++ {
		min := i // 标记为最小值
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		if min != i { // 找出了最小值
			arr[i], arr[min] = arr[min], arr[i]
		}
		fmt.Println(arr)
	}
	return arr
}
