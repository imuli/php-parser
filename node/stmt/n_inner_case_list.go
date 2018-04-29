package stmt

import (
	"github.com/z7zmey/php-parser/walker"
)

// InnerCaseList node
type InnerCaseList struct {
	CaseList *CaseList
}

// NewInnerCaseList node constructor
func NewInnerCaseList(CaseList *CaseList) *InnerCaseList {
	return &InnerCaseList{
		CaseList,
	}
}

// Attributes returns node attributes as map
func (n *InnerCaseList) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *InnerCaseList) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.CaseList != nil {
		vv := v.GetChildrenVisitor("CaseList")
		n.CaseList.Walk(vv)
	}

	v.LeaveNode(n)
}
