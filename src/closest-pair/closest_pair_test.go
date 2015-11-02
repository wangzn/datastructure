package closest_pair

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	p1 := Point{3, 3}
	p2 := Point{4, 4}
	p3 := Point{5, 5}
	p4 := Point{1, 5}
	ps := make([]Point, 0)
	ps = append(ps, p1, p2, p3, p4)
	dis := Min(ps)
	fmt.Println(ps, dis)
}
