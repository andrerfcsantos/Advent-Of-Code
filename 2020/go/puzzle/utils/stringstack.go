package utils

type StringStack struct {
	elems []string
}

func NewStringStack() StringStack {
	return StringStack{}
}

func (s *StringStack) Push(elem string) {
	s.elems = append(s.elems, elem)
}

func (s *StringStack) Pop() string {
	elem := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return elem
}

func (s *StringStack) Elems() int {
	return len(s.elems)
}
