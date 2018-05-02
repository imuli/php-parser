package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Class node
type Class struct {
	PhpDocComment     string
	ClassName         node.Node
	Modifiers         []node.Node
	InnerArgumentList *node.InnerArgumentList
	Extends           node.Node
	Implements        []node.Node
	StmtList          *StmtList
}

// NewClass node constructor
func NewClass(ClassName node.Node, Modifiers []node.Node, InnerArgumentList *node.InnerArgumentList, Extends node.Node, Implements []node.Node, StmtList *StmtList, PhpDocComment string) *Class {
	return &Class{
		PhpDocComment,
		ClassName,
		Modifiers,
		InnerArgumentList,
		Extends,
		Implements,
		StmtList,
	}
}

// Attributes returns node attributes as map
func (n *Class) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"PhpDocComment": n.PhpDocComment,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Class) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.ClassName != nil {
		vv := v.GetChildrenVisitor("ClassName")
		n.ClassName.Walk(vv)
	}

	if n.Modifiers != nil {
		vv := v.GetChildrenVisitor("Modifiers")
		for _, nn := range n.Modifiers {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	if n.InnerArgumentList != nil {
		vv := v.GetChildrenVisitor("InnerArgumentList")
		n.InnerArgumentList.Walk(vv)
	}

	if n.Extends != nil {
		vv := v.GetChildrenVisitor("Extends")
		n.Extends.Walk(vv)
	}

	if n.Implements != nil {
		vv := v.GetChildrenVisitor("Implements")
		for _, nn := range n.Implements {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	if n.StmtList != nil {
		vv := v.GetChildrenVisitor("StmtList")
		n.StmtList.Walk(vv)
	}

	v.LeaveNode(n)
}
