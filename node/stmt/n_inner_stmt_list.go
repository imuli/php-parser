package stmt

import (
	"github.com/z7zmey/php-parser/walker"
)

// InnerStmtList node
type InnerStmtList struct {
	Stmts *StmtList
}

// NewInnerStmtList node constructor
func NewInnerStmtList(Stmts *StmtList) *InnerStmtList {
	return &InnerStmtList{
		Stmts,
	}
}

// Attributes returns node attributes as map
func (n *InnerStmtList) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *InnerStmtList) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Stmts != nil {
		vv := v.GetChildrenVisitor("Stmts")
		n.Stmts.Walk(vv)
	}

	v.LeaveNode(n)
}
