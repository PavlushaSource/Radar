package geom

// Distance is the interface of the function that counts the distance between two point,
// taking into account barriers on the field
type Distance func(first Point, second Point, barriers []Barrier) float64

// achievability is the interface of the function that check the achievability of two points,
// taking into account barriers on the field
type achievability func(first Point, second Point, barrier Barrier) bool

// arePointsAchievable returns `true` if two points is achievable
func arePointsAchievable(first Point, second Point, barriers []Barrier, achievabilityFunc achievability) bool {
	var flag = true
	for _, barrier := range barriers {
		flag = flag && achievabilityFunc(first, second, barrier)
	}
	return flag
}
