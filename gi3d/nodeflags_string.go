// Code generated by "stringer -type=NodeFlags"; DO NOT EDIT.

package gi3d

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[WorldMatrixUpdated-24]
	_ = x[VectorsUpdated-25]
	_ = x[NodeFlagsN-26]
}

const _NodeFlags_name = "WorldMatrixUpdatedVectorsUpdatedNodeFlagsN"

var _NodeFlags_index = [...]uint8{0, 18, 32, 42}

func (i NodeFlags) String() string {
	i -= 24
	if i < 0 || i >= NodeFlags(len(_NodeFlags_index)-1) {
		return "NodeFlags(" + strconv.FormatInt(int64(i+24), 10) + ")"
	}
	return _NodeFlags_name[_NodeFlags_index[i]:_NodeFlags_index[i+1]]
}

func StringToNodeFlags(s string) (NodeFlags, error) {
	for i := 0; i < len(_NodeFlags_index)-1; i++ {
		if s == _NodeFlags_name[_NodeFlags_index[i]:_NodeFlags_index[i+1]] {
			return NodeFlags(i + 24), nil
		}
	}
	return 0, errors.New("String: " + s + " is not a valid option for type: NodeFlags")
}
