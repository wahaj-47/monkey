package object

import "fmt"

type ObjectType string

const (
	INTEGER_OBJECT = "INTEGER"
	FLOAT_OBJECT   = "FLOAT"
	BOOLEAN_OBJECT = "BOOLEAN"
	NULL_OBJECT    = "NULL"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJECT }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJECT }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJECT }
func (n *Null) Inspect() string  { return "null" }

type Float struct {
	Value float64
}

func (f *Float) Type() ObjectType { return FLOAT_OBJECT }
func (f *Float) Inspect() string  { return fmt.Sprintf("%f", f.Value) }
