package main

import "fmt"

// 不用额外变量交换两个变量值
func swap(a int, b int) {
	fmt.Println(a, b)
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a, b)
}

func main() {
	swap(1, 2)
}
