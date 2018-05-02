package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// SimpleUse node
type SimpleUse struct {
	UseType      node.Node
	InnerUseList *InnerUseList
}

// NewSimpleUse node constructor
func NewSimpleUse(UseType node.Node, InnerUseList *InnerUseList) *SimpleUse {
	return &SimpleUse{
		UseType,
		InnerUseList,
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

	if n.InnerUseList != nil {
		vv := v.GetChildrenVisitor("InnerUseList")
		n.InnerUseList.Walk(vv)
	}

	v.LeaveNode(n)
}
