package main

import "fmt"

type Node struct {
	Value int
	Node  *Node
}
type LinkList struct {
	Head *Node
}

func (l *LinkList) Print() {
	head := l.Head

	if head == nil {
		fmt.Println("List is empty!")
		return
	}

	for head != nil {
		fmt.Println(head.Value)
		head = head.Node
	}
}

func (l *LinkList) Append(v int) {
	head := l.Head
	newNode := &Node{Value: v}

	if head == nil {
		l.Head = newNode
		return
	}

	/* If head not empty set head to next node*/
	for head.Node != nil {
		head = head.Node
	}

	head.Node = newNode
}

func (l *LinkList) Reverse() {
	cur := l.Head

	var next, prev *Node
	for cur != nil {
		next = cur.Node
		cur.Node = prev
		prev = cur
		cur = next
	}

	l.Head = prev
}

func (l *LinkList) AppendAfterHead(v int) {
	head := l.Head
	newNode := &Node{Value: v}

	if head != nil {
		prev := head.Node
		l.Head.Node = newNode
		l.Head.Node.Node = prev
	}
}

func (l *LinkList) ChangeHeadValue(v int) {
	head := l.Head
	newNode := &Node{Value: v}

	if head != nil {
		prev := head.Node
		l.Head = newNode
		l.Head.Node = prev
	}
}

func (l *LinkList) Delete(v int) {
	head := l.Head
	if head.Value == v {
		l.Head = head.Node
		return
	}

	for head.Node != nil {
		if head.Node.Value == v {
			head.Node = head.Node.Node
			break
		} else {
			head = head.Node
		}
	}
}

func (l *LinkList) Search(v int) {
	head := l.Head

	if head.Value == v {
		fmt.Printf("found head value: %v", v)
		return
	}

	for head.Node != nil {
		if head.Node.Value == v {
			fmt.Printf("found value: %v", v)
			break
		} else {
			head = head.Node
		}
	}
}
func main() {
	link := LinkList{}
	link.Append(1)
	link.Append(2)
	link.Append(3)
	link.Append(4)
	//link.AppendAfterHead(5)
	//link.ChangeHeadValue(6)
	//link.Delete(4)
	link.Print()
	link.Reverse()
	link.Print()
	//link.Search(1)
}
