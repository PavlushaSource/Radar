package utils

import (
	"fyne.io/fyne/v2"
	"github.com/PavlushaSource/Radar/view/domain"
)

func resourceCat(color domain.Color) fyne.Resource {
	switch color {
	case domain.Red:
		return ResourceCatRedSvg
	case domain.Purple:
		return ResourceCatPurpleSvg
	case domain.Blue:
		return ResourceCatBlueSvg
	default:
		panic("resourceCat: unknown color")
	}
}
