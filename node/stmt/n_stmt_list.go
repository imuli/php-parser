package stmt

import (
	"github.com/z7zmey/php-parser/walker"
)

// StmtList node
type StmtList struct {
	InnerStmtList *InnerStmtList
}

// NewStmtList node constructor
func NewStmtList(InnerStmtList *InnerStmtList) *StmtList {
	return &StmtList{
		InnerStmtList,
	}
}

// Attributes returns node attributes as map
func (n *StmtList) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *StmtList) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.InnerStmtList != nil {
		vv := v.GetChildrenVisitor("InnerStmtList")
		n.InnerStmtList.Walk(vv)
	}

	v.LeaveNode(n)
}
