//go:generate fyne bundle -o svg.go ../../svg

package runner

import "math/rand/v2"

func GenerateBackendCats(count int) []CatBackend {
	res := make([]CatBackend, count)

	minX, maxX := 0, 1920
	minY, maxY := 0, 1080

	for i := 0; i < count; i++ {
		currX := rand.IntN(maxX-minX) + minX
		currY := rand.IntN(maxY-minY) + minY

		res[i] = CatBackend{X: float32(currX), Y: float32(currY)}
		res = append(res, CatBackend{X: float32(currX), Y: float32(currY)})
	}
	return res
}
