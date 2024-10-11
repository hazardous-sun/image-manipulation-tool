package models

type ChangesStack []interface{}

func (s *ChangesStack) Push(x interface{}) {
	*s = append(*s, x)
}

func (s *ChangesStack) Pop() interface{} {
	h := *s
	var el interface{}
	l := len(h)
	el, *s = h[l-1], h[0:l-1]
	return el
}

func (s *ChangesStack) Length() int {
	return len(*s)
}

func (s *ChangesStack) Empty() bool {
	return s.Length() == 0
}

func (s *ChangesStack) Clear() {
	*s = (*s)[:0]
}

func NewStack() *ChangesStack {
	return &ChangesStack{}
}
