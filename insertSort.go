package main

import "fmt"

// 前N个子序列依次有序
// 1， 2， 3...
func insertSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	for i := 1; i < len(arr); i++ {
		// 当前值与前子序列从后向前依次比较
		// 遇到大的交换，遇到小的就退出当前循环
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				swap(arr, j, j-1)
			} else {
				continue
			}
		}
		fmt.Println(arr)
	}
}

func swap(arr []int, i int, j int) {
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]
}

func main() {
	arr := []int{3, 8, 1, 10, 233, 78, 6, 1}

	insertSort(arr)
}
