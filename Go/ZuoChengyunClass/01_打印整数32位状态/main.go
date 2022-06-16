package main

import "fmt"

func print(num int32) {
	for i := 31; i >= 0; i-- {
		if num&(1<<i) == 0 {
			fmt.Print("0")
		} else {
			fmt.Print("1")
		}
	}
	fmt.Println()
}

func main() {
	var num int32 = 1
	print(num >> 2)
	// print(num << 1)
	// print(num << 2)
	// print(num << 8)

}
