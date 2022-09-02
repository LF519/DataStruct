/*
一个数组中有一种数出现了奇数次, 其他数都出现了偶数次, 怎么找到并打印这种数
*/
package main

import "fmt"

func printOddNum(arr []int) {
	//方法很简单，就是用异或的性质：N^0=N,N^N=0
	// 说明1次【奇数次】异或完还是自己，2次【偶数次】异或完为0
	eor := 0
	for i := 0; i < len(arr); i++ {
		eor ^= arr[i]
	}
	fmt.Println(eor)
}

func main() {
	arr := []int{0, 0, 1, 2, 2, 2, 2, 3, 3}
	printOddNum(arr)
}
