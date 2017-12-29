package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type InstanceOf struct {
	name  string
	expr  node.Node
	class node.Node
}

func NewInstanceOf(expr node.Node, class node.Node) node.Node {
	return InstanceOf{
		"InstanceOf",
		expr,
		class,
	}
}

func (n InstanceOf) Name() string {
	return "InstanceOf"
}

func (n InstanceOf) Attributes() map[string]interface{} {
	return nil
}

func (n InstanceOf) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	if n.class != nil {
		vv := v.GetChildrenVisitor("class")
		n.class.Walk(vv)
	}

	v.LeaveNode(n)
}