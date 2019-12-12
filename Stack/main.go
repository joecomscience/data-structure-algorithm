package main

import "fmt"

type Stack struct {
	Items []int
}

func (s *Stack) Create() {
	if s.Items == nil {
		s.Items = make([]int, 0)
	}
}

func (s *Stack) IsEmpty() {
	if len(s.Items) == 0 {
		fmt.Println("Stack is empty!")
		return
	} else {
		fmt.Println("Stack in not empty")
	}
}

func (s *Stack) Push(v int) {
	s.Items = append(s.Items, v)
	fmt.Println(s.Items)
}

func (s *Stack) Pop() {
	s.Items = s.Items[:len(s.Items)-1]
	fmt.Println(s.Items)
}

func (s *Stack) PopAll() {
	s.Items = make([]int, 0)
	fmt.Println(s.Items)
}

func (s *Stack) Peek() {
	fmt.Println(s.Items[len(s.Items)-1])
}

func main() {
	s := Stack{}
	s.Create()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Peek()
	s.Pop()
	s.IsEmpty()
	s.PopAll()

}
