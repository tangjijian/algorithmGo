package sort

func InsertSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {

		//for i := 1; i < length; i++ { // 跳过第一个
		//	backup := arr[i] // 备份插入的数据
		//	j := i - 1       //
		//	for j >= 0 && backup < arr[j] {
		//		arr[j+1] = arr[j]
		//		j--
		//	}
		//	arr[j+1] = backup
		//}
		//3, 2, 4, 6, 8, 1, 5, 9
		//3, 2, 4, 6, 8, 1, 5, 9
		for i := 1; i < length; i++ {
			backup := arr[i]
			j := i - 1
			for j >= 0 {
				if backup < arr[j] {
					arr[j+1] = arr[j]
					j--
				} else {
					break
				}

			}
			arr[j+1] = backup
		}
	}
	return arr
}
