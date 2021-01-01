// Code generated by "stringer -type=BorderStyles"; DO NOT EDIT.

package gist

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BorderSolid-0]
	_ = x[BorderDotted-1]
	_ = x[BorderDashed-2]
	_ = x[BorderDouble-3]
	_ = x[BorderGroove-4]
	_ = x[BorderRidge-5]
	_ = x[BorderInset-6]
	_ = x[BorderOutset-7]
	_ = x[BorderNone-8]
	_ = x[BorderHidden-9]
	_ = x[BorderN-10]
}

const _BorderStyles_name = "BorderSolidBorderDottedBorderDashedBorderDoubleBorderGrooveBorderRidgeBorderInsetBorderOutsetBorderNoneBorderHiddenBorderN"

var _BorderStyles_index = [...]uint8{0, 11, 23, 35, 47, 59, 70, 81, 93, 103, 115, 122}

func (i BorderStyles) String() string {
	if i < 0 || i >= BorderStyles(len(_BorderStyles_index)-1) {
		return "BorderStyles(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _BorderStyles_name[_BorderStyles_index[i]:_BorderStyles_index[i+1]]
}

func (i *BorderStyles) FromString(s string) error {
	for j := 0; j < len(_BorderStyles_index)-1; j++ {
		if s == _BorderStyles_name[_BorderStyles_index[j]:_BorderStyles_index[j+1]] {
			*i = BorderStyles(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: BorderStyles")
}