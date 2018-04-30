package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/walker"
)

// Closure node
type Closure struct {
	ReturnsRef    bool
	Static        bool
	PhpDocComment string
	Params        []node.Node
	Uses          []node.Node
	ReturnType    node.Node
	InnerStmtList *stmt.InnerStmtList
}

// NewClosure node constructor
func NewClosure(Params []node.Node, Uses []node.Node, ReturnType node.Node, InnerStmtList *stmt.InnerStmtList, Static bool, ReturnsRef bool, PhpDocComment string) *Closure {
	return &Closure{
		ReturnsRef,
		Static,
		PhpDocComment,
		Params,
		Uses,
		ReturnType,
		InnerStmtList,
	}
}

// Attributes returns node attributes as map
func (n *Closure) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"ReturnsRef":    n.ReturnsRef,
		"Static":        n.Static,
		"PhpDocComment": n.PhpDocComment,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Closure) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Params != nil {
		vv := v.GetChildrenVisitor("Params")
		for _, nn := range n.Params {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	if n.Uses != nil {
		vv := v.GetChildrenVisitor("Uses")
		for _, nn := range n.Uses {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	if n.ReturnType != nil {
		vv := v.GetChildrenVisitor("ReturnType")
		n.ReturnType.Walk(vv)
	}

	if n.InnerStmtList != nil {
		vv := v.GetChildrenVisitor("InnerStmtList")
		n.InnerStmtList.Walk(vv)
	}

	v.LeaveNode(n)
}
