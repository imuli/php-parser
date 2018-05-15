package stmt_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node/scalar"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestAltSwitch(t *testing.T) {
	t.Helper()
	src := `<? 
		switch (1) :
			case 1:
			default:
			case 2;
		endswitch;
	`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.AltSwitch{
				Cond: &scalar.Lnumber{Value: "1"},
				CaseList: &stmt.CaseList{
					InnerCaseList: &stmt.InnerCaseList{
						Cases: []node.Node{
							&stmt.Case{
								Cond: &scalar.Lnumber{Value: "1"},
								InnerStmtList: &stmt.InnerStmtList{
									Stmts: []node.Node{},
								},
							},
							&stmt.Default{
								InnerStmtList: &stmt.InnerStmtList{
									Stmts: []node.Node{},
								},
							},
							&stmt.Case{
								Cond: &scalar.Lnumber{Value: "2"},
								InnerStmtList: &stmt.InnerStmtList{
									Stmts: []node.Node{},
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

func TestAltSwitchSemicolon(t *testing.T) {
	t.Helper()
	src := `<? 
		switch (1) :;
			case 1;
			case 2;
		endswitch;
	`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.AltSwitch{
				Cond: &scalar.Lnumber{Value: "1"},
				CaseList: &stmt.CaseList{
					InnerCaseList: &stmt.InnerCaseList{
						Cases: []node.Node{
							&stmt.Case{
								Cond: &scalar.Lnumber{Value: "1"},
								InnerStmtList: &stmt.InnerStmtList{
									Stmts: []node.Node{},
								},
							},
							&stmt.Case{
								Cond: &scalar.Lnumber{Value: "2"},
								InnerStmtList: &stmt.InnerStmtList{
									Stmts: []node.Node{},
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

func TestSwitch(t *testing.T) {
	t.Helper()
	src := `<? 
		switch (1) {
			case 1: break;
			case 2: break;
		}
	`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Switch{
				Cond: &scalar.Lnumber{Value: "1"},
				CaseList: &stmt.CaseList{
					InnerCaseList: &stmt.InnerCaseList{
						Cases: []node.Node{
							&stmt.Case{
								Cond: &scalar.Lnumber{Value: "1"},
								InnerStmtList: &stmt.InnerStmtList{
									Stmts: []node.Node{
										&stmt.Break{},
									},
								},
							},
							&stmt.Case{
								Cond: &scalar.Lnumber{Value: "2"},
								InnerStmtList: &stmt.InnerStmtList{
									Stmts: []node.Node{
										&stmt.Break{},
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

func TestSwitchSemicolon(t *testing.T) {
	t.Helper()
	src := `<? 
		switch (1) {;
			case 1; break;
			case 2; break;
		}
	`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Switch{
				Cond: &scalar.Lnumber{Value: "1"},
				CaseList: &stmt.CaseList{
					InnerCaseList: &stmt.InnerCaseList{
						Cases: []node.Node{
							&stmt.Case{
								Cond: &scalar.Lnumber{Value: "1"},
								InnerStmtList: &stmt.InnerStmtList{
									Stmts: []node.Node{
										&stmt.Break{},
									},
								},
							},
							&stmt.Case{
								Cond: &scalar.Lnumber{Value: "2"},
								InnerStmtList: &stmt.InnerStmtList{
									Stmts: []node.Node{
										&stmt.Break{},
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
