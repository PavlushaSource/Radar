package main

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
	"image/color"
	"log"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

// Game object used by ebiten
type game struct {
	ui *ebitenui.UI
}

/*
The Anchor layout is used to specify where to draw the elements within the container.
This layout does not account for overlap as seen in this example where we have two widgets set to be centered in the parent.
This is often used to align a widget or another container to the center of the parent container.
*/
func main() {
	face, _ := loadFont(20)
	// construct a new container that serves as the root of the UI hierarchy
	rootContainer := widget.NewContainer(
		// the container will use a plain color as its background
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0x13, 0x1a, 0x22, 0xff})),
		// the container will use an anchor layout to layout its single child widget
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(
			//Set how much padding before displaying content
			widget.AnchorLayoutOpts.Padding(widget.NewInsetsSimple(30)),
		)),
	)

	innerContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{255, 0, 0, 255})),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
				StretchHorizontal:  true,
				StretchVertical:    false,
			}),
			widget.WidgetOpts.MinSize(100, 100),
		),
	)

	label := widget.NewLabel(
		widget.LabelOpts.Text("DISTANC", face, &widget.LabelColor{
			Idle: color.Black,
		}),
		widget.LabelOpts.TextOpts(
			widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
			widget.TextOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(widget.RowLayoutData{
					Position: widget.RowLayoutPositionCenter,
				}),
			),
		),
	)
	rootContainer.AddChild(innerContainer)
	rootContainer.AddChild(label)

	//innerContainer2 := widget.NewContainer(
	//	widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0, 200, 255, 255})),
	//	widget.ContainerOpts.WidgetOpts(
	//		widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
	//			HorizontalPosition: widget.AnchorLayoutPositionCenter,
	//			VerticalPosition:   widget.AnchorLayoutPositionCenter,
	//			StretchHorizontal:  false,
	//			StretchVertical:    true,
	//		}),
	//		widget.WidgetOpts.MinSize(100, 100),
	//	),
	//)
	//rootContainer.AddChild(innerContainer2)
	// construct the UI
	ui := ebitenui.UI{
		Container: rootContainer,
	}

	// Ebiten setup
	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowTitle("Ebiten UI - Anchor Layout")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	game := game{
		ui: &ui,
	}

	// run Ebiten main loop
	err := ebiten.RunGame(&game)
	if err != nil {
		log.Println(err)
	}
}

// Layout implements Game.
func (g *game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

// Update implements Game.
func (g *game) Update() error {
	// update the UI
	g.ui.Update()
	return nil
}

// Draw implements Ebiten's Draw method.
func (g *game) Draw(screen *ebiten.Image) {
	// draw the UI onto the screen
	g.ui.Draw(screen)
}

func loadFont(size float64) (text.Face, error) {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &text.GoTextFace{
		Source: s,
		Size:   size,
	}, nil
}
