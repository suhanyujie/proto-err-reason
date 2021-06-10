// Code generated by "stringer -type ErrCode -linecomment -output code_string.go"; DO NOT EDIT.

package example

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ERROR_REASON_UNSPECIFIED-0]
	_ = x[USER_NOT_FOUND-1]
	_ = x[Role_Create_Fail-300001001]
}

const (
	_ErrCode_name_0 = "unknownnot found"
	_ErrCode_name_1 = "新增角色失败"
)

var (
	_ErrCode_index_0 = [...]uint8{0, 7, 16}
)

func (i ErrCode) String() string {
	switch {
	case 0 <= i && i <= 1:
		return _ErrCode_name_0[_ErrCode_index_0[i]:_ErrCode_index_0[i+1]]
	case i == 300001001:
		return _ErrCode_name_1
	default:
		return "ErrCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}