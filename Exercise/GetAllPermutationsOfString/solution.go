package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "ABCD"
	permute(str, 0, len(str)-1)
}

/**
 * permutation function
 * @param str string to calculate permutation for
 * @param l starting index
 * @param r end index
 */
func permute(str string, l int, r int) {
	if l == r {
		fmt.Println(str)
	} else {
		for i := l; i <= r; i++ {
			str = swap(str, l, i)
			permute(str, l+1, r)
			str = swap(str, l, i)
		}
	}
}

func swap(s string, a int, b int) string {
	arr := strings.Split(s, "")
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp

	return strings.Join(arr, "")
}
