// Code generated by "stringer -type=VpFlags"; DO NOT EDIT.

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
	_ = x[VpFlagPopup-29]
	_ = x[VpFlagMenu-30]
	_ = x[VpFlagCompleter-31]
	_ = x[VpFlagCorrector-32]
	_ = x[VpFlagTooltip-33]
	_ = x[VpFlagPopupDestroyAll-34]
	_ = x[VpFlagSVG-35]
	_ = x[VpFlagUpdatingNode-36]
	_ = x[VpFlagNeedsFullRender-37]
	_ = x[VpFlagDoingFullRender-38]
	_ = x[VpFlagsN-39]
}

const _VpFlags_name = "VpFlagPopupVpFlagMenuVpFlagCompleterVpFlagCorrectorVpFlagTooltipVpFlagPopupDestroyAllVpFlagSVGVpFlagUpdatingNodeVpFlagNeedsFullRenderVpFlagDoingFullRenderVpFlagsN"

var _VpFlags_index = [...]uint8{0, 11, 21, 36, 51, 64, 85, 94, 112, 133, 154, 162}

func (i VpFlags) String() string {
	i -= 29
	if i < 0 || i >= VpFlags(len(_VpFlags_index)-1) {
		return "VpFlags(" + strconv.FormatInt(int64(i+29), 10) + ")"
	}
	return _VpFlags_name[_VpFlags_index[i]:_VpFlags_index[i+1]]
}

func StringToVpFlags(s string) (VpFlags, error) {
	for i := 0; i < len(_VpFlags_index)-1; i++ {
		if s == _VpFlags_name[_VpFlags_index[i]:_VpFlags_index[i+1]] {
			return VpFlags(i + 29), nil
		}
	}
	return 0, errors.New("String: " + s + " is not a valid option for type: VpFlags")
}