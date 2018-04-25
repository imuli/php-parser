package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Namespace node
type Namespace struct {
	NamespaceName node.Node
	Stmt          *InnerStmtList
}

// NewNamespace node constructor
func NewNamespace(NamespaceName node.Node, Stmt *InnerStmtList) *Namespace {
	return &Namespace{
		NamespaceName,
		Stmt,
	}
}

// Attributes returns node attributes as map
func (n *Namespace) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Namespace) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.NamespaceName != nil {
		vv := v.GetChildrenVisitor("NamespaceName")
		n.NamespaceName.Walk(vv)
	}

	if n.Stmt != nil {
		vv := v.GetChildrenVisitor("Stmt")
		n.Stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
