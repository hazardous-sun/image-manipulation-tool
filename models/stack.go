package models

type ChangesStack []interface{}

// Push :
// Inserts a new value at the top of the stack structure.
func (s *ChangesStack) Push(x interface{}) {
	*s = append(*s, x)
}

// Pop :
// Removes the value at the top of the stack and returns it.
func (s *ChangesStack) Pop() interface{} {
	h := *s
	var el interface{}
	l := len(h)
	if l == 0 {
		return el
	}
	el, *s = h[l-1], h[0:l-1]
	return el
}

// Length :
// Returns the amount of group of elements in the stack.
func (s *ChangesStack) Length() int {
	return len(*s)
}

func (s *ChangesStack) Empty() bool {
	return s.Length() == 0
}

func (s *ChangesStack) Clear() {
	*s = (*s)[:0]
}

// Constructor ---------------------------------------------------------------------------------------------------------

// NewStack :
// Returns a reference to an empty stack.
func NewStack() *ChangesStack {
	return &ChangesStack{}
}
