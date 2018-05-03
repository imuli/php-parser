package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// ForExprList node
type ForExprList struct {
	Expressions []node.Node
}

// NewForExprList node constructor
func NewForExprList(Expressions []node.Node) *ForExprList {
	return &ForExprList{
		Expressions,
	}
}

// Attributes returns node attributes as map
func (n *ForExprList) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ForExprList) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expressions != nil {
		vv := v.GetChildrenVisitor("Expressions")
		for _, nn := range n.Expressions {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
