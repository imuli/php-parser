package stmt_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node/expr/binary"

	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/expr/assign"

	"github.com/z7zmey/php-parser/node/scalar"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestFor(t *testing.T) {
	t.Helper()
	src := `<? for($i = 0; $i < 10; $i++, $i++) {}`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.For{
				Init: &stmt.ForExprList{
					Expressions: []node.Node{
						&assign.Assign{
							Variable:   &expr.Variable{VarName: &node.Identifier{Value: "i"}},
							Expression: &scalar.Lnumber{Value: "0"},
						},
					},
				},
				Cond: &stmt.ForExprList{
					Expressions: []node.Node{
						&binary.Smaller{
							Left:  &expr.Variable{VarName: &node.Identifier{Value: "i"}},
							Right: &scalar.Lnumber{Value: "10"},
						},
					},
				},
				Loop: &stmt.ForExprList{
					Expressions: []node.Node{
						&expr.PostInc{
							Variable: &expr.Variable{VarName: &node.Identifier{Value: "i"}},
						},
						&expr.PostInc{
							Variable: &expr.Variable{VarName: &node.Identifier{Value: "i"}},
						},
					},
				},
				Stmt: &stmt.StmtList{
					InnerStmtList: &stmt.InnerStmtList{Stmts: []node.Node{}},
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

func TestAltFor(t *testing.T) {
	t.Helper()
	src := `<? for(; $i < 10; $i++) : endfor;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.AltFor{
				Init: &stmt.ForExprList{
					Expressions: nil,
				},
				Cond: &stmt.ForExprList{
					Expressions: []node.Node{
						&binary.Smaller{
							Left:  &expr.Variable{VarName: &node.Identifier{Value: "i"}},
							Right: &scalar.Lnumber{Value: "10"},
						},
					},
				},
				Loop: &stmt.ForExprList{
					Expressions: []node.Node{
						&expr.PostInc{
							Variable: &expr.Variable{VarName: &node.Identifier{Value: "i"}},
						},
					},
				},
				Stmt: &stmt.StmtList{InnerStmtList: &stmt.InnerStmtList{Stmts: []node.Node{}}},
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
