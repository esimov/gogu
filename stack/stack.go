package stack

type Stack[T comparable] struct {
	Items []T
}

func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(item T) {
	s.Items = append(s.Items, item)
}

func (s *Stack[T]) Pop() (item T) {
	if s.Size() == 0 {
		return
	}

	item = s.Items[s.Size()-1]
	s.Items = s.Items[:s.Size()-1]

	return
}

func (s *Stack[T]) Peek() (item T) {
	if s.Size() == 0 {
		return
	}
	return s.Items[s.Size()-1]
}

func (s *Stack[T]) Search(item T) bool {
	for i := 0; i < s.Size(); i++ {
		if s.Items[i] == item {
			return true
		}
	}

	return false
}

func (s *Stack[T]) Size() int {
	return len(s.Items)
}
