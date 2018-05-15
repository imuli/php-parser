package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Case node
type Case struct {
	Cond          node.Node
	InnerStmtList *InnerStmtList
}

// NewCase node constructor
func NewCase(Cond node.Node, InnerStmtList *InnerStmtList) *Case {
	return &Case{
		Cond,
		InnerStmtList,
	}
}

// Attributes returns node attributes as map
func (n *Case) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Case) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Cond != nil {
		vv := v.GetChildrenVisitor("Cond")
		n.Cond.Walk(vv)
	}

	if n.InnerStmtList != nil {
		vv := v.GetChildrenVisitor("InnerStmtList")
		n.InnerStmtList.Walk(vv)
	}

	v.LeaveNode(n)
}
