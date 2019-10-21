package main

import "fmt"

type Node struct {
	Value string
	Next  *Node
}

type List struct {
	Head *Node
}

func (l *List) Print() {
	if l.Head == nil {
		fmt.Println("List is empty")
		return
	}

	c := l.Head
	for c != nil {
		fmt.Printf("%s\t", c.Value)
		c = c.Next
	}
	fmt.Println()
}

func (l *List) Append(v string) {
	n := &Node{
		Value: v,
	}
	if l.Head == nil {
		l.Head = n
		return
	}

	c := l.Head
	for c.Next != nil {
		c = c.Next
	}
	c.Next = n
	l.Print()
}

func (l *List) delete(v string) {
	c := l.Head
	if l.Head.Value == v {
		l.Head = c.Next
		l.Print()
		return
	}
	for c != nil {
		if c.Next.Value == v {
			c.Next = c.Next.Next
			break
		} else {
			c = c.Next
		}
	}
	l.Print()
}

func (l *List) reverse() {
	c := l.Head
	var n, p *Node
	for c != nil {
		n = c.Next
		c.Next = p
		p = c
		c = n
	}
	l.Head = p
	l.Print()
}

func (l *List) search(v string) bool {
	h := l.Head
	for h != nil {
		if h.Next.Value == v {
			return true
		} else {
			h = h.Next
		}
	}
	return false

}

func main() {
	n := List{}
	n.Append("joe")
	n.Append("walker")
	n.Append("test1")
	n.Append("test2")
	n.delete("joe")
	n.reverse()
	n.delete("test2")
}
