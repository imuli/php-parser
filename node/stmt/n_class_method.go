package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// ClassMethod node
type ClassMethod struct {
	ReturnsRef    bool
	PhpDocComment string
	MethodName    node.Node
	Modifiers     []node.Node
	ParameterList *node.ParameterList
	ReturnType    node.Node
	StmtList      *StmtList
}

// NewClassMethod node constructor
func NewClassMethod(MethodName node.Node, Modifiers []node.Node, ReturnsRef bool, ParameterList *node.ParameterList, ReturnType node.Node, StmtList *StmtList, PhpDocComment string) *ClassMethod {
	return &ClassMethod{
		ReturnsRef,
		PhpDocComment,
		MethodName,
		Modifiers,
		ParameterList,
		ReturnType,
		StmtList,
	}
}

// Attributes returns node attributes as map
func (n *ClassMethod) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"ReturnsRef":    n.ReturnsRef,
		"PhpDocComment": n.PhpDocComment,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ClassMethod) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.MethodName != nil {
		vv := v.GetChildrenVisitor("MethodName")
		n.MethodName.Walk(vv)
	}

	if n.Modifiers != nil {
		vv := v.GetChildrenVisitor("Modifiers")
		for _, nn := range n.Modifiers {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	if n.ParameterList != nil {
		vv := v.GetChildrenVisitor("ParameterList")
		n.ParameterList.Walk(vv)
	}

	if n.ReturnType != nil {
		vv := v.GetChildrenVisitor("ReturnType")
		n.ReturnType.Walk(vv)
	}

	if n.StmtList != nil {
		vv := v.GetChildrenVisitor("StmtList")
		n.StmtList.Walk(vv)
	}

	v.LeaveNode(n)
}
