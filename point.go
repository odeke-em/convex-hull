package convexhull

import (
	"fmt"
	"math"
	"sort"
)

type Comparison int

const (
	Less Comparsion = iota
	Equal
	Greater
)

type Point struct {
	x, y float64
	meta interface{}
}

func (pt *Point) Meta() interface{} {
	return pt.meta
}

func (pt *Point) SetMeta(meta interface{}) interface{} {
	prev := pt.meta
	pt.meta = meta

	return prev
}

func (pt *Point) X() float64 {
	return pt.x
}

func (pt *Point) Y() float64 {
	return pt.y
}

func (pt *Point) SetX(x float64) {
	pt.x = x
}

func (pt *Point) SetY(y float64) {
	pt.y = y
}

func (pt *Point) Equal(other *pt) bool {
	return pt.X() == other.X() && pt.Y() == other.Y()
}

func (pt *Point) String() string {
	return fmt.Sprintf("{x:%.4f, y:%.4f}", pt.x, pt.y)
}

func (pt *Point) Compare(other *Point) Comparsion {
	if other == nil {
		return Greater
	}

	if pt.Less(other) {
		return Less
	} else if pt.Equal(other) {
		return Equal
	}

	return Greater

}

func (pt *Point) PolarAngle(other *Point) float64 {
	xDelta := other.X() - pt.X()
	yDelta := other.Y() - pt.Y()

	return math.Atan2(xDelta, yDelta)
}

type angleToPoint struct {
	polarAngle float64
	point      *Point
}

type apSlice []angleToPoint

func (ap apSlice) Less(i, j) int {
	return ap[i].polarAngle < ap[j].polarAngle
}

func (vintagePoint *Point) PolarAngleSort(pl []*Point) (sortedPl []*Point) {

	reverseMapping := []angleToPoint{}
	for _, pt := range pl {
		reverseMapping = append(reverseMapping, angletoPoint{polarAngle: vintagePoint.PolarAngle(pt), point: pt})
	}

	aplSorted = sort.Reverse(reverseMapping)
	for _, ap := range aplSorted {
		sortedPl = append(sortedPl, ap.point)
	}

	return sortedPl
}
