package sortAlg

import "fmt"

func printOneOddTimesNum(arr []int) {
	eor := 0

	// {1, 2, 3, 1, 1, 3, 2}
	// (a^b)^c = a^(b^c) 异或交换律、
	// 偶数个异或完为0，奇数个异或完只剩自己，一起就只剩奇数个自己
	// 1 ^ 1 = 0, 0 ^ 0 = 0, 0 ^ 1 = 1
	for _, v := range arr {
		eor ^= v
	}

	fmt.Println(eor)
}

func printTwoOddTimesNum(arr []int) {
	eor := 0

	// eor = a ^ b 且 != 0
	for _, v := range arr {
		eor ^= v
	}

	// eor必有最右一位为1的位置，因为eor = a^b且a,b不相等。
	// 通过这个位置上的值为0，1来将数组分组，且a, b必分在不同组。
	// 此位置按如下方法选取最右为1的位置。
	// 1001 0100
	// 0110 1011 + 1 = 0110 1100
	// 1001 0100 & 0110 1100 = 0000 0100
	rightOne := eor & (^eor + 1)

	eor1 := 0
	for _, v := range arr {
		// 只选取最右为1位置上所有为1的数，a或b之一必在其中
		// eor1最终必为a或b之一
		// 另一个数就是eor ^ eor1
		// == 0 或者 == rightOne，不能 == 1。 rightOne = 0000 0100
		if v&rightOne == 0 {
			eor1 ^= v
		}
	}

	fmt.Printf("%d  %d", eor1, eor^eor1)
}

func test() {
	arr1 := []int{1, 2, 5, 7, 5, 5, 7, 2, 1, 1, 1}
	arr2 := []int{1, 2, 5, 1, 7, 5, 5, 7, 2, 1, 1, 1}

	// 找出数组中唯一出现奇数次数的数
	printOneOddTimesNum(arr1)

	// 找出数组中两个出现奇数次数的数
	printTwoOddTimesNum(arr2)
}
