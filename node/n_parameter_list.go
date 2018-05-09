package node

import (
	"github.com/z7zmey/php-parser/walker"
)

// ParameterList node
type ParameterList struct {
	InnerParameterList *InnerParameterList
}

// NewParameterList node constructor
func NewParameterList(InnerParameterList *InnerParameterList) *ParameterList {
	return &ParameterList{
		InnerParameterList,
	}
}

// Attributes returns node attributes as map
func (n *ParameterList) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ParameterList) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.InnerParameterList != nil {
		vv := v.GetChildrenVisitor("InnerParameterList")
		n.InnerParameterList.Walk(vv)
	}

	v.LeaveNode(n)
}
