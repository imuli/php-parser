package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Catch node
type Catch struct {
	Types         []node.Node
	Variable      node.Node
	InnerStmtList *InnerStmtList
}

// NewCatch node constructor
func NewCatch(Types []node.Node, Variable node.Node, InnerStmtList *InnerStmtList) *Catch {
	return &Catch{
		Types,
		Variable,
		InnerStmtList,
	}
}

// Attributes returns node attributes as map
func (n *Catch) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Catch) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Types != nil {
		vv := v.GetChildrenVisitor("Types")
		for _, nn := range n.Types {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	if n.InnerStmtList != nil {
		vv := v.GetChildrenVisitor("InnerStmtList")
		n.InnerStmtList.Walk(vv)
	}

	v.LeaveNode(n)
}
