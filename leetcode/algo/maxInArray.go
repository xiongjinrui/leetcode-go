package algo

func GetMaxNumberOfArray(arr []int) int {
	return process(arr, 0, len(arr)-1)
}

// 递归查找L到R之间的最大值
func process(arr []int, L int, R int) int {
	if L == R {
		return arr[L]
	}

	// 防止（L + R)/2发生溢出
	// 通过>>右移一位，相当于/2，但位运算要快
	mid := L + (R-L)>>1
	lMax := process(arr, L, mid)
	rMax := process(arr, mid+1, R)

	if rMax > lMax {
		return rMax
	} else {
		return lMax
	}
}
