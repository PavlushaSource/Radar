package main

import (
	"bytes"
	"fmt"
	"image/color"
	"log"
	"strings"
	"unicode/utf8"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

// Game object used by ebiten
type game struct {
	ui           *ebitenui.UI
	textInputs   []*widget.TextInput // Хранение ссылок на текстовые поля
	updateButton *widget.Button      // Кнопка "Update Text"
}

type ListEntry struct {
	id   int
	name string
}

func main() {
	ebiten.SetWindowSize(400, 600)
	ebiten.SetWindowTitle("Ebiten UI - TextInput with Labels")

	buttonImage, _ := loadButtonImage()

	// Создаем объект игры
	game := &game{}

	// Загружаем шрифт
	face, _ := loadFont(20)

	// Создаем корневой контейнер с макетом GridLayout
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(2),                         // 2 колонки: для метки и текстового поля
			widget.GridLayoutOpts.Stretch([]bool{false, false}, nil), // Метка фиксирована, поле растягивается
			widget.GridLayoutOpts.Spacing(10, 10),
		)),
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0x13, 0x1a, 0x22, 0xff})),
	)

	// Функция для добавления метки и текстового поля
	addTextInputWithLabel := func(labelText string, placeholder string, validationFunc func(string) (bool, *string)) *widget.TextInput {
		// Метка
		label := widget.NewLabel(
			widget.LabelOpts.Text(labelText, face, &widget.LabelColor{
				Idle: color.White,
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

		// Текстовое поле
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
				Idle:          color.NRGBA{254, 255, 255, 255},
				Disabled:      color.NRGBA{R: 200, G: 200, B: 200, A: 255},
				Caret:         color.NRGBA{254, 255, 255, 255},
				DisabledCaret: color.NRGBA{R: 200, G: 200, B: 200, A: 255},
			}),
			widget.TextInputOpts.Padding(widget.NewInsetsSimple(5)),
			widget.TextInputOpts.CaretOpts(widget.CaretOpts.Size(face, 2)),
			widget.TextInputOpts.Placeholder(placeholder),
			widget.TextInputOpts.Validation(validationFunc),
			widget.TextInputOpts.ChangedHandler(func(args *widget.TextInputChangedEventArgs) {
				fmt.Println("Text Changed: ", args.InputText)
			}),
		)

		// Добавляем метку и поле в контейнер
		rootContainer.AddChild(label)
		rootContainer.AddChild(textInput)

		return textInput
	}

	// Добавляем поля с метками
	game.textInputs = append(game.textInputs, addTextInputWithLabel("Hello World", "Standard Textbox", nil))
	game.textInputs = append(game.textInputs, addTextInputWithLabel("Hello World", "Password Textbox", nil))
	game.textInputs = append(game.textInputs, addTextInputWithLabel("Hello World", "Max length (5)", func(newInputText string) (bool, *string) {
		if utf8.RuneCountInString(newInputText) > 5 {
			return false, nil
		}
		return true, nil
	}))
	game.textInputs = append(game.textInputs, addTextInputWithLabel("Hello World", "All Caps", func(newInputText string) (bool, *string) {
		newInputText = strings.ToUpper(newInputText)
		return false, &newInputText
	}))

	numEntries := 20
	entries := make([]any, 0, numEntries)
	for i := 1; i <= numEntries; i++ {
		entries = append(entries, ListEntry{i, fmt.Sprintf("Entry %d", i)})
	}
	// construct a combobox
	comboBox := widget.NewListComboButton(
		widget.ListComboButtonOpts.SelectComboButtonOpts(
			widget.SelectComboButtonOpts.ComboButtonOpts(
				//Set the max height of the dropdown list
				widget.ComboButtonOpts.MaxContentHeight(150),
				//Set the parameters for the primary displayed button
				widget.ComboButtonOpts.ButtonOpts(
					widget.ButtonOpts.Image(buttonImage),
					widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
					widget.ButtonOpts.Text("", face, &widget.ButtonTextColor{
						Idle:     color.White,
						Disabled: color.White,
					}),
					widget.ButtonOpts.WidgetOpts(
						//Set how wide the button should be
						widget.WidgetOpts.MinSize(150, 0),
						//Set the combobox's position
						widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
							HorizontalPosition: widget.AnchorLayoutPositionCenter,
							VerticalPosition:   widget.AnchorLayoutPositionCenter,
						})),
				),
			),
		),
		widget.ListComboButtonOpts.ListOpts(
			//Set how wide the dropdown list should be
			widget.ListOpts.ContainerOpts(widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.MinSize(150, 0))),
			//Set the entries in the list
			widget.ListOpts.Entries(entries),
			widget.ListOpts.ScrollContainerOpts(
				//Set the background images/color for the dropdown list
				widget.ScrollContainerOpts.Image(&widget.ScrollContainerImage{
					Idle:     image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
					Disabled: image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
					Mask:     image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
				}),
			),
			widget.ListOpts.SliderOpts(
				//Set the background images/color for the background of the slider track
				widget.SliderOpts.Images(&widget.SliderTrackImage{
					Idle:  image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
					Hover: image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
				}, buttonImage),
				widget.SliderOpts.MinHandleSize(5),
				//Set how wide the track should be
				widget.SliderOpts.TrackPadding(widget.NewInsetsSimple(2))),
			//Set the font for the list options
			widget.ListOpts.EntryFontFace(face),
			//Set the colors for the list
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
			//Padding for each entry
			widget.ListOpts.EntryTextPadding(widget.NewInsetsSimple(5)),
		),
		//Define how the entry is displayed
		widget.ListComboButtonOpts.EntryLabelFunc(
			func(e any) string {
				//Button Label function
				return e.(ListEntry).name
			},
			func(e any) string {
				//List Label function
				return "List: " + e.(ListEntry).name
			}),
		//Callback when a new entry is selected
		widget.ListComboButtonOpts.EntrySelectedHandler(func(args *widget.ListComboButtonEntrySelectedEventArgs) {
			fmt.Println("Selected Entry: ", args.Entry)
		}),
	)
	label := widget.NewLabel(
		widget.LabelOpts.Text("DISTANC", face, &widget.LabelColor{
			Idle: color.White,
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

	rootContainer.AddChild(label)
	rootContainer.AddChild(comboBox)

	// Создаем кнопку "Update Text"
	game.updateButton = widget.NewButton(
		widget.ButtonOpts.Text("Update Text", face, &widget.ButtonTextColor{
			Idle: color.NRGBA{0xdf, 0xf4, 0xff, 0xff},
		}),
		widget.ButtonOpts.Image(&widget.ButtonImage{
			Idle:     image.NewNineSliceColor(color.NRGBA{R: 100, G: 150, B: 200, A: 255}),
			Hover:    image.NewNineSliceColor(color.NRGBA{R: 120, G: 170, B: 220, A: 255}),
			Pressed:  image.NewNineSliceColor(color.NRGBA{R: 80, G: 120, B: 180, A: 255}),
			Disabled: image.NewNineSliceColor(color.NRGBA{R: 50, G: 50, B: 50, A: 255}),
		}),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			// Сброс всех текстовых полей
			for _, input := range game.textInputs {
				input.SetText("")
			}
		}),
	)

	runButton := widget.NewButton(
		widget.ButtonOpts.Text("Run", face, &widget.ButtonTextColor{
			Idle: color.NRGBA{0xdf, 0xf4, 0xff, 0xff},
		}),
		widget.ButtonOpts.Image(&widget.ButtonImage{
			Idle:     image.NewNineSliceColor(color.NRGBA{R: 100, G: 150, B: 200, A: 255}),
			Hover:    image.NewNineSliceColor(color.NRGBA{R: 120, G: 170, B: 220, A: 255}),
			Pressed:  image.NewNineSliceColor(color.NRGBA{R: 80, G: 120, B: 180, A: 255}),
			Disabled: image.NewNineSliceColor(color.NRGBA{R: 50, G: 50, B: 50, A: 255}),
		}),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			// Сброс всех текстовых полей
			for _, input := range game.textInputs {
				input.SetText("")
			}
		}),
	)

	// Добавляем кнопку в правый нижний угол
	rootContainer.AddChild(game.updateButton)
	rootContainer.AddChild(runButton)

	// Создаем UI
	ui := ebitenui.UI{
		Container: rootContainer,
	}
	game.ui = &ui

	// Запускаем игру
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
	// Обновление UI
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		g.ui.ChangeFocus(widget.FOCUS_PREVIOUS)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		g.ui.ChangeFocus(widget.FOCUS_NEXT)
	}
	g.ui.Update()
	return nil
}

// Draw implements Ebiten's Draw method.
func (g *game) Draw(screen *ebiten.Image) {
	// Рисуем UI
	g.ui.Draw(screen)
}

// Загрузка шрифта
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
