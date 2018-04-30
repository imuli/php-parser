package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Interface node
type Interface struct {
	PhpDocComment string
	InterfaceName node.Node
	Extends       []node.Node
	InnerStmtList *InnerStmtList
}

// NewInterface node constructor
func NewInterface(InterfaceName node.Node, Extends []node.Node, InnerStmtList *InnerStmtList, PhpDocComment string) *Interface {
	return &Interface{
		PhpDocComment,
		InterfaceName,
		Extends,
		InnerStmtList,
	}
}

// Attributes returns node attributes as map
func (n *Interface) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"PhpDocComment": n.PhpDocComment,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Interface) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.InterfaceName != nil {
		vv := v.GetChildrenVisitor("InterfaceName")
		n.InterfaceName.Walk(vv)
	}

	if n.Extends != nil {
		vv := v.GetChildrenVisitor("Extends")
		for _, nn := range n.Extends {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	if n.InnerStmtList != nil {
		vv := v.GetChildrenVisitor("InnerStmtList")
		n.InnerStmtList.Walk(vv)
	}

	v.LeaveNode(n)
}
