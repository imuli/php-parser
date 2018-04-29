package node

import (
	"github.com/z7zmey/php-parser/walker"
)

// InnerArgumentList node
type InnerArgumentList struct {
	ArgumentList *ArgumentList
}

// NewInnerArgumentList node constructor
func NewInnerArgumentList(ArgumentList *ArgumentList) *InnerArgumentList {
	return &InnerArgumentList{
		ArgumentList,
	}
}

// Attributes returns node attributes as map
func (n *InnerArgumentList) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *InnerArgumentList) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.ArgumentList != nil {
		vv := v.GetChildrenVisitor("ArgumentList")
		n.ArgumentList.Walk(vv)
	}

	v.LeaveNode(n)
}
