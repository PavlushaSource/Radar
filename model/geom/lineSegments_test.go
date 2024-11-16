package geom

import (
	"testing"
)

type LineSegmentsIntersectionTest struct {
	arg1, arg2 LineSegment
	expected   bool
}

var LineSegmentsIntersectionTests = []LineSegmentsIntersectionTest{
	{arg1: LineSegment{newPoint(1, 1), newPoint(1, 1)},
		arg2:     LineSegment{newPoint(1, 1), newPoint(1, 1)},
		expected: true},
	{arg1: LineSegment{newPoint(1, 1), newPoint(1, 1)},
		arg2:     LineSegment{newPoint(2, 2), newPoint(2, 2)},
		expected: false},
	{arg1: LineSegment{newPoint(1, 1), newPoint(1, 6)},
		arg2:     LineSegment{newPoint(2, 2), newPoint(0, 2)},
		expected: true},
	{arg1: LineSegment{newPoint(1, 1), newPoint(1, 6)},
		arg2:     LineSegment{newPoint(2, 2), newPoint(3, 2)},
		expected: false},
	{arg1: LineSegment{newPoint(4, 2), newPoint(3, 4)},
		arg2:     LineSegment{newPoint(4, 2), newPoint(5, 4)},
		expected: true},
	{arg1: LineSegment{newPoint(4, 2), newPoint(3, 4)},
		arg2:     LineSegment{newPoint(4, 1.99), newPoint(5, 4)},
		expected: false},
	{arg1: LineSegment{newPoint(4, 2), newPoint(3, 4)},
		arg2:     LineSegment{newPoint(4, 2.01), newPoint(5, 4)},
		expected: false},
	{arg1: LineSegment{newPoint(4, 2), newPoint(3, 4)},
		arg2:     LineSegment{newPoint(3.99, 2), newPoint(5, 4)},
		expected: true},
	{arg1: LineSegment{newPoint(4, 2), newPoint(3, 4)},
		arg2:     LineSegment{newPoint(4, 4), newPoint(0, 4.000000001)},
		expected: false},
}

func TestLineSegmentsIntersection(t *testing.T) {
	for _, test := range LineSegmentsIntersectionTests {
		if output := IsLineSegmentsIntersect(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %t not equal to expected %t", output, test.expected)
		}
	}
}
