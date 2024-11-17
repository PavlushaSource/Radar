package geom

type Distance func(first Point, second Point, barriers []Barrier) float64

type achievability func(first Point, second Point, barrier Barrier) bool

func arePointsAchievable(first Point, second Point, barriers []Barrier, achievabilityFunc achievability) bool {
	var flag = true
	for _, barrier := range barriers {
		flag = flag && achievabilityFunc(first, second, barrier)
	}
	return flag
}
