package stmt

import (
	"github.com/z7zmey/php-parser/walker"
)

// CaseList node
type CaseList struct {
	InnerCaseList *InnerCaseList
}

// NewCaseList node constructor
func NewCaseList(InnerCaseList *InnerCaseList) *CaseList {
	return &CaseList{
		InnerCaseList,
	}
}

// Attributes returns node attributes as map
func (n *CaseList) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *CaseList) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.InnerCaseList != nil {
		vv := v.GetChildrenVisitor("InnerCaseList")
		n.InnerCaseList.Walk(vv)
	}

	v.LeaveNode(n)
}
