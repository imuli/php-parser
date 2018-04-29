package stmt

import (
	"github.com/z7zmey/php-parser/walker"
)

// InnerTraitAdaptationList node
type InnerTraitAdaptationList struct {
	TraitAdaptationList *TraitAdaptationList
}

// NewInnerTraitAdaptationList node constructor
func NewInnerTraitAdaptationList(traitAdaptationList *TraitAdaptationList) *InnerTraitAdaptationList {
	return &InnerTraitAdaptationList{
		traitAdaptationList,
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

	if n.TraitAdaptationList != nil {
		vv := v.GetChildrenVisitor("TraitAdaptationList")
		n.TraitAdaptationList.Walk(vv)
	}

	v.LeaveNode(n)
}
