package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// InnerCaseList node
type InnerCaseList struct {
	Cases []node.Node
}

// NewInnerCaseList node constructor
func NewInnerCaseList(Cases []node.Node) *InnerCaseList {
	return &InnerCaseList{
		Cases,
	}
}

// Attributes returns node attributes as map
func (n *InnerCaseList) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *InnerCaseList) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Cases != nil {
		vv := v.GetChildrenVisitor("Cases")
		for _, nn := range n.Cases {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
