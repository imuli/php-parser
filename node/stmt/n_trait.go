package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Trait node
type Trait struct {
	PhpDocComment string
	TraitName     node.Node
	StmtList      *StmtList
}

// NewTrait node constructor
func NewTrait(TraitName node.Node, StmtList *StmtList, PhpDocComment string) *Trait {
	return &Trait{
		PhpDocComment,
		TraitName,
		StmtList,
	}
}

// Attributes returns node attributes as map
func (n *Trait) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"PhpDocComment": n.PhpDocComment,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Trait) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.TraitName != nil {
		vv := v.GetChildrenVisitor("TraitName")
		n.TraitName.Walk(vv)
	}

	if n.StmtList != nil {
		vv := v.GetChildrenVisitor("StmtList")
		n.StmtList.Walk(vv)
	}

	v.LeaveNode(n)
}
