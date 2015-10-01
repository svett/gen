package stack

import (
	"io"

	"github.com/clipperhouse/typewriter"
)

func init() {
	if err := typewriter.Register(NewWriter()); err != nil {
		panic(err)
	}
}

type writer struct{}

// Creates a new stack type writer
func NewWriter() typewriter.Interface {
	return &writer{}
}

func (tw *writer) Name() string {
	return "stack"
}

func (tw *writer) Imports(t typewriter.Type) (result []typewriter.ImportSpec) {
	return
}

func (tw *writer) Write(w io.Writer, t typewriter.Type) error {
	tag, found := t.FindTag(tw)

	if !found {
		return nil
	}

	header := "// DO NOT MODIFY. Auto-generated code."
	if _, err := w.Write([]byte(header)); err != nil {
		return err
	}

	tmpl, err := templates.ByTag(t, tag)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(w, t); err != nil {
		return err
	}
	return nil
}
