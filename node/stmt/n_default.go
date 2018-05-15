package stmt

import (
	"github.com/z7zmey/php-parser/walker"
)

// Default node
type Default struct {
	InnerStmtList *InnerStmtList
}

// NewDefault node constructor
func NewDefault(InnerStmtList *InnerStmtList) *Default {
	return &Default{
		InnerStmtList,
	}
}

// Attributes returns node attributes as map
func (n *Default) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Default) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.InnerStmtList != nil {
		vv := v.GetChildrenVisitor("InnerStmtList")
		n.InnerStmtList.Walk(vv)
	}

	v.LeaveNode(n)
}
