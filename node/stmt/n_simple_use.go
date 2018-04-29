package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// SimpleUse node
type SimpleUse struct {
	UseType node.Node
	UseList *UseList
}

// NewSimpleUse node constructor
func NewSimpleUse(UseType node.Node, UseList *UseList) *SimpleUse {
	return &SimpleUse{
		UseType,
		UseList,
	}
}

// Attributes returns node attributes as map
func (n *SimpleUse) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *SimpleUse) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.UseType != nil {
		vv := v.GetChildrenVisitor("UseType")
		n.UseType.Walk(vv)
	}

	if n.UseList != nil {
		vv := v.GetChildrenVisitor("UseList")
		n.UseList.Walk(vv)
	}

	v.LeaveNode(n)
}
