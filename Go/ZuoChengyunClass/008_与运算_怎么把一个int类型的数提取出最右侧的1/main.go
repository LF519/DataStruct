/*
怎么把一个int类型的数, 提取出最右侧的1(二进制)来
*/
package main

import "fmt"

func main() {
	// a 与 a的取反加一(也就是-a)
	a := 12
	fmt.Printf("%b\n", a)
	fmt.Printf("%b\n", a&(-a))
}
