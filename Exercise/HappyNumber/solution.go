package main

import "fmt"

func main() {
	x := isHappyNumber(19)
	fmt.Println(x)
}

func isHappyNumber(n int) bool {
	slow, fast := n, n
	for {
		slow = numSquareSum(slow)
		fast = numSquareSum(numSquareSum(fast))
		if slow == fast {
			break
		}
	}
	return slow == 1
}

func numSquareSum(n int) int {
	sum := 0
	for n != 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}

	return sum
}