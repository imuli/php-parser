package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// InnerTraitAdaptationList node
type InnerTraitAdaptationList struct {
	Adaptations []node.Node
}

// NewInnerTraitAdaptationList node constructor
func NewInnerTraitAdaptationList(Adaptations []node.Node) *InnerTraitAdaptationList {
	return &InnerTraitAdaptationList{
		Adaptations,
	}
}

// Attributes returns node attributes as map
func (n *InnerTraitAdaptationList) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *InnerTraitAdaptationList) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Adaptations != nil {
		vv := v.GetChildrenVisitor("Adaptations")
		for _, nn := range n.Adaptations {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
