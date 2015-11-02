package sieve

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	ret := Sieve(100)
	fmt.Println(ret)
}
