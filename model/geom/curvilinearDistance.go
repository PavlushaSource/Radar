package geom

import "math"

// CurvilinearDistance returns a distance between two points, which is considered to be half of the length of the circle
// constructed through these points.
func CurvilinearDistance(first Point, second Point, barriers []Barrier) float64 {
	if !arePointsAchievable(first, second, barriers, CurvilinearAchievability) {
		return InfDistance
	}

	radius := EuclideanDistance(first, second, make([]Barrier, 0)) / 2
	return math.Pi * radius
}

// CurvilinearAchievability returns true if the achievability is performed on both semicircles
func CurvilinearAchievability(first Point, second Point, barrier Barrier) bool {
	return len(IntersectCurvilinearAndBarrier(first, second, barrier)) == 0
}

// IntersectCurvilinearAndBarrier returns the points of intersection of the curvilinear function and the barrier
func IntersectCurvilinearAndBarrier(first Point, second Point, barrier Barrier) []Point {
	ans := make([]Point, 0)

	if barrierIsPoint(barrier) && circleIsPoint(first, second) {
		if barrier.StartPoint().X() == first.X() && barrier.FinishPoint().Y() == first.Y() {
			p := NewPoint(first.X(), first.Y())
			if isPointOnBarrier(p, barrier) {
				ans = append(ans, p)
			}
		}

		return ans
	}

	R := EuclideanDistance(first, second, make([]Barrier, 0)) / 2
	X0 := (second.X()-first.X())/2 + first.X()
	Y0 := (second.Y()-first.Y())/2 + first.Y()

	if barrierIsPoint(barrier) {
		A := (barrier.StartPoint().X() - X0) * (barrier.StartPoint().X() - X0)
		B := (barrier.StartPoint().Y() - Y0) * (barrier.StartPoint().Y() - Y0)

		if A+B == R*R {
			p := NewPoint(barrier.StartPoint().X(), barrier.StartPoint().Y())
			if isPointOnBarrier(p, barrier) {
				ans = append(ans, p)
			}
		}

		return ans
	}

	if circleIsPoint(first, second) {
		if isPointOnBarrier(first, barrier) {
			ans = append(ans, NewPoint(first.X(), first.Y()))
		}
		return ans
	}

	A := barrier.FinishPoint().Y() - barrier.StartPoint().Y()
	B := barrier.StartPoint().X()
	C := barrier.FinishPoint().X() - barrier.StartPoint().X()
	D := barrier.StartPoint().Y()

	if C == 0 {
		E := (B - X0) * (B - X0)
		F := R * R
		G := 2 * Y0
		H := Y0 * Y0
		I := H + E - F

		Di := G*G - 4*I

		if Di == 0 {
			X := B
			Y := G / 2

			p := NewPoint(X, Y)

			if isPointOnBarrier(p, barrier) {
				ans = append(ans, p)
			}
		}

		if Di > 0 {
			X := B
			Y := (G + math.Sqrt(Di)) / 2

			p := NewPoint(X, Y)

			if isPointOnBarrier(p, barrier) {
				ans = append(ans, p)
			}

			X1 := B
			Y1 := (G - math.Sqrt(Di)) / 2

			p2 := NewPoint(X1, Y1)

			if isPointOnBarrier(p2, barrier) {
				ans = append(ans, p2)
			}

		}

		return ans
	}

	E := A / C
	F := -E*B + D

	G := F - Y0

	H := 2 * X0
	I := X0 * X0
	J := E * E
	K := 2 * E * G
	L := G * G
	M := R * R

	N := J + 1
	O := I + L - M
	P := -H + K

	Di := P*P - 4*N*O

	if Di == 0 {
		X := -P / (2 * N)
		Y := E*X + F

		p := NewPoint(X, Y)

		if isPointOnBarrier(p, barrier) {
			ans = append(ans, p)
		}
	}

	if Di > 0 {
		X1 := (-P + math.Sqrt(Di)) / (2 * N)
		Y1 := E*X1 + F

		p := NewPoint(X1, Y1)

		if isPointOnBarrier(p, barrier) {
			ans = append(ans, p)
		}

		X2 := (-P - math.Sqrt(Di)) / (2 * N)
		Y2 := E*X2 + F

		p2 := NewPoint(X2, Y2)

		if isPointOnBarrier(p2, barrier) {
			ans = append(ans, p2)
		}
	}

	return ans
}

// barrierIsPoint returns true if the barrier is a point
func barrierIsPoint(barrier Barrier) bool {
	return barrier.StartPoint().X() == barrier.FinishPoint().X() &&
		barrier.StartPoint().Y() == barrier.FinishPoint().Y()
}

// circleIsPoint returns true is the circle is a point
func circleIsPoint(first Point, second Point) bool {
	return first.X() == second.X() && first.Y() == second.Y()
}

// isPointOnBarrier returns true if the point on the barrier
func isPointOnBarrier(point Point, barrier Barrier) bool {
	A := barrier.FinishPoint().Y() - barrier.StartPoint().Y()
	B := barrier.StartPoint().X()
	C := barrier.FinishPoint().X() - barrier.StartPoint().X()
	D := barrier.StartPoint().Y()

	minX := math.Min(barrier.StartPoint().X(), barrier.FinishPoint().X())
	maxX := math.Max(barrier.StartPoint().X(), barrier.FinishPoint().X())

	minY := math.Min(barrier.StartPoint().Y(), barrier.FinishPoint().Y())
	maxY := math.Max(barrier.StartPoint().Y(), barrier.FinishPoint().Y())

	return A*point.X()-B*A == C*point.Y()-D*C &&
		minX <= point.X() && point.X() <= maxX &&
		minY <= point.Y() && point.Y() <= maxY
}
