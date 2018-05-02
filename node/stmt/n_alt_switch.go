package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// AltSwitch node
type AltSwitch struct {
	Cond          node.Node
	InnerCaseList *InnerCaseList
}

// NewAltSwitch node constructor
func NewAltSwitch(Cond node.Node, InnerCaseList *InnerCaseList) *AltSwitch {
	return &AltSwitch{
		Cond,
		InnerCaseList,
	}
}

// Attributes returns node attributes as map
func (n *AltSwitch) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *AltSwitch) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Cond != nil {
		vv := v.GetChildrenVisitor("Cond")
		n.Cond.Walk(vv)
	}

	if n.InnerCaseList != nil {
		vv := v.GetChildrenVisitor("InnerCaseList")
		n.InnerCaseList.Walk(vv)
	}

	v.LeaveNode(n)
}
