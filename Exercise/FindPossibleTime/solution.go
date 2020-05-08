package main

import (
	"fmt"
	"strings"
)

func main() {
	txt := "1222"
	u := make(map[string]int)
	permutations(txt, 0, len(txt)-1, u)
	t := getTime()
	r := findTime(u, t)
	fmt.Println(len(r))
}

func permutations(str string, start int, end int, m map[string]int) {
	if start == end {
		strArr := strings.Split(str, "")
		key := strArr[0] + strArr[1] + ":" + strArr[2] + strArr[3]
		m[key] = 1
	} else {
		for i := start; i <= end; i++ {
			str = swap(str, start, i)
			permutations(str, start+1, end, m)
			str = swap(str, start, i)
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

func getTime() map[string]int {
	t := make(map[string]int)
	for i := 0; i < 24; i++ {
		for j := 0; j <= 60; j++ {
			time := fmt.Sprintf("%02d:%02d", i, j)
			t[time] = 1
		}
	}

	return t
}

func findTime(u map[string]int, t map[string]int) []string {
	var s []string
	r := make(map[string]int)

	for uk, _ := range u {
		for _, _ = range t {
			if t[uk] == 1 && r[uk] == 0 {
				r[uk] = 1
			}
		}
	}

	for k, _ := range r {
		s = append(s, k)
	}
	fmt.Println(s)
	return s
}
