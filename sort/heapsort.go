package sort

func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastlength := length - i
		HeapSortMax(arr, lastlength)
		arr[0], arr[lastlength-1] = arr[lastlength-1], arr[0]
	}
	return arr
}
func HeapSortMax(arr []int, length int) []int {

	//length := len(arr)
	depth := length/2 - 1 // 层级数
	for i := depth; i >= 0; i-- {
		leftchild := 2*i + 1
		rigthchild := 2*i + 2
		topmax := i // 标记默认最大为顶端元素
		if leftchild < length && arr[leftchild] > arr[topmax] {
			topmax = leftchild
		}
		if rigthchild < length && arr[rigthchild] > arr[topmax] {
			topmax = rigthchild
		}
		if topmax != i { // 交换值
			arr[topmax], arr[i] = arr[i], arr[topmax]
		}

	}
	return arr
}
