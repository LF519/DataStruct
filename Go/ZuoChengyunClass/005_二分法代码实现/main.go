package main

// 找到有序数组中某个num是否存在
func exist(sortedArr []int, num int) bool {
	if sortedArr == nil || len(sortedArr) == 0 {
		return false
	}
	L := 0
	R := len(sortedArr) - 1
	mid := L
	for L < R {
		mid = L + ((R - L) >> 1)
		if sortedArr[mid] == num {
			return true
		} else if sortedArr[mid] < num {
			L = mid + 1
		} else {
			R = mid - 1
		}
	}
	// L = R时退出了for循环, 所以最后还要验证一下L本身
	return sortedArr[L] == num
}

func main() {

}
