package main

import "fmt"

var output [][]int

func main() {
	getPermutations([]int{1, 2, 3, 4})
	for _, v := range output {
		fmt.Println(v)
	}
	fmt.Println(len(output))
}

func getPermutations(arr []int) {
	generate(len(arr), arr)
}

func generate(n int, heapArr []int) {
	if n == 1 {
		output = append(output, heapArr)
		return
	}

	generate(n-1, heapArr)
	for i := 0; i < n-1; i++ {
		if n%2 == 0 {
			swapIndex(heapArr, i, n-1)
		} else {
			swapIndex(heapArr, 0, n-1)
		}
		generate(n-1, heapArr)
	}
}

func swapIndex(arr []int, a int, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}
