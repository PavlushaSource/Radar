package main

import (
	"bytes"
	"fmt"
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
	"image/color"
	"log"
)

type game struct {
	ui           *ebitenui.UI
	textInputs   []*widget.TextInput // Хранение ссылок на текстовые поля
	updateButton *widget.Button      // Кнопка "Update Text"
	inMainMenu   bool
}

var face, _ = loadFont(20)

var InputFields = make(map[string]*widget.TextInput)

func NewLabel(labelText string) *widget.Label {
	return widget.NewLabel(
		widget.LabelOpts.Text(labelText, face, &widget.LabelColor{
			Idle: color.White,
		}),
		widget.LabelOpts.TextOpts(
			widget.TextOpts.Position(widget.TextPositionEnd, widget.TextPositionEnd),
			widget.TextOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(widget.RowLayoutData{
					Position: widget.RowLayoutPositionEnd,
				}),
			),
		),
	)
}

func AddInputFieldWithText(container *widget.Container, labelText string, placeholder string) {
	label := NewLabel(labelText)

	textInput := widget.NewTextInput(
		widget.TextInputOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(200, 30),
		),
		widget.TextInputOpts.Image(&widget.TextInputImage{
			Idle:     image.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 100, A: 255}),
			Disabled: image.NewNineSliceColor(color.NRGBA{R: 50, G: 50, B: 50, A: 255}),
		}),
		widget.TextInputOpts.Face(face),
		widget.TextInputOpts.Color(&widget.TextInputColor{
			Idle:          color.NRGBA{R: 254, G: 255, B: 255, A: 255},
			Disabled:      color.NRGBA{R: 200, G: 200, B: 200, A: 255},
			Caret:         color.NRGBA{R: 254, G: 255, B: 255, A: 255},
			DisabledCaret: color.NRGBA{R: 200, G: 200, B: 200, A: 255},
		}),
		widget.TextInputOpts.Padding(widget.NewInsetsSimple(5)),
		widget.TextInputOpts.CaretOpts(widget.CaretOpts.Size(face, 2)),
		widget.TextInputOpts.Placeholder(placeholder),
	)

	InputFields[labelText] = textInput
	container.AddChild(label)
	container.AddChild(textInput)

}

func NewCenterContainer() *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{R: 0x13, G: 0x1a, B: 0x22, A: 0xff})),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(
			widget.AnchorLayoutOpts.Padding(widget.NewInsetsSimple(30)),
		)),
	)
}

func NewGridContainer() *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(2),                         // 2 колонки: для метки и текстового поля
			widget.GridLayoutOpts.Stretch([]bool{false, false}, nil), // Метка фиксирована, поле растягивается
			widget.GridLayoutOpts.Spacing(10, 10),
		)),
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{R: 0x13, G: 0x1a, B: 0x22, A: 0xff})),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
				StretchHorizontal:  true,
				StretchVertical:    false,
			})),
	)
}

func AddComboboxFieldWithText(container *widget.Container, labelText string, entries []any,
	selectHandler func(args *widget.ListComboButtonEntrySelectedEventArgs)) {
	buttonImage, _ := loadButtonImage()

	comboBox := widget.NewListComboButton(
		widget.ListComboButtonOpts.SelectComboButtonOpts(
			widget.SelectComboButtonOpts.ComboButtonOpts(
				widget.ComboButtonOpts.MaxContentHeight(150),
				widget.ComboButtonOpts.ButtonOpts(
					widget.ButtonOpts.Image(buttonImage),
					widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
					widget.ButtonOpts.Text("", face, &widget.ButtonTextColor{
						Idle:     color.White,
						Disabled: color.White,
					}),
					widget.ButtonOpts.WidgetOpts(
						widget.WidgetOpts.MinSize(150, 0),
					),
				),
			),
		),
		widget.ListComboButtonOpts.ListOpts(
			widget.ListOpts.ContainerOpts(widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.MinSize(150, 0))),
			widget.ListOpts.Entries(entries),
			widget.ListOpts.ScrollContainerOpts(
				widget.ScrollContainerOpts.Image(&widget.ScrollContainerImage{
					Idle:     image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
					Disabled: image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
					Mask:     image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
				}),
			),
			widget.ListOpts.SliderOpts(
				widget.SliderOpts.Images(&widget.SliderTrackImage{
					Idle:  image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
					Hover: image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
				}, buttonImage),
				widget.SliderOpts.MinHandleSize(5),
				widget.SliderOpts.TrackPadding(widget.NewInsetsSimple(2))),
			widget.ListOpts.EntryFontFace(face),
			widget.ListOpts.EntryColor(&widget.ListEntryColor{
				Selected:                   color.NRGBA{254, 255, 255, 255},             //Foreground color for the unfocused selected entry
				Unselected:                 color.NRGBA{254, 255, 255, 255},             //Foreground color for the unfocused unselected entry
				SelectedBackground:         color.NRGBA{R: 130, G: 130, B: 200, A: 255}, //Background color for the unfocused selected entry
				SelectedFocusedBackground:  color.NRGBA{R: 130, G: 130, B: 170, A: 255}, //Background color for the focused selected entry
				FocusedBackground:          color.NRGBA{R: 170, G: 170, B: 180, A: 255}, //Background color for the focused unselected entry
				DisabledUnselected:         color.NRGBA{100, 100, 100, 255},             //Foreground color for the disabled unselected entry
				DisabledSelected:           color.NRGBA{100, 100, 100, 255},             //Foreground color for the disabled selected entry
				DisabledSelectedBackground: color.NRGBA{100, 100, 100, 255},             //Background color for the disabled selected entry
			}),
			widget.ListOpts.EntryTextPadding(widget.NewInsetsSimple(5)),
		),
		widget.ListComboButtonOpts.EntryLabelFunc(
			func(e any) string {
				return e.(string)
			},
			func(e any) string {
				return e.(string)
			}),
		widget.ListComboButtonOpts.EntrySelectedHandler(selectHandler),
	)

	label := NewLabel(labelText)

	container.AddChild(label)
	container.AddChild(comboBox)
}

func NewButton(buttonText string, clickHandler func(args *widget.ButtonClickedEventArgs)) *widget.Button {
	return widget.NewButton(
		widget.ButtonOpts.Text(buttonText, face, &widget.ButtonTextColor{
			Idle: color.NRGBA{R: 0xdf, G: 0xf4, B: 0xff, A: 0xff},
		}),
		widget.ButtonOpts.Image(&widget.ButtonImage{
			Idle:     image.NewNineSliceColor(color.NRGBA{R: 100, G: 150, B: 200, A: 255}),
			Hover:    image.NewNineSliceColor(color.NRGBA{R: 120, G: 170, B: 220, A: 255}),
			Pressed:  image.NewNineSliceColor(color.NRGBA{R: 80, G: 120, B: 180, A: 255}),
			Disabled: image.NewNineSliceColor(color.NRGBA{R: 50, G: 50, B: 50, A: 255}),
		}),
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(150, 35),
		),
		widget.ButtonOpts.ClickedHandler(clickHandler),
	)
}

func main() {

	ebiten.SetWindowSize(400, 400)
	ebiten.SetWindowTitle("Ebiten UI - TextInput with Labels")

	game := &game{inMainMenu: true}

	rootContainer := NewCenterContainer()
	innerContainer := NewGridContainer()

	AddInputFieldWithText(innerContainer, "Count cats", "integer number")
	AddInputFieldWithText(innerContainer, "Update time", "time in second")
	AddInputFieldWithText(innerContainer, "Angry radius", "in pixels")
	AddInputFieldWithText(innerContainer, "Hissing radius", "in pixels")

	AddComboboxFieldWithText(innerContainer, "Distance type", []any{"Euclidean", "Manhattan", "Curvilinear"}, func(args *widget.ListComboButtonEntrySelectedEventArgs) {
		fmt.Println(args.Entry)
	})

	AddComboboxFieldWithText(innerContainer, "Geometry type", []any{"Simple", "Vector"}, func(args *widget.ListComboButtonEntrySelectedEventArgs) {
		fmt.Println(args.Entry)
	})

	resetButton := NewButton("Reset", func(args *widget.ButtonClickedEventArgs) {
		for _, input := range InputFields {
			input.SetText("")
		}
	})

	runButton := NewButton("Run", func(args *widget.ButtonClickedEventArgs) {
		game.inMainMenu = false
		ebiten.SetFullscreen(true)

		// TODO validate appConfig HERE
	})

	innerContainer.AddChild(resetButton)
	innerContainer.AddChild(runButton)

	rootContainer.AddChild(innerContainer)

	ui := ebitenui.UI{
		Container: rootContainer,
	}
	game.ui = &ui

	err := ebiten.RunGame(game)
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

	if g.inMainMenu {
		// Обновление UI
		if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			g.ui.ChangeFocus(widget.FOCUS_PREVIOUS)
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			g.ui.ChangeFocus(widget.FOCUS_NEXT)
		}
		g.ui.Update()
	} else {
		//g.ui.Update()
	}
	return nil
}

// Draw implements Ebiten's Draw method.
func (g *game) Draw(screen *ebiten.Image) {
	// Рисуем UI
	//fmt.Println(g.inMainMenu)
	if g.inMainMenu {
		g.ui.Draw(screen)
	} else {
		screen.Fill(color.NRGBA{R: 100, G: 150, B: 200, A: 255})
	}
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

func loadButtonImage() (*widget.ButtonImage, error) {
	idle := image.NewNineSliceColor(color.NRGBA{R: 170, G: 170, B: 180, A: 255})

	hover := image.NewNineSliceColor(color.NRGBA{R: 130, G: 130, B: 150, A: 255})

	pressed := image.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 120, A: 255})

	return &widget.ButtonImage{
		Idle:    idle,
		Hover:   hover,
		Pressed: pressed,
	}, nil
}
