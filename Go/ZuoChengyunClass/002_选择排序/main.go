package main

import "fmt"

func selectSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	N := len(arr)
	for i := 0; i < N; i++ {
		minValueIndex := i
		for j := i + 1; j < N; j++ {
			if arr[j] < arr[minValueIndex] {
				minValueIndex = j
			}
		}
		swap(arr, i, minValueIndex)
	}
}

func swap(arr []int, i int, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func main() {
	newAarry := []int{2, 4, 1, 5, 8, 5}
	selectSort(newAarry)
	fmt.Println(newAarry)
}
