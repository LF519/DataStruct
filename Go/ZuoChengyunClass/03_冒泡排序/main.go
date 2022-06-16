package main

import "fmt"

func bubbleSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	N := len(arr)
	for i := N - 1; i >= 0; i-- {
		for j := 0; j < i-1; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
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
	bubbleSort(newAarry)
	fmt.Println(newAarry)
}
