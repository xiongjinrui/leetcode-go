package sortAlg

func bubbleSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	// 除去沉底的最大数，继续遍历
	for i := len(arr) - 1; i > 0; i-- {
		//fmt.Println("----------------------------")
		// 将最大数通过一次遍历沉底
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				swap1(arr, j, j+1)
			}
			//fmt.Println(arr)
		}
	}
}

// 通过异或运算来交换。
// 省空间，且运算速度快。
// 但i, j不能相等。 同数异或将变为0。
// a = a ^ b
// b = (a ^ b) ^ b = a
// a = (a ^ b) ^ a = b
func swap1(arr []int, i int, j int) {
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]
}
