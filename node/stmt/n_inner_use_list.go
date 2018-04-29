package stmt

import (
	"github.com/z7zmey/php-parser/walker"
)

// InnerUseList node
type InnerUseList struct {
	UseList *UseList
}

// NewInnerUseList node constructor
func NewInnerUseList(Stmts *UseList) *InnerUseList {
	return &InnerUseList{
		Stmts,
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

	if n.UseList != nil {
		vv := v.GetChildrenVisitor("UseList")
		n.UseList.Walk(vv)
	}

	v.LeaveNode(n)
}
