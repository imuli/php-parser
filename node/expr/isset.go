package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Isset struct {
	name      string
	variables []node.Node
}

func NewIsset(variables []node.Node) node.Node {
	return Isset{
		"Isset",
		variables,
	}
}

func (n Isset) Name() string {
	return "Isset"
}

func (n Isset) Attributes() map[string]interface{} {
	return nil
}

func (n Isset) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.variables != nil {
		vv := v.GetChildrenVisitor("variables")
		for _, nn := range n.variables {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}