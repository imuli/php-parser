package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Require struct {
	name string
	expr node.Node
}

func NewRequire(expression node.Node) node.Node {
	return Require{
		"Require",
		expression,
	}
}

func (n Require) Name() string {
	return "Require"
}

func (n Require) Attributes() map[string]interface{} {
	return nil
}

func (n Require) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}