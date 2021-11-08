package algo

import "fmt"

func ReversePairs(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}

	return countPairs(arr, 0, len(arr)-1)
}

func countPairs(arr []int, l int, r int) int {
	if l == r {
		return 0
	}

	mid := l + (r-l)>>1

	return countPairs(arr, l, mid) + countPairs(arr, mid+1, r) + countMerge(arr, l, mid, r)
}

// 将两段有序数组插入到一个临时数组中去，生成一个新的有序数组
func countMerge(arr []int, l int, mid int, r int) int {
	var temp []int
	p1 := l
	p2 := mid + 1
	res := 0

	fmt.Println(arr[l:r])
	fmt.Println(arr[l:mid])
	fmt.Println(arr[mid:r])

	for p1 <= mid && p2 <= r {
		if arr[p1] < arr[p2] {
			temp = append(temp, arr[p1])
			p1++
		} else {
			res += (mid + l) - (l + p1)
			temp = append(temp, arr[p2])
			p2++
		}
	}

	for ; p1 <= mid; p1++ {
		temp = append(temp, arr[p1])
	}

	for ; p2 < r; p2++ {
		temp = append(temp, arr[p2])
	}

	for i := 0; i < len(temp); i++ {
		arr[l+i] = temp[i]
	}

	fmt.Println(res)

	return res
}
