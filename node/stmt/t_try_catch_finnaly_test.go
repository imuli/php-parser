package stmt_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/name"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestTry(t *testing.T) {
	t.Helper()
	src := `<? 
		try {}
	`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Try{
				StmtList: &stmt.StmtList{
					InnerStmtList: &stmt.InnerStmtList{
						Stmts: []node.Node{},
					},
				},
				Catches: []node.Node{},
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

func TestTryCatch(t *testing.T) {
	t.Helper()
	src := `<? 
		try {} catch (Exception $e) {}
	`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Try{
				StmtList: &stmt.StmtList{
					InnerStmtList: &stmt.InnerStmtList{
						Stmts: []node.Node{},
					},
				},
				Catches: []node.Node{
					&stmt.Catch{
						Types: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Exception"},
								},
							},
						},
						Variable: &expr.Variable{
							VarName: &node.Identifier{Value: "e"},
						},
						StmtList: &stmt.StmtList{
							InnerStmtList: &stmt.InnerStmtList{
								Stmts: []node.Node{},
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

func TestPhp7TryCatch(t *testing.T) {
	t.Helper()
	src := `<? 
		try {} catch (Exception|RuntimeException $e) {}
	`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Try{
				StmtList: &stmt.StmtList{
					InnerStmtList: &stmt.InnerStmtList{
						Stmts: []node.Node{},
					},
				},
				Catches: []node.Node{
					&stmt.Catch{
						Types: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Exception"},
								},
							},
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "RuntimeException"},
								},
							},
						},
						Variable: &expr.Variable{
							VarName: &node.Identifier{Value: "e"},
						},
						StmtList: &stmt.StmtList{
							InnerStmtList: &stmt.InnerStmtList{
								Stmts: []node.Node{},
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

func TestTryCatchCatch(t *testing.T) {
	t.Helper()
	src := `<? 
		try {} catch (Exception $e) {} catch (RuntimeException $e) {}
	`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Try{
				StmtList: &stmt.StmtList{
					InnerStmtList: &stmt.InnerStmtList{
						Stmts: []node.Node{},
					},
				},
				Catches: []node.Node{
					&stmt.Catch{
						Types: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Exception"},
								},
							},
						},
						Variable: &expr.Variable{
							VarName: &node.Identifier{Value: "e"},
						},
						StmtList: &stmt.StmtList{
							InnerStmtList: &stmt.InnerStmtList{
								Stmts: []node.Node{},
							},
						},
					},
					&stmt.Catch{
						Types: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "RuntimeException"},
								},
							},
						},
						Variable: &expr.Variable{
							VarName: &node.Identifier{Value: "e"},
						},
						StmtList: &stmt.StmtList{
							InnerStmtList: &stmt.InnerStmtList{
								Stmts: []node.Node{},
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

func TestTryCatchFinally(t *testing.T) {
	t.Helper()
	src := `<? 
		try {} catch (Exception $e) {} finally {}
	`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Try{
				StmtList: &stmt.StmtList{
					InnerStmtList: &stmt.InnerStmtList{
						Stmts: []node.Node{},
					},
				},
				Catches: []node.Node{
					&stmt.Catch{
						Types: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Exception"},
								},
							},
						},
						Variable: &expr.Variable{
							VarName: &node.Identifier{Value: "e"},
						},
						StmtList: &stmt.StmtList{
							InnerStmtList: &stmt.InnerStmtList{
								Stmts: []node.Node{},
							},
						},
					},
				},
				Finally: &stmt.Finally{
					StmtList: &stmt.StmtList{
						InnerStmtList: &stmt.InnerStmtList{
							Stmts: []node.Node{},
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

func TestTryCatchCatchCatch(t *testing.T) {
	t.Helper()
	src := `<? try {} catch (Exception $e) {} catch (\RuntimeException $e) {} catch (namespace\AdditionException $e) {}`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Try{
				StmtList: &stmt.StmtList{
					InnerStmtList: &stmt.InnerStmtList{
						Stmts: []node.Node{},
					},
				},
				Catches: []node.Node{
					&stmt.Catch{
						Types: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Exception"},
								},
							},
						},
						Variable: &expr.Variable{
							VarName: &node.Identifier{Value: "e"},
						},
						StmtList: &stmt.StmtList{
							InnerStmtList: &stmt.InnerStmtList{
								Stmts: []node.Node{},
							},
						},
					},
					&stmt.Catch{
						Types: []node.Node{
							&name.FullyQualified{
								Parts: []node.Node{
									&name.NamePart{Value: "RuntimeException"},
								},
							},
						},
						Variable: &expr.Variable{
							VarName: &node.Identifier{Value: "e"},
						},
						StmtList: &stmt.StmtList{
							InnerStmtList: &stmt.InnerStmtList{
								Stmts: []node.Node{},
							},
						},
					},
					&stmt.Catch{
						Types: []node.Node{
							&name.Relative{
								Parts: []node.Node{
									&name.NamePart{Value: "AdditionException"},
								},
							},
						},
						Variable: &expr.Variable{
							VarName: &node.Identifier{Value: "e"},
						},
						StmtList: &stmt.StmtList{
							InnerStmtList: &stmt.InnerStmtList{
								Stmts: []node.Node{},
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
