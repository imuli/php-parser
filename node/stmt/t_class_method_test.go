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

func TestSimpleClassMethod(t *testing.T) {
	t.Helper()
	src := `<? class foo{ function bar() {} }`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Class{
				ClassName: &node.Identifier{Value: "foo"},
				InnerStmtList: &stmt.InnerStmtList{
					Stmts: &stmt.StmtList{
						Stmts: []node.Node{
							&stmt.ClassMethod{
								PhpDocComment: "",
								MethodName:    &node.Identifier{Value: "bar"},
								InnerStmtList: &stmt.InnerStmtList{
									Stmts: &stmt.StmtList{
										Stmts: []node.Node{},
									},
								},
							},
						},
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

func TestPrivateProtectedClassMethod(t *testing.T) {
	t.Helper()
	src := `<? class foo{ final private function bar() {} protected function baz() {} }`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Class{
				ClassName: &node.Identifier{Value: "foo"},
				InnerStmtList: &stmt.InnerStmtList{
					Stmts: &stmt.StmtList{
						Stmts: []node.Node{
							&stmt.ClassMethod{
								PhpDocComment: "",
								ReturnsRef:    false,
								MethodName:    &node.Identifier{Value: "bar"},
								Modifiers: []node.Node{
									&node.Identifier{Value: "final"},
									&node.Identifier{Value: "private"},
								},
								InnerStmtList: &stmt.InnerStmtList{
									Stmts: &stmt.StmtList{
										Stmts: []node.Node{},
									},
								},
							},
							&stmt.ClassMethod{
								PhpDocComment: "",
								ReturnsRef:    false,
								MethodName:    &node.Identifier{Value: "baz"},
								Modifiers: []node.Node{
									&node.Identifier{Value: "protected"},
								},
								InnerStmtList: &stmt.InnerStmtList{
									Stmts: &stmt.StmtList{
										Stmts: []node.Node{},
									},
								},
							},
						},
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

func TestPhp5ClassMethod(t *testing.T) {
	t.Helper()
	src := `<? class foo{ public static function &bar() {} }`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Class{
				ClassName: &node.Identifier{Value: "foo"},
				InnerStmtList: &stmt.InnerStmtList{
					Stmts: &stmt.StmtList{
						Stmts: []node.Node{
							&stmt.ClassMethod{
								PhpDocComment: "",
								ReturnsRef:    true,
								MethodName:    &node.Identifier{Value: "bar"},
								Modifiers: []node.Node{
									&node.Identifier{Value: "public"},
									&node.Identifier{Value: "static"},
								},
								InnerStmtList: &stmt.InnerStmtList{
									Stmts: &stmt.StmtList{
										Stmts: []node.Node{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual := php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestPhp7ClassMethod(t *testing.T) {
	t.Helper()
	src := `<? class foo{ public static function &bar(): void {} }`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Class{
				ClassName: &node.Identifier{Value: "foo"},
				InnerStmtList: &stmt.InnerStmtList{
					Stmts: &stmt.StmtList{
						Stmts: []node.Node{
							&stmt.ClassMethod{
								PhpDocComment: "",
								ReturnsRef:    true,
								MethodName:    &node.Identifier{Value: "bar"},
								Modifiers: []node.Node{
									&node.Identifier{Value: "public"},
									&node.Identifier{Value: "static"},
								},
								ReturnType: &name.Name{
									Parts: []node.Node{
										&name.NamePart{Value: "void"},
									},
								},
								InnerStmtList: &stmt.InnerStmtList{
									Stmts: &stmt.StmtList{
										Stmts: []node.Node{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestAbstractClassMethod(t *testing.T) {
	t.Helper()
	src := `<? abstract class Foo{ abstract public function bar(); }`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Class{
				Modifiers: []node.Node{&node.Identifier{Value: "abstract"}},
				ClassName: &node.Identifier{Value: "Foo"},
				InnerStmtList: &stmt.InnerStmtList{
					Stmts: &stmt.StmtList{
						Stmts: []node.Node{
							&stmt.ClassMethod{
								PhpDocComment: "",
								ReturnsRef:    false,
								MethodName:    &node.Identifier{Value: "bar"},
								Modifiers: []node.Node{
									&node.Identifier{Value: "abstract"},
									&node.Identifier{Value: "public"},
								},
							},
						},
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

func TestPhp7AbstractClassMethod(t *testing.T) {
	t.Helper()
	src := `<? abstract class Foo{ public function bar(): void; }`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Class{
				Modifiers: []node.Node{&node.Identifier{Value: "abstract"}},
				ClassName: &node.Identifier{Value: "Foo"},
				InnerStmtList: &stmt.InnerStmtList{
					Stmts: &stmt.StmtList{
						Stmts: []node.Node{
							&stmt.ClassMethod{
								PhpDocComment: "",
								ReturnsRef:    false,
								MethodName:    &node.Identifier{Value: "bar"},
								Modifiers: []node.Node{
									&node.Identifier{Value: "public"},
								},
								ReturnType: &name.Name{
									Parts: []node.Node{
										&name.NamePart{Value: "void"},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)
}
