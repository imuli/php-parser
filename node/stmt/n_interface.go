package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Interface node
type Interface struct {
	PhpDocComment string
	InterfaceName node.Node
	Extends       *InterfaceExtends
	StmtList      *StmtList
}

// NewInterface node constructor
func NewInterface(InterfaceName node.Node, Extends *InterfaceExtends, StmtList *StmtList, PhpDocComment string) *Interface {
	return &Interface{
		PhpDocComment,
		InterfaceName,
		Extends,
		StmtList,
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
		n.Extends.Walk(vv)
	}

	if n.StmtList != nil {
		vv := v.GetChildrenVisitor("StmtList")
		n.StmtList.Walk(vv)
	}

	v.LeaveNode(n)
}
