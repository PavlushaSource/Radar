package utils

import (
	"fyne.io/fyne/v2"
	"github.com/PavlushaSource/Radar/viewModel"
)

func ResourceCat(color viewModel.Color) fyne.Resource {
	switch color {
	case viewModel.Red:
		return ResourceCatRedSvg
	case viewModel.Purple:
		return ResourceCatPurpleSvg
	case viewModel.Blue:
		return ResourceCatBlueSvg
	default:
		panic("ResourceCat: unknown color")
	}
}
