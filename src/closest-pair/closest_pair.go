package closest_pair

import (
	"fmt"
	"math"
	"sort"
)

type Point struct {
	x int
	y int
}

type Points []Point
type ByY Points
type ByX Points

func (ps Points) Extreme() (minx, miny, maxx, maxy int) {
	minx = math.MaxInt64
	miny = minx
	maxx = math.MinInt64
	maxy = maxx
	for _, p := range ps {
		if p.x > maxx {
			maxx = p.x
		}
		if p.y > maxy {
			maxy = p.y
		}
		if p.x < minx {
			minx = p.x
		}
		if p.y < miny {
			miny = p.y
		}
	}
	return
}

func (ps Points) Max(by string) int {
	max := math.MinInt64
	for _, p := range ps {
		if by == "x" {
			if p.x > max {
				max = p.x
			}
		} else {
			if p.y > max {
				max = p.y
			}
		}
	}
	return max
}

func (ps Points) MaxY() int {
	return ps.Max("y")
}

func (ps Points) MaxX() int {
	return ps.Max("x")
}

func (ps Points) Print() {
	minx, miny, maxx, maxy := ps.Extreme()
	diffx := maxx - minx + 1
	diffy := maxy - miny + 1
	c := make([][]int, diffy)
	for i := 0; i < len(c); i++ {
		c[i] = make([]int, diffx)
	}
	for _, p := range ps {
		x := maxy - p.y
		y := p.x - minx
		c[x][y] = 1
	}

	for i := 0; i < len(c); i++ {
		fmt.Print("|")
		for j := 0; j < len(c[i]); j++ {
			if c[i][j] == 0 {
				fmt.Print("   ")
			} else {
				fmt.Print(" * ")
			}
		}
		fmt.Println()
	}
	for j := 0; j < len(c[0]); j++ {
		fmt.Print("---")
	}
	fmt.Println()
}

func (ps Points) Len() int {
	return len(ps)
}

func (ps Points) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps Points) Less(i, j int) bool {
	return ps[i].x < ps[j].x
}

func (ps ByY) Len() int {
	return len(ps)
}

func (ps ByY) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps ByY) Less(i, j int) bool {
	return ps[i].x < ps[j].x
}

func (ps ByX) Len() int {
	return len(ps)
}

func (ps ByX) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps ByX) Less(i, j int) bool {
	return ps[i].x < ps[j].x
}

func Dis(a, b Point) float64 {
	diffx := math.Abs(float64(a.x - b.x))
	diffy := math.Abs(float64(a.y - b.y))
	return math.Sqrt(math.Pow(diffx, 2.0) + math.Pow(diffy, 2.0))
}

func Min(ps Points) float64 {
	if len(ps) <= 1 {
		return math.MaxFloat64
	}
	if len(ps) == 2 {
		return Dis(ps[0], ps[1])
	}
	var byx ByX
	byx = ByX(ps)
	sort.Sort(byx)
	mid := len(ps) / 2
	midX := ps[mid].x
	min1 := Min(ps[:mid])
	min2 := Min(ps[mid+1:])
	min := min1
	if min2 < min {
		min = min2
	}
	temp := make([]Point, 0)
	for i := 0; i < len(ps); i++ {
		if math.Abs(float64(ps[i].x-midX)) < float64(min) {
			temp = append(temp, ps[i])
		}
	}
	var byy ByY
	byy = temp
	sort.Sort(byy)
	for i := 0; i < len(temp); i++ {
		for j := i + 1; j < i+7 && j < len(temp); j++ {
			dis := Dis(temp[i], temp[j])
			if dis < min {
				min = dis
			}
		}
	}
	return min
}
