package stack

import "github.com/clipperhouse/typewriter"

var templates = typewriter.TemplateSlice{
	stackTmpl,
}

var stackTmpl = &typewriter.Template{
	Name: "Stack",
	Text: `
// A structure that represents a stack data structure
// for {{.Name}} type
//
// Example:
// stack := &stack.Stack{}
// stack.Push(new(TValue))
// value, err := stack.Pop()
type {{.Name}}Stack struct {
	data []{{.Pointer}}{{.Name}}
}

// Adds an element on top of the stack
func (s *{{.Name}}Stack) Push(value {{.Pointer}}{{.Name}}) {
	s.data = append(s.data, value)
}

// Removes an element from top of the stack.
// If the stack is empty, it returns an error.
func (s *{{.Name}}Stack) Pop() ({{.Pointer}}{{.Name}}, error) {
	length := len(s.data)
	if length == 0 {
		return nil, errors.New("Stack is empty")
	}

	value := s.data[length-1]
	s.data = s.data[:length-1]
	return value, nil
}`,
}
