package viewModel

import (
	"github.com/PavlushaSource/Radar/model/engine"
)

func ConvertStateToVMCat(state engine.State) []Cat {
	vmCats := make([]Cat, 0, state.NumCats())

	for _, c := range state.Cats() {
		//fmt.Println(c.Status())
		//fmt.Println(ConvertStatusToColor(c))
		vmCats = append(vmCats, Cat{X: float32(c.X()), Y: float32(c.Y()), Color: ConvertStatusToColor(c)})
	}

	return vmCats
}

func ConvertStatusToColor(cat engine.Cat) Color {
	switch cat.Status() {
	case engine.Calm:
		return Blue
	case engine.Hissing:
		return Purple
	case engine.Fighting:
		return Red
	default:
		panic("Undefined Color")
	}
}
