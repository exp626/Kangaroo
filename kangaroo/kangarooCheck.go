package kangaroo

import "math"

// KangarooCheck checks will be 2 kangaroos in one place, if their starting position and jump length are known
// x1 float64 start point first kangaroo
// v1 float64 step first kangaroo
// x2 float64 start point second kangaroo
// v2 float64 step second kangaroo
// -10000 ≤ x1, x2 ≤ 10000
// -10000 ≤ v1, v2 ≤ 10000
func KangarooCheck(x1 int, v1 int, x2 int, v2 int) bool {

	if x1 < x2 && v1 < v2 {
		return false
	}

	xd := x1 - x2
	vd := v2 - v1

	// division by zero
	if xd == 0 && vd == 0 {
		return true
	} else if vd == 0 {
		return false
	}

	// jumping crossing
	mod := math.Mod(float64(xd), float64(vd))
	if mod != 0 {
		return false
	}

	val := xd / vd
	if val < 0 {
		return false
	}
	return true
}
