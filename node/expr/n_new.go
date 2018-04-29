package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// New node
type New struct {
	Class             node.Node
	InnerArgumentList *node.InnerArgumentList
}

// NewNew node constructor
func NewNew(Class node.Node, InnerArgumentList *node.InnerArgumentList) *New {
	return &New{
		Class,
		InnerArgumentList,
	}
}

// Attributes returns node attributes as map
func (n *New) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *New) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Class != nil {
		vv := v.GetChildrenVisitor("Class")
		n.Class.Walk(vv)
	}

	if n.InnerArgumentList != nil {
		vv := v.GetChildrenVisitor("InnerArgumentList")
		n.InnerArgumentList.Walk(vv)
	}

	v.LeaveNode(n)
}
