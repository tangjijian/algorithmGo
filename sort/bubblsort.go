package sort

import "fmt"

func BubbleSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	for i := 0; i < length-1; i++ {
		exchange := false
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				exchange = true
			}

		}
		if !exchange {
			break
		}
		fmt.Println(arr)
	}
	return arr
}
