package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Try node
type Try struct {
	StmtList *StmtList
	Catches  []node.Node
	Finally  node.Node
}

// NewTry node constructor
func NewTry(StmtList *StmtList, Catches []node.Node, Finally node.Node) *Try {
	return &Try{
		StmtList,
		Catches,
		Finally,
	}
}

// Attributes returns node attributes as map
func (n *Try) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Try) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.StmtList != nil {
		vv := v.GetChildrenVisitor("StmtList")
		n.StmtList.Walk(vv)
	}

	if n.Catches != nil {
		vv := v.GetChildrenVisitor("Catches")
		for _, nn := range n.Catches {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	if n.Finally != nil {
		vv := v.GetChildrenVisitor("Finally")
		n.Finally.Walk(vv)
	}

	v.LeaveNode(n)
}
