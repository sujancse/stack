package stack

type node[T any] struct {
	next *node[T]
	data T
}

type Stack[T any] struct {
	head *node[T]
}

func (s Stack[T]) push(data T) {
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

func (s Stack[T]) pop() T {
	data := s.head.data
	s.head = s.head.next

	return data
}

func (s Stack[T]) peek() T {
	return s.head.data
}

func (s Stack[T]) isEmpty() bool {
	return s.head == nil
}
