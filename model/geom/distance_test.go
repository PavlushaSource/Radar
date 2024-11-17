package geom

import (
	"math"
	"testing"
)

type distanceFuncArgs struct {
	first, second Point
	barriers      []Barrier
}

type distanceTest struct {
	arg      distanceFuncArgs
	expected float64
}

var emptyBarrierSlice []Barrier

var barrierSlice = []Barrier{
	{LineSegment{newPoint(1, 1), newPoint(3, 3)}},
	{LineSegment{newPoint(4, 1), newPoint(4, 4)}},
}

var distanceFuncArgsForTests = []distanceFuncArgs{
	{newPoint(1, 2), newPoint(1, 2), emptyBarrierSlice},
	{newPoint(10, 2), newPoint(8, 2), emptyBarrierSlice},
	{newPoint(1, 20), newPoint(1, 10), emptyBarrierSlice},
	{newPoint(1, 2), newPoint(3, 4), emptyBarrierSlice},
	{newPoint(3, 4), newPoint(1, 2), emptyBarrierSlice},
	{newPoint(1.5, 1.5), newPoint(3.5, 3.5), emptyBarrierSlice},
	{newPoint(-1.5, -1.5), newPoint(3.5, 3.5), emptyBarrierSlice},
	{newPoint(0, 0), newPoint(525252.52, 525252.52), emptyBarrierSlice},
	{newPoint(1, 1), newPoint(1, 1), barrierSlice},
	{newPoint(0, 0), newPoint(5, 0), barrierSlice},
	{newPoint(0, 0), newPoint(5, 2), barrierSlice},
	{newPoint(1, 3), newPoint(3, 1), barrierSlice},
	{newPoint(3, 1), newPoint(5, 5), barrierSlice},
}

var euclideanDistanceTests = []distanceTest{
	{distanceFuncArgsForTests[0], 0},
	{distanceFuncArgsForTests[1], 2},
	{distanceFuncArgsForTests[2], 10},
	{distanceFuncArgsForTests[3], 2.828427124},
	{distanceFuncArgsForTests[4], 2.828427124},
	{distanceFuncArgsForTests[5], 2.828427124},
	{distanceFuncArgsForTests[6], 7.071067811},
	{distanceFuncArgsForTests[7], 742819.237454645},
	{distanceFuncArgsForTests[8], InfDistance},
	{distanceFuncArgsForTests[9], 5.0},
	{distanceFuncArgsForTests[10], InfDistance},
	{distanceFuncArgsForTests[11], InfDistance},
	{distanceFuncArgsForTests[12], InfDistance},
}

func TestEuclideanDistance(t *testing.T) {
	RunDistanceTests(t, euclideanDistanceTests, EuclideanDistance)
}

var manhattanDistanceDistanceTests = []distanceTest{
	{distanceFuncArgsForTests[0], 0},
	{distanceFuncArgsForTests[1], 2},
	{distanceFuncArgsForTests[2], 10},
	{distanceFuncArgsForTests[3], 4},
	{distanceFuncArgsForTests[4], 4},
	{distanceFuncArgsForTests[5], 4},
	{distanceFuncArgsForTests[6], 10},
	{distanceFuncArgsForTests[7], 1050505.04},
	{distanceFuncArgsForTests[8], InfDistance},
	{distanceFuncArgsForTests[9], 5.0},
	{distanceFuncArgsForTests[10], 7.0},
	{distanceFuncArgsForTests[11], InfDistance},
	{distanceFuncArgsForTests[12], InfDistance},
}

func TestManhattanDistance(t *testing.T) {
	RunDistanceTests(t, manhattanDistanceDistanceTests, ManhattanDistance)
}

func RunDistanceTests(t *testing.T, tests []distanceTest, testingFunc Distance) {
	for _, test := range tests {
		if output := testingFunc(test.arg.first, test.arg.second, test.arg.barriers); math.Abs(output-test.expected) > Eps {
			t.Errorf("Output %f not equal to expected %f", output, test.expected)
		}
	}
}
