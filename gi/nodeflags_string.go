// Code generated by "stringer -type=NodeFlags"; DO NOT EDIT.

package gi

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NoLayout-11]
	_ = x[EventsConnected-12]
	_ = x[CanFocus-13]
	_ = x[HasFocus-14]
	_ = x[FullReRender-15]
	_ = x[ReRenderAnchor-16]
	_ = x[Invisible-17]
	_ = x[Inactive-18]
	_ = x[Selected-19]
	_ = x[MouseHasEntered-20]
	_ = x[DNDHasEntered-21]
	_ = x[NodeDragging-22]
	_ = x[InstaDrag-23]
	_ = x[NodeFlagsN-24]
	_ = x[TextFieldFocusActive-24]
}

const _NodeFlags_name = "NoLayoutEventsConnectedCanFocusHasFocusFullReRenderReRenderAnchorInvisibleInactiveSelectedMouseHasEnteredDNDHasEnteredNodeDraggingInstaDragNodeFlagsN"

var _NodeFlags_index = [...]uint8{0, 8, 23, 31, 39, 51, 65, 74, 82, 90, 105, 118, 130, 139, 149}

func (i NodeFlags) String() string {
	i -= 11
	if i < 0 || i >= NodeFlags(len(_NodeFlags_index)-1) {
		return "NodeFlags(" + strconv.FormatInt(int64(i+11), 10) + ")"
	}
	return _NodeFlags_name[_NodeFlags_index[i]:_NodeFlags_index[i+1]]
}

func StringToNodeFlags(s string) (NodeFlags, error) {
	for i := 0; i < len(_NodeFlags_index)-1; i++ {
		if s == _NodeFlags_name[_NodeFlags_index[i]:_NodeFlags_index[i+1]] {
			return NodeFlags(i + 11), nil
		}
	}
	return 0, errors.New("String: " + s + " is not a valid option for type: NodeFlags")
}
