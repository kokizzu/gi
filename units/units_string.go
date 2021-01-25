// Code generated by "stringer -type=Units"; DO NOT EDIT.

package units

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Px-0]
	_ = x[Dp-1]
	_ = x[Pct-2]
	_ = x[Rem-3]
	_ = x[Em-4]
	_ = x[Ex-5]
	_ = x[Ch-6]
	_ = x[Vw-7]
	_ = x[Vh-8]
	_ = x[Vmin-9]
	_ = x[Vmax-10]
	_ = x[Cm-11]
	_ = x[Mm-12]
	_ = x[Q-13]
	_ = x[In-14]
	_ = x[Pc-15]
	_ = x[Pt-16]
	_ = x[Dot-17]
	_ = x[UnitsN-18]
}

const _Units_name = "PxDpPctRemEmExChVwVhVminVmaxCmMmQInPcPtDotUnitsN"

var _Units_index = [...]uint8{0, 2, 4, 7, 10, 12, 14, 16, 18, 20, 24, 28, 30, 32, 33, 35, 37, 39, 42, 48}

func (i Units) String() string {
	if i < 0 || i >= Units(len(_Units_index)-1) {
		return "Units(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Units_name[_Units_index[i]:_Units_index[i+1]]
}

func (i *Units) FromString(s string) error {
	for j := 0; j < len(_Units_index)-1; j++ {
		if s == _Units_name[_Units_index[j]:_Units_index[j+1]] {
			*i = Units(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: Units")
}