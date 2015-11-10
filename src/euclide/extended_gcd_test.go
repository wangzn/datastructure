package euclide

import (
	"testing"

	"github.com/wangzn/datastructure/src/sorting/util"
)

func checkTuple(a, b int) bool {
	g := Gcd(a, b)
	x, y := ExtGcd(a, b)
	if x*a+y*b != g {
		return false
	}
	return true
}

func TestMain(t *testing.T) {
	a := util.GenArray(100)
	for i := 0; i < 100; i = i + 2 {
		r := checkTuple(a[i], a[i+1])
		if !r {
			t.Error("Error", a[i], a[i+1], Gcd(a[i], a[i+1]))
		}
	}
}
