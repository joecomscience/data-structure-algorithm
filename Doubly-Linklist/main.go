package main

import "fmt"

type Node struct {
	Value      int
	Next, Prev *Node
}

type DoublyLinkList struct {
	Head *Node
}

func (dl *DoublyLinkList) Print() {
	h := dl.Head
	if h == nil {
		fmt.Println("List is empty")
		return
	}

	for h != nil {
		fmt.Printf("value: %d, next: %v, prev: %v \n", h.Value, h.Next, h.Prev)
		h = h.Next
	}
}

func (dl *DoublyLinkList) Append(v int) {
	h := dl.Head
	n := &Node{Value: v}

	if h == nil {
		dl.Head = n
		return
	}

	for h.Next != nil {
		h = h.Next
	}

	n.Prev = h
	h.Next = n
	h = n
}

func (dl *DoublyLinkList) Delete(v int) {
	h := dl.Head

	if h.Value == v {
		h = h.Next
	}

	for h.Next != nil {
		if h.Next.Value == v {
			h.Next = h.Next.Next
		} else {
			h = h.Next
		}
	}
}

func main() {
	dl := DoublyLinkList{}
	dl.Append(1)
	dl.Append(2)
	dl.Append(3)
	dl.Append(4)
	dl.Print()
	fmt.Println("-----")
	dl.Delete(3)
	dl.Print()
}
