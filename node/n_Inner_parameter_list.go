package node

import (
	"github.com/z7zmey/php-parser/walker"
)

// InnerParameterList node
type InnerParameterList struct {
	Parameters []Node
}

// NewInnerParameterList node constructor
func NewInnerParameterList(Parameters []Node) *InnerParameterList {
	return &InnerParameterList{
		Parameters,
	}
}

// Attributes returns node attributes as map
func (n *InnerParameterList) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *InnerParameterList) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Parameters != nil {
		vv := v.GetChildrenVisitor("Parameters")
		for _, nn := range n.Parameters {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
