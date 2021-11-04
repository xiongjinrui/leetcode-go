package main

import (
	"fmt"
	"leetcodego/leetcode/algo"
	"math/rand"
	"reflect"
	"sort"
	"time"
)

func generateRandomArray(maxSize int, maxValue int) []int {
	arr := make([]int, maxSize)
	for i := 0; i < maxSize-1; i++ {
		arr[i] = rand.Intn(maxValue)
	}

	return arr
}

func main() {
	testTimes := 100
	maxSize := 10
	maxValue := 100
	succeed := true

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < testTimes; i++ {
		arr1 := generateRandomArray(maxSize, maxValue)
		arr2 := arr1

		algo.InsertSort(arr1)
		sort.Ints(arr2)

		fmt.Println("--------------------------------------------")
		fmt.Println(arr1)
		fmt.Println(arr2)

		if !reflect.DeepEqual(arr1, arr2) {
			succeed = false
			break
		}

		fmt.Println(algo.GetMaxNumberOfArray(arr1))
	}

	if succeed {
		fmt.Println("Right!")
	} else {
		fmt.Println("Wrong!!")
	}

}
