package algo

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	// 递归，左段有序，右段有序，合并成新的有序数组
	// 时间复杂度：master公式 T(N) = aT(N/b) + O(N^d)
	// 其中 a >= 1 and b > 1 是常量，其表示的意义是n表示问题的规模，
	// a表示递归的次数也就是生成的子问题数，b表示每次递归是原来的1/b之一个规模，
	// f（n）表示分解和合并所要花费的时间之和。
	// 当d<logb a时，时间复杂度为O(n^(logb a))
	// 当d=logb a时，时间复杂度为O((n^d)*logn)
	// 当d>logb a时，时间复杂度为O(n^d)

	// 2T(N/2)+O(N) = O(NlgN)
	left := MergeSort(arr[:len(arr)/2])
	right := MergeSort(arr[(1 + len(arr)/2):])
	return merge(left, right) // O(N)
}

// 将两段有序数组插入到一个临时数组中去，生成一个新的有序数组
func merge(left []int, right []int) []int {
	var temp []int
	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			temp = append(temp, left[i])
			i++
		} else {
			temp = append(temp, right[j])
			j++
		}
	}

	for ; i < len(left); i++ {
		temp = append(temp, left[i])
	}

	for ; j < len(right); j++ {
		temp = append(temp, right[j])
	}

	return temp
}
