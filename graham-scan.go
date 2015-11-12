package convexhull

func excluding(pl []*Point, q *Point) (excludes []*Point) {
	for _, p := range pl {
		if p != q {
			excludes = append(excludes, p)
		}
	}

	return
}

func GrahamScan(Q []*Point) (S []*Point) {
	if len(Q) < 3 {
		return Q
	}

	minPo := minYCoordPoint(Q)

	minPoExcluded := excluding(Q, minPo)
	sortedPl := minPo.PolarAngleSort(minPoExcluded)

	firsts, rest := sortedPl[0:2], sortedPl[2:]

	// In reverse since PUSH(po, S), PUSH(p1, S), PUSH(p2, S)
	S = append(S, firsts[1], firsts[0], minPo)

	for i, n := 0, len(rest); i < n; i++ {
		pi := rest[i]

		for len(S) > 1 {
			top, nextToTop := S[0], S[1]
			if pi.PolarAngle(nextToTop) >= pi.PolarAngle(top) {
				break
			}

			// POP(S)
			S = S[1:]
		}

		// PUSH(S, pi)
		S = append([]*Point{pi}, S...)
	}

	return
}
