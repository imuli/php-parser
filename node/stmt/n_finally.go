package stmt

import (
	"github.com/z7zmey/php-parser/walker"
)

// Finally node
type Finally struct {
	StmtList *StmtList
}

// NewFinally node constructor
func NewFinally(StmtList *StmtList) *Finally {
	return &Finally{
		StmtList,
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

	if n.StmtList != nil {
		vv := v.GetChildrenVisitor("StmtList")
		n.StmtList.Walk(vv)
	}

	v.LeaveNode(n)
}
