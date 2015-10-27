package matrix

import (
	"testing"
)

func TestMain(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6}
	m := NewMatrix(data, 3, 3)
	v, err := m.Get(1, 1)
	if err != nil || v != 5 {
		t.Error("Get value error", v, err)
	}
}
