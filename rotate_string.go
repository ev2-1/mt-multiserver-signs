// Code generated by "stringer -type Rotate"; DO NOT EDIT.

package signs

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[North-0]
	_ = x[North22_5-1]
	_ = x[North45-2]
	_ = x[North67_5-3]
	_ = x[East-4]
	_ = x[East22_5-5]
	_ = x[East45-6]
	_ = x[East67_5-7]
	_ = x[South-8]
	_ = x[South22_5-9]
	_ = x[South45-10]
	_ = x[South67_5-11]
	_ = x[West-12]
	_ = x[West22_5-13]
	_ = x[West45-14]
	_ = x[West67_5-15]
}

const _Rotate_name = "NorthNorth22_5North45North67_5EastEast22_5East45East67_5SouthSouth22_5South45South67_5WestWest22_5West45West67_5"

var _Rotate_index = [...]uint8{0, 5, 14, 21, 30, 34, 42, 48, 56, 61, 70, 77, 86, 90, 98, 104, 112}

func (i Rotate) String() string {
	if i >= Rotate(len(_Rotate_index)-1) {
		return "Rotate(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Rotate_name[_Rotate_index[i]:_Rotate_index[i+1]]
}
