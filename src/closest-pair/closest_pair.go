package closest_pair

import (
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
