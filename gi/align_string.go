// Code generated by "stringer -type=Align"; DO NOT EDIT.

package gi

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

const _Align_name = "AlignLeftAlignTopAlignCenterAlignMiddleAlignRightAlignBottomAlignBaselineAlignJustifyAlignSpaceAroundAlignFlexStartAlignFlexEndAlignTextTopAlignTextBottomAlignSubAlignSuperAlignN"

var _Align_index = [...]uint8{0, 9, 17, 28, 39, 49, 60, 73, 85, 101, 115, 127, 139, 154, 162, 172, 178}

func (i Align) String() string {
	if i < 0 || i >= Align(len(_Align_index)-1) {
		return "Align(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Align_name[_Align_index[i]:_Align_index[i+1]]
}

func (i *Align) FromString(s string) error {
	for j := 0; j < len(_Align_index)-1; j++ {
		if s == _Align_name[_Align_index[j]:_Align_index[j+1]] {
			*i = Align(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: Align")
}