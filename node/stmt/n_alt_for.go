package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// AltFor node
type AltFor struct {
	Init *ForExprList
	Cond *ForExprList
	Loop *ForExprList
	Stmt node.Node
}

// NewAltFor node constructor
func NewAltFor(Init *ForExprList, Cond *ForExprList, Loop *ForExprList, Stmt node.Node) *AltFor {
	return &AltFor{
		Init,
		Cond,
		Loop,
		Stmt,
	}
}

// Attributes returns node attributes as map
func (n *AltFor) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *AltFor) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Init != nil {
		vv := v.GetChildrenVisitor("Init")
		n.Init.Walk(vv)
	}

	if n.Cond != nil {
		vv := v.GetChildrenVisitor("Cond")
		n.Cond.Walk(vv)
	}

	if n.Loop != nil {
		vv := v.GetChildrenVisitor("Loop")
		n.Loop.Walk(vv)
	}

	if n.Stmt != nil {
		vv := v.GetChildrenVisitor("Stmt")
		n.Stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
