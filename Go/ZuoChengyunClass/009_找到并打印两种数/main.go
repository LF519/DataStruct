/*
一个数组中有两种数出现了奇数次, 其他数都出现了偶数次, 怎么找到并打印这种数
*/
package main

import "fmt"

func printOddNum(arr []int) {
	eor := 0
	for i := 0; i < len(arr); i++ {
		eor ^= arr[i]
	}
	// 设a和b是上面的两种数
	// eor != 0
	// 取出来eor二进制形式中最右侧的1, 提取出来, 如
	// eor: 		0011010110111000
	// rightOne: 	0000000000001000
	rightOne := eor & -eor

	eor1 := 0 // eor'

	// eor的结果是a^b, a必然不等于b, 则他们在二进制位上必然会存在至少某一个位置上不等
	// 取eor中最右侧值为1的数, 则a跟b在此位数必然不等(a^b在此位的值为1), 此数设为变量rightOne
	// 按照rightOne, 将所有数分成两类, 一类是该位数为0的, 一类是该位数为1的
	for i := 0; i < len(arr); i++ {
		// 将其中一类(如该位数为0)的数全部异或, 得到其中一个数, 值为eor1
		if arr[i]&rightOne == 0 {
			eor1 ^= arr[i]
		}
	}
	// a
	fmt.Println(eor1)
	// eor ^ eor1 = a ^ b ^ a = b
	fmt.Println(eor ^ eor1)

}

func main() {
	arr := []int{6, 6, 6, 10, 4, 4, 12, 12, 12, 12, 3, 3}
	printOddNum(arr)
}
