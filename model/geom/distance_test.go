package geom

import (
	"math"
	"testing"
)

type pointPair struct {
	A, B Point
}

type distanceTest struct {
	arg      pointPair
	expected float64
}

var PointPairsForTests = []pointPair{
	{newPoint(1, 2), newPoint(1, 2)},
	{newPoint(10, 2), newPoint(8, 2)},
	{newPoint(1, 20), newPoint(1, 10)},
	{newPoint(1, 2), newPoint(3, 4)},
	{newPoint(3, 4), newPoint(1, 2)},
	{newPoint(1.5, 1.5), newPoint(3.5, 3.5)},
	{newPoint(-1.5, -1.5), newPoint(3.5, 3.5)},
	{newPoint(0, 0), newPoint(525252.52, 525252.52)},
}

var euclideanDistanceTests = []distanceTest{
	{PointPairsForTests[0], 0},
	{PointPairsForTests[1], 2},
	{PointPairsForTests[2], 10},
	{PointPairsForTests[3], 2.828427124},
	{PointPairsForTests[4], 2.828427124},
	{PointPairsForTests[5], 2.828427124},
	{PointPairsForTests[6], 7.071067811},
	{PointPairsForTests[7], 742819.237454645},
}

func TestEuclideanDistance(t *testing.T) {
	RunDistanceTests(t, euclideanDistanceTests, EuclideanDistance)
}

var manhattanDistanceDistanceTests = []distanceTest{
	{PointPairsForTests[0], 0},
	{PointPairsForTests[1], 2},
	{PointPairsForTests[2], 10},
	{PointPairsForTests[3], 4},
	{PointPairsForTests[4], 4},
	{PointPairsForTests[5], 4},
	{PointPairsForTests[6], 10},
	{PointPairsForTests[7], 1050505.04},
}

func TestManhattanDistance(t *testing.T) {
	RunDistanceTests(t, manhattanDistanceDistanceTests, ManhattanDistance)
}

func RunDistanceTests(t *testing.T, tests []distanceTest, testingFunc Distance) {
	for _, test := range tests {
		if output := testingFunc(test.arg.A, test.arg.B); math.Abs(output-test.expected) > Eps {
			t.Errorf("Output %f not equal to expected %f", output, test.expected)
		}
	}
}
