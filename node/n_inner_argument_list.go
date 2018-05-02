package node

import (
	"github.com/z7zmey/php-parser/walker"
)

// InnerArgumentList node
type InnerArgumentList struct {
	Arguments []Node
}

// NewInnerArgumentList node constructor
func NewInnerArgumentList(Arguments []Node) *InnerArgumentList {
	return &InnerArgumentList{
		Arguments,
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

	if n.Arguments != nil {
		vv := v.GetChildrenVisitor("Arguments")
		for _, nn := range n.Arguments {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
