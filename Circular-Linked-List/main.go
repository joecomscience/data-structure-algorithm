package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type List struct {
	Head *Node
}

func (l *List) Print() {
	h := l.Head
	if h == nil {
		fmt.Println("List is empty!")
		return
	}

	for h.Next != nil {
		if l.Head.Value == h.Next.Value {
			fmt.Printf("value: %d, next: %v\n", h.Value, h.Next)
			break
		}
		fmt.Printf("value: %d, next: %v\n", h.Value, h.Next)
		h = h.Next
	}
}

func (l *List) Append(v int) {
	h := l.Head
	n := &Node{Value: v}
	var prev *Node

	if h == nil {
		l.Head = n
		l.Head.Next = n
		return
	}

	prev = h
	for l.Head.Value != h.Next.Value {
		h = h.Next
	}

	n.Next = prev
	h.Next = n
}

func (l *List) Delete(v int) {
	h := l.Head
	if h == nil {
		fmt.Println("List is empty!")
		return
	}

	if h.Value == v {
		l.Head = h.Next
		return
	}

	for l.Head.Value != h.Next.Value {
		if h.Value == v {
			h = h.Next.Next
		} else  {
			h = h.Next
		}
	}
}

func main() {
	l := List{}
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)
	l.Print()
	fmt.Println("----")
	l.Delete(1)
	l.Print()
}
