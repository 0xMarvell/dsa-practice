package main

import "fmt"

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
	fmt.Println(*s)
}

// func (s *Stack) Pop() (string, bool) {

// }

func main() {
	var stack Stack
	stack.Push("D")
	stack.Push("S")
	stack.Push("A")
}
