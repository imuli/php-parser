package stmt_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node/name"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestNamespace(t *testing.T) {
	t.Helper()
	src := `<? namespace Foo;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Namespace{
				NamespaceName: &name.Name{
					Parts: []node.Node{
						&name.NamePart{Value: "Foo"},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestNamespaceStmts(t *testing.T) {
	t.Helper()
	src := `<? namespace Foo {}`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Namespace{
				NamespaceName: &name.Name{
					Parts: []node.Node{
						&name.NamePart{Value: "Foo"},
					},
				},
				Stmt: &stmt.InnerStmtList{
					Stmts: &stmt.StmtList{
						Stmts: []node.Node{},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestAnonymousNamespace(t *testing.T) {
	t.Helper()
	src := `<? namespace {}`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Namespace{
				Stmt: &stmt.InnerStmtList{
					Stmts: &stmt.StmtList{
						Stmts: []node.Node{},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}
