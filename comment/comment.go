package comment

import (
	"github.com/z7zmey/php-parser/position"
)

type Side int

const (
	Before Side = iota
	After
)

type Kind int

const (
	Outline Kind = iota
	Inline
)

// Comment aggrigates information about comment /**
type Comment struct {
	value    string
	position *position.Position
	side     Side
	kind     Kind
}

// NewComment - Comment constructor
func NewComment(value string, pos *position.Position) *Comment {
	return &Comment{
		value,
		pos,
		Before,
		Outline,
	}
}

// SetSide - set comment side [before = 0, after = 1]
func (c *Comment) SetSide(side Side) {
	c.side = side
}

// Side - returns comment side [before = 0, after = 1]
func (c *Comment) Side() Side {
	return c.side
}

// SetKind - set comment kind [Outline = 0, Inline = 1]
func (c *Comment) SetKind(kind Kind) {
	c.kind = kind
}

// Kind - returns comment kind [Outline = 0, Inline = 1]
func (c *Comment) Kind() Kind {
	return c.kind
}

func (c *Comment) String() string {
	return c.value
}

// Position returns comment position
func (c *Comment) Position() *position.Position {
	return c.position
}
