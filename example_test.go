package stack

import "fmt"

func ExampleStack_Push() {
	s := Stack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
}

func ExampleStack_Pop() {
	s := Stack[int]{}
	s.Push(1)
	s.Push(2)

	fmt.Println(s.Pop())
	// Output: 2

	fmt.Println(s.Pop())
	// Output: 1
}

func ExampleStack_Peek() {
	s := Stack[int]{}
	s.Push(1)
	s.Push(2)

	fmt.Println(s.Peek())
	// Output: 2
}

func ExampleStack_IsEmpty() {
	s := Stack[int]{}
	s.Push(1)
	s.Push(2)

	fmt.Println(s.IsEmpty())
	// Output: false

	s.Pop()
	s.Pop()

	fmt.Println(s.IsEmpty())
	// Output: true
}
