package main

import "fmt"

func insertSort1(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	N := len(arr)
	for end := 1; end < N; end++ {
		curNumberIndex := end
		for curNumberIndex-1 >= 0 && arr[curNumberIndex-1] > arr[curNumberIndex] {
			swap(arr, curNumberIndex-1, curNumberIndex)
			curNumberIndex--
		}
	}

}

func insertSort2(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	N := len(arr)
	for end := 1; end < N; end++ {
		for curIndex := end; curIndex-1 >= 0 && arr[curIndex-1] > arr[curIndex]; curIndex-- {
			swap(arr, curIndex-1, curIndex)
		}
	}

}

func swap(arr []int, i int, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func main() {
	newAarry := []int{2, 4, 1, 5, 8, 5}
	insertSort2(newAarry)
	fmt.Println(newAarry)
}
