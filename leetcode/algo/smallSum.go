package algo

func SmallSum(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	return SmallSum(arr[:len(arr)/2]) + SmallSum(arr[(1+len(arr)/2):]) + merge1(arr[:len(arr)/2], arr[(1+len(arr)/2):])
}

// 将两段有序数组插入到一个临时数组中去，生成一个新的有序数组
func merge1(left []int, right []int) int {
	var temp []int
	i := 0
	j := 0
	res := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res += left[i] * (len(right) - j)
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

	return res
}
