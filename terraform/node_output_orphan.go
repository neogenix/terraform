package terraform

import (
	"fmt"
)

// NodeOutputOrphan represents an output that is an orphan.
type NodeOutputOrphan struct {
	OutputName string
	PathValue  []string
}

func (n *NodeOutputOrphan) Name() string {
	result := fmt.Sprintf("output.%s (orphan)", n.OutputName)
	if len(n.PathValue) > 1 {
		result = fmt.Sprintf("%s.%s", modulePrefixStr(n.PathValue), result)
	}

	return result
}

// GraphNodeSubPath
func (n *NodeOutputOrphan) Path() []string {
	return n.PathValue
}

// GraphNodeEvalable
func (n *NodeOutputOrphan) EvalTree() EvalNode {
	return &EvalDeleteOutput{
		Name: n.OutputName,
	}
}
