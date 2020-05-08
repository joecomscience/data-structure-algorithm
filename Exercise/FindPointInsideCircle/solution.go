package main

import (
	"fmt"
	"strconv"
)

func main() {
	inner := 2
	outer := 4
	//pointX := []int{0, 1, 2, -2, 3}
	pointX := []int{4, 0, 1, -2}
	//pointY := []int{0, 1, 4, 1, 0}
	pointY := []int{-4, 4, 3, 0}

	fmt.Printf("result: %v\n", findPoint(inner, outer, pointX, pointY))
}

func findPoint(i int, o int, pointX []int, pointY []int) [][]int {
	inner, outer := getPointInsideCircle(i, o, pointX, pointY)
	a := make(map[string]bool)
	for _, v := range outer {
		str := strconv.Itoa(v[0]) + strconv.Itoa(v[1])
		a[str] = true
	}

	r := [][]int{}
	for _, v := range inner {
		str := strconv.Itoa(v[0]) + strconv.Itoa(v[1])
		if _, ok := a[str]; !ok {
			r = append(r, v)
		}
	}

	return r
}

func getPointInsideCircle(i int, o int, pointX []int, pointY []int) ([][]int, [][]int) {
	var inner [][]int
	var outer [][]int
	pointList := getPointList(pointX, pointY)
	for _, v := range pointList {
		x := v[0]
		y := v[1]
		_outer := (x*x)+(y*y) < o*o
		_inner := (x*x)+(y*y) <= i*i

		if _outer {
			inner = append(inner, v)
		}
		if _inner {
			outer = append(outer, v)
		}
	}

	fmt.Println(inner)
	fmt.Println(outer)
	return inner, outer
}

func getPointList(x []int, y []int) [][]int {
	var r [][]int
	for i, _ := range x {
		r = append(r, []int{x[i], y[i]})
	}

	return r
}
