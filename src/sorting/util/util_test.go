package util

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	a := GenArray(10)
	fmt.Println(a)
	if len(a) != 10 {
		t.Error("Get array size error")
	}
	b := []int{1, 2, 3, 4, 5}
	c := CheckArray(b)
	if c != 0 {
		t.Error("Check corrent array error", c)
	}
	b = []int{2, 1, 3}
	c = CheckArray(b)
	if c != 1 {
		t.Error("Check invalid array error", c)
	}

}
