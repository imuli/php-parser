package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// For node
type For struct {
	Init *ForExprList
	Cond *ForExprList
	Loop *ForExprList
	Stmt node.Node
}

// NewFor node constructor
func NewFor(Init *ForExprList, Cond *ForExprList, Loop *ForExprList, Stmt node.Node) *For {
	return &For{
		Init,
		Cond,
		Loop,
		Stmt,
	}
}

// Attributes returns node attributes as map
func (n *For) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *For) Walk(v walker.Visitor) {
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
