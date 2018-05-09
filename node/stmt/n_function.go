package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Function node
type Function struct {
	ReturnsRef    bool
	PhpDocComment string
	FunctionName  node.Node
	ParameterList *node.ParameterList
	ReturnType    node.Node
	StmtList      *StmtList
}

// NewFunction node constructor
func NewFunction(FunctionName node.Node, ReturnsRef bool, Params *node.ParameterList, ReturnType node.Node, StmtList *StmtList, PhpDocComment string) *Function {
	return &Function{
		ReturnsRef,
		PhpDocComment,
		FunctionName,
		Params,
		ReturnType,
		StmtList,
	}
}

// Attributes returns node attributes as map
func (n *Function) Attributes() map[string]interface{} {
	// return n.attributes
	return map[string]interface{}{
		"ReturnsRef":    n.ReturnsRef,
		"PhpDocComment": n.PhpDocComment,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Function) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.FunctionName != nil {
		vv := v.GetChildrenVisitor("FunctionName")
		n.FunctionName.Walk(vv)
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
