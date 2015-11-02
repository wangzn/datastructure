package closest_pair

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	p1 := Point{3, 3}
	p2 := Point{4, 4}
	p3 := Point{5, 5}
	ps := make([]Point, 0)
	ps = append(ps, p1, p2, p3)
	Points(ps).Print()
	p4 := Point{1, 5}
	p5 := Point{2, 6}
	p6 := Point{6, 2}
	p7 := Point{4, 3}
	p8 := Point{1, 2}
	ps = make([]Point, 0)
	ps = append(ps, p1, p2, p3, p4, p5, p6, p7, p8)
	dis := Min(ps)
	fmt.Println(ps, dis)
	Points(ps).Print()
}
