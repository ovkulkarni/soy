package soyjs

import (
	"errors"
	"io"

	"github.com/robfig/soy/soymsg"
	"github.com/robfig/soy/template"
)

type JSFormat string

type JSFormatter interface {
	Template(fmt JSFormat, name string) (string, string)
	Call(fmt JSFormat, name string) (string, string)
	Directive(fmt JSFormat, name string) (PrintDirective, string)
	Function(fmt JSFormat, name string) (Func, string)
}

const (
	ES5 = "ES5"
	ES6 = "ES6"
)

// Options for js source generation.
type Options struct {
	Messages  soymsg.Bundle
	Format    JSFormat
	Formatter JSFormatter
}

// Generator provides an interface to a template registry capable of generating
// javascript to execute the embodied templates.
// The generated javascript requires lib/soyutils.js to already have been loaded.
type Generator struct {
	registry *template.Registry
}

// NewGenerator returns a new javascript generator capable of producing
// javascript for the templates contained in the given registry.
func NewGenerator(registry *template.Registry) *Generator {
	return &Generator{registry}
}

var ErrNotFound = errors.New("file not found")

// WriteFile generates javascript corresponding to the soy file of the given name.
func (gen *Generator) WriteFile(out io.Writer, filename string) error {
	for _, soyfile := range gen.registry.SoyFiles {
		if soyfile.Name == filename {
			return Write(out, soyfile, Options{})
		}
	}
	return ErrNotFound
}
