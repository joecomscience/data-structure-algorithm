package main

import "fmt"

func main() {
	head := &Node{value: 1}
	head.push(20)
	head.push(4)
	head.push(15)
	head.push(10)

	head.next.next.next.next = head
	head.detectLoop()
}

type Node struct {
	value int
	next  *Node
}

func (n *Node) push(v int)  {
	nd := &Node{value: v}
	nd.next = n
	n = nd
}

func (n *Node) detectLoop() int {
	slow_p, fast_p := n, n

	for slow_p != nil && fast_p != nil && fast_p.next != nil {
		slow_p = slow_p.next
		fast_p = fast_p.next.next
		if slow_p == fast_p {
			fmt.Println("found")
			return 1
		}
	}

	return 0
}