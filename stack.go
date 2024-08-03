package stack

type node[T any] struct {
	next *node[T]
	data T
}

type Stack[T any] struct {
	head *node[T]
}

func (s *Stack[T]) Push(data T) {
	t := &node[T]{
		data: data,
		next: nil,
	}

	if s.head == nil {
		s.head = t
	} else {
		t.next = s.head
		s.head = t
	}
}

func (s *Stack[T]) Pop() T {
	data := s.head.data
	s.head = s.head.next

	return data
}

func (s *Stack[T]) Peek() T {
	return s.head.data
}

func (s *Stack[T]) IsEmpty() bool {
	return s.head == nil
}
