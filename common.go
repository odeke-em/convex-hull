package convexhull

func minYCoordPoint(pl []*Point) (minPo *Point) {
	if len(pl) < 1 {
		return
	}
	minPo = pl[0]

	for i, n := 1, len(pl); i < n; i++ {
		cur := pl[i]
		comp := valueCompare(minPo.Y(), cur.Y())

		if comp == Less {
			minPo = cur
		} else if comp == Equal && minPo.X() > cur.Y() { // Find the left most one
			minPo = cur
		}
	}

	return
}

func valueCompare(v1, v2 float64) Comparison {
	if v1 < v2 {
		return Less
	}
	if v1 > v2 {
		return Greater
	}
	return Equal
}
