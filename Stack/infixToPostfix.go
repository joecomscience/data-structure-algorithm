package main

import "fmt"

type stack struct {
	stack []string
}

func (s *stack) push(v string) {
	s.stack = append(s.stack, v)
}

func (s *stack) pop() string {
	if len(s.stack) < 1 {
		return ""
	}
	v := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return v
}

func (s *stack) isEmpty() bool {
	return len(s.stack) == 0
}

func checkOperator(o string) bool {
	switch o {
	case "+":
		return true
	case "-":
		return true
	case "*":
		return true
	case "/":
		return true
	case "%":
		return true
	case "(":
		return true
	default:
		return false
	}
}

func infixToPostfix(infix string) string {
	postfix := ""
	s := stack{}

	for _, ch := range infix {
		sts := string(ch)
		if checkOperator(sts) {
			s.push(sts)
		} else if sts == ")" {
			buffer := s.pop()
			for buffer != "(" {
				postfix += buffer
				buffer = s.pop()
			}
		} else {
			postfix += sts
		}
	}

	for !s.isEmpty() {
		postfix += s.pop()
	}

	fmt.Println(postfix)
	return postfix
}

func main() {
	infix := "a-(b+c*d)/e"
	infixToPostfix(infix)
}
