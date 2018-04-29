package stmt

import (
	"github.com/z7zmey/php-parser/walker"
)

// Finally node
type Finally struct {
	InnerStmtList *InnerStmtList
}

// NewFinally node constructor
func NewFinally(InnerStmtList *InnerStmtList) *Finally {
	return &Finally{
		InnerStmtList,
	}
}

// Attributes returns node attributes as map
func (n *Finally) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Finally) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.InnerStmtList != nil {
		vv := v.GetChildrenVisitor("InnerStmtList")
		n.InnerStmtList.Walk(vv)
	}

	v.LeaveNode(n)
}
