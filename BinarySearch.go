package main

import "fmt"

func main() {
	nums := []int{2, 3, 4, 10, 40}
	x := binarySearch(nums, 10)
	fmt.Println(x)

	y := bsRecursive(nums, 0, len(nums)-1, 10)
	fmt.Println(y)
}

func binarySearch(arr []int, v int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == v {
			return v
		} else if arr[mid] < v {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func bsRecursive(arr []int, l int, r int, v int) int {
	mid := (l + r) / 2

	if l <= r {
		if arr[mid] == v {
			return v
		} else if arr[mid] < v {
			return bsRecursive(arr, mid+1, r, v)
		} else {
			return bsRecursive(arr, l, mid-1, v)
		}
	}
	return -1
}
