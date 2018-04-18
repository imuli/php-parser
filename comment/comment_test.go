package comment_test

import (
	"testing"

	"github.com/z7zmey/php-parser/position"

	"github.com/z7zmey/php-parser/comment"
)

func TestCommentGetPosition(t *testing.T) {
	expected := position.NewPosition(0, 0, 0, 0)

	comment := comment.NewComment("/** hello world */", expected)

	actual := comment.Position()

	if expected != actual {
		t.Errorf("expected and actual are not equal\n")
	}
}

func TestCommentPrint(t *testing.T) {
	expected := "/** hello world */"

	comment := comment.NewComment(expected, nil)

	actual := comment.String()

	if expected != actual {
		t.Errorf("expected and actual are not equal\n")
	}
}

func TestCommentDefaultSide(t *testing.T) {
	c := comment.NewComment("/** hello world */", nil)

	expected := comment.Before
	actual := c.Side()

	if expected != actual {
		t.Errorf("expected and actual are not equal\n")
	}
}

func TestCommentSetSide(t *testing.T) {
	c := comment.NewComment("/** hello world */", nil)

	expected := comment.After
	c.SetSide(expected)
	actual := c.Side()

	if expected != actual {
		t.Errorf("expected and actual are not equal\n")
	}
}

func TestCommentDefaultKind(t *testing.T) {
	c := comment.NewComment("/** hello world */", nil)

	expected := comment.Outline
	actual := c.Kind()

	if expected != actual {
		t.Errorf("expected and actual are not equal\n")
	}
}

func TestCommentSetKind(t *testing.T) {
	c := comment.NewComment("/** hello world */", nil)

	expected := comment.Inline
	c.SetKind(expected)
	actual := c.Kind()

	if expected != actual {
		t.Errorf("expected and actual are not equal\n")
	}
}


