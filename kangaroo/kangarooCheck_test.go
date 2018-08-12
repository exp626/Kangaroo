package kangaroo

import (
	"fmt"
	"testing"
)

func Test_KangarooCheck(t *testing.T) {

	type testpair struct {
		values []int
		res    bool
	}

	var tests = []testpair{
		{[]int{0, 3, 4, 2}, true},
		{[]int{0, 2, 5, 3}, false},
		{[]int{0, 2, 5, -3}, true},
		{[]int{0, 6, 4, 3}, false},
		{[]int{0, -2, 4, -3}, true},
		{[]int{-10000, -2, 4, -3}, true},
		{[]int{10000, -2, -10000, 3}, true},
		{[]int{0, 2, 4, 2}, false},
	}

	for _, data := range tests {
		if KangarooCheck(data.values[0], data.values[1], data.values[2], data.values[3]) != data.res {
			fmt.Println("Fail with data:")
			fmt.Println(data.values[0], data.values[1], data.values[2], data.values[3])
			t.Fail()
		}
	}
}
