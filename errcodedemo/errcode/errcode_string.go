// Code generated by "stringer -type ErrCode -linecomment"; DO NOT EDIT.

package errcode

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ErrUnknown-1]
	_ = x[ErrNotFound-2]
}

const _ErrCode_name = "未知错误未找到"

var _ErrCode_index = [...]uint8{0, 12, 21}

func (i ErrCode) String() string {
	i -= 1
	if i < 0 || i >= ErrCode(len(_ErrCode_index)-1) {
		return "ErrCode(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _ErrCode_name[_ErrCode_index[i]:_ErrCode_index[i+1]]
}
