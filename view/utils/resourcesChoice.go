package utils

import (
	"fyne.io/fyne/v2"
	"github.com/PavlushaSource/Radar/viewModel"
)

func resourceCat(color producer.Color) fyne.Resource {
	switch color {
	case producer.Red:
		return ResourceCatRedSvg
	case producer.Purple:
		return ResourceCatPurpleSvg
	case producer.Blue:
		return ResourceCatBlueSvg
	default:
		panic("resourceCat: unknown color")
	}
}
