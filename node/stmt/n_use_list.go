package stmt

import (
	"github.com/z7zmey/php-parser/walker"
)

// UseList node
type UseList struct {
	InnerUseList *InnerUseList
}

// NewUseList node constructor
func NewUseList(Stmts *InnerUseList) *UseList {
	return &UseList{
		Stmts,
	}
}

// Attributes returns node attributes as map
func (n *UseList) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *UseList) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.InnerUseList != nil {
		vv := v.GetChildrenVisitor("InnerUseList")
		n.InnerUseList.Walk(vv)
	}

	v.LeaveNode(n)
}
