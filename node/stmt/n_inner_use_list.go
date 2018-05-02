package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// InnerUseList node
type InnerUseList struct {
	Uses []node.Node
}

// NewInnerUseList node constructor
func NewInnerUseList(Uses []node.Node) *InnerUseList {
	return &InnerUseList{
		Uses,
	}
}

// Attributes returns node attributes as map
func (n *InnerUseList) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *InnerUseList) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Uses != nil {
		vv := v.GetChildrenVisitor("Uses")
		for _, nn := range n.Uses {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
