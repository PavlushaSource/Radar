package geom

import (
	"math"
	"testing"
)

// Testing Distance Calculation Functions
//
// The following scenarios are tested:
// - Both objects have positive coordinates
// - One of the objects has negative coordinates
// - Both objects have negative coordinates
// - Objects are at a large distance from each other
// - Objects have identical coordinates
// - Objects are not enclosed by a barrier
// - Objects are enclosed by a barrier
//      (+ case where the function touches but does not cross the barrier)
//
//
// The test coverage can be considered sufficient because all the main scenarios of using the algorithm are tested
//
//
// How are the tests organized?
//
// Predefined test data covering the main scenarios of the algorithm are stored in [distanceFuncArgsForTests]
//
// For each distance calculation function, `Expected` values corresponding to the test data were prepared
// These values are stored in [euclideanDistanceTests], [manhattanDistanceDistanceTests], and [curvilinearDistanceDistanceTests]
//
// Using Go's testing tools, the tested function is executed on the test data,
// followed by a correctness check (comparison of `Actual` and `Expected` results)

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
	NewBarrier(NewPoint(1, 1), NewPoint(3, 3)),
	NewBarrier(NewPoint(4, 1), NewPoint(4, 4)),
}

var distanceFuncArgsForTests = []distanceFuncArgs{
	{NewPoint(1, 2), NewPoint(1, 2), emptyBarrierSlice},
	{NewPoint(10, 2), NewPoint(8, 2), emptyBarrierSlice},
	{NewPoint(1, 20), NewPoint(1, 10), emptyBarrierSlice},
	{NewPoint(1, 2), NewPoint(3, 4), emptyBarrierSlice},
	{NewPoint(3, 4), NewPoint(1, 2), emptyBarrierSlice},
	{NewPoint(1.5, 1.5), NewPoint(3.5, 3.5), emptyBarrierSlice},
	{NewPoint(-1.5, -1.5), NewPoint(3.5, 3.5), emptyBarrierSlice},
	{NewPoint(0, 0), NewPoint(525252.52, 525252.52), emptyBarrierSlice},
	{NewPoint(1, 1), NewPoint(1, 1), barrierSlice},
	{NewPoint(0, 0), NewPoint(5, 0), barrierSlice},
	{NewPoint(0, 0), NewPoint(5, 2), barrierSlice},
	{NewPoint(1, 3), NewPoint(3, 1), barrierSlice},
	{NewPoint(3, 1), NewPoint(5, 5), barrierSlice},
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

var curvilinearDistanceDistanceTests = []distanceTest{
	{distanceFuncArgsForTests[0], 0},
	{distanceFuncArgsForTests[1], math.Pi},
	{distanceFuncArgsForTests[2], 5 * math.Pi},
	{distanceFuncArgsForTests[3], math.Sqrt(2) * math.Pi},
	{distanceFuncArgsForTests[4], math.Sqrt(2) * math.Pi},
	{distanceFuncArgsForTests[5], math.Sqrt(2) * math.Pi},
	{distanceFuncArgsForTests[6], 5 / math.Sqrt(2) * math.Pi},
	{distanceFuncArgsForTests[7], math.Sqrt(525252.52*525252.52*2) / 2 * math.Pi},
	{distanceFuncArgsForTests[8], InfDistance},
	{distanceFuncArgsForTests[9], InfDistance},
	{distanceFuncArgsForTests[10], InfDistance},
	{distanceFuncArgsForTests[11], InfDistance},
	{distanceFuncArgsForTests[12], InfDistance},
}

func TestCurvilinearDistance(t *testing.T) {
	RunDistanceTests(t, curvilinearDistanceDistanceTests, CurvilinearDistance)
}

func RunDistanceTests(t *testing.T, tests []distanceTest, testingFunc Distance) {
	for _, test := range tests {
		if output := testingFunc(test.arg.first, test.arg.second, test.arg.barriers); math.Abs(output-test.expected) > Eps {
			t.Errorf("Output %f not equal to expected %f", output, test.expected)
		}
	}
}
