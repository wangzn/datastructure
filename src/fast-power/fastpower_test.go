package fastpower

import (
	"testing"
)

func TestMain(t *testing.T) {
	p := FastPower(2, 3, 31)
	if p != 2 {
		t.Error("fast power error")
	}
	p = FastPower(7, 3, 1)
	if p != 1 {
		t.Error("fast power error")
	}
}

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastPower(2, 3, 31)
	}
}
