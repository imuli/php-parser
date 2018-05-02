package stmt

import (
	"github.com/z7zmey/php-parser/walker"
)

// TraitAdaptationList node
type TraitAdaptationList struct {
	InnerTraitAdaptationList *InnerTraitAdaptationList
}

// NewTraitAdaptationList node constructor
func NewTraitAdaptationList(traitAdaptationList *InnerTraitAdaptationList) *TraitAdaptationList {
	return &TraitAdaptationList{
		traitAdaptationList,
	}
}

// Attributes returns node attributes as map
func (n *TraitAdaptationList) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *TraitAdaptationList) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.InnerTraitAdaptationList != nil {
		vv := v.GetChildrenVisitor("InnerTraitAdaptationList")
		n.InnerTraitAdaptationList.Walk(vv)
	}

	v.LeaveNode(n)
}
