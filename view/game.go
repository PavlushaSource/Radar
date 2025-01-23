package view

import (
	"bytes"
	"fmt"
	"github.com/PavlushaSource/Radar/resources"
	"github.com/PavlushaSource/Radar/view/config"
	"github.com/PavlushaSource/Radar/view/utils"
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image"
	"image/color"
	"log"
	"math"
)

var (
	dogImageFight *ebiten.Image
	dogImageHiss  *ebiten.Image
	borderImage   *ebiten.Image
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(resources.DogFight_png))
	if err != nil {
		log.Fatal(err)
	}
	dogImageFight = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(resources.DogHiss_png))
	if err != nil {
		log.Fatal(err)
	}
	dogImageHiss = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(resources.Border_png))
	if err != nil {
		log.Fatal(err)
	}
	borderImage = ebiten.NewImageFromImage(img)
}

type Application struct {
	dogs  []*Dog
	pause bool

	bg *ebiten.Image

	borders []utils.Line

	appConfig *config.ApplicationConfig

	cursor utils.Cursor

	Menu         *ebitenui.UI
	textInputs   []*widget.TextInput // Хранение ссылок на текстовые поля
	updateButton *widget.Button      // Кнопка "Update Text"

	InMainMenu bool
}

func NewApplication() *Application {
	dogs := make([]*Dog, 0)

	for i := 0; i < dogCount; i++ {
		dogs = append(dogs, NewDog(dogImageFight))
	}

	bg := ebiten.NewImage(1920, 1080)
	bg.Fill(color.RGBA{R: 51, G: 232, B: 78, A: 0xff})

	game := &Application{}
	game.bg = bg
	game.borders = make([]utils.Line, 0)
	game.dogs = dogs
	game.Menu = NewUI(game)
	game.appConfig = config.NewApplicationConfig()

	return game
}

func (a *Application) Update() error {

	var scrollY float64
	if ebiten.IsKeyPressed(ebiten.KeyC) || ebiten.IsKeyPressed(ebiten.KeyPageDown) {
		scrollY = -0.25
	} else if ebiten.IsKeyPressed(ebiten.KeyE) || ebiten.IsKeyPressed(ebiten.KeyPageUp) {
		scrollY = 0.25
	} else {
		_, scrollY = ebiten.Wheel()
		if scrollY < -1 {
			scrollY = -1
		} else if scrollY > 1 {
			scrollY = 1
		}
	}
	a.appConfig.CamScaleTo += scrollY * (a.appConfig.CamScaleTo / 7)

	// Clamp target zoom level.
	if a.appConfig.CamScaleTo < 0.01 {
		a.appConfig.CamScaleTo = 0.01
	} else if a.appConfig.CamScaleTo > 100 {
		a.appConfig.CamScaleTo = 100
	}

	// Smooth zoom transition.
	div := 10.0
	if a.appConfig.CamScaleTo > a.appConfig.CamScale {
		a.appConfig.CamScale += (a.appConfig.CamScaleTo - a.appConfig.CamScale) / div
	} else if a.appConfig.CamScaleTo < a.appConfig.CamScale {
		a.appConfig.CamScale -= (a.appConfig.CamScale - a.appConfig.CamScaleTo) / div
	}

	// Pan camera via keyboard.
	pan := 7.0 / a.appConfig.CamScale
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		a.appConfig.CamX -= pan
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		a.appConfig.CamX += pan
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		a.appConfig.CamY -= pan
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		a.appConfig.CamY += pan
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		a.appConfig.CamScaleTo = 1
		a.appConfig.CamScale = 1
		a.appConfig.CamX = 0
		a.appConfig.CamY = 0
	}

	var mx, my int = ebiten.CursorPosition()
	//mouseX
	mx -= int(a.appConfig.CamX * a.appConfig.CamScale)
	my += int(a.appConfig.CamY * a.appConfig.CamScale)
	//if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		fmt.Println(a.cursor.Pressed)
		if !a.cursor.Pressed {
			a.cursor.Pressed = true
			a.cursor.StartX, a.cursor.StartY = mx, my
		}
		//bx, by := int(float64(borderImage.Bounds().Dx())*borderImgScale*a.camScale/1.75), int(float64(borderImage.Bounds().Dy())*borderImgScale*a.camScale/1.75)
		//if a.cursor.x-(bx) <= mx && a.cursor.x+(bx) >= mx && a.cursor.y-(by) <= my && a.cursor.y+by >= my {
		//} else {
		//	//fmt.Println("MOUSE POS", mx, my)
		//	a.cursor = Cursor{
		//		x:       mx,
		//		y:       my,
		//		pressed: true,
		//	}
		//}
		a.cursor.EndX, a.cursor.EndY = mx, my
	} else {
		a.cursor.Pressed = false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		a.pause = !a.pause
	}

	if a.pause {
		return nil
	}
	for _, d := range a.dogs {
		d.Update()
	}
	return nil
}

func (a *Application) DrawBorder(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	scale := a.appConfig.CamScale

	for _, line := range a.borders {

		dx := float64(line.EndX - line.StartX)
		dy := float64(line.EndY - line.StartY)

		length := math.Hypot(dx, dy)

		dirX, dirY := dx/length, dy/length

		r := 16.0
		step := r * 1.5

		for i := 0.0; i < length; i += step {
			op.GeoM.Reset()

			op.GeoM.Translate(-160, -160)
			op.GeoM.Scale(borderImgScale, borderImgScale)

			centerX := float64(line.StartX) + dirX*i
			centerY := float64(line.StartY) + dirY*i

			op.GeoM.Translate(centerX, centerY)

			op.GeoM.Translate(a.appConfig.CamX, -a.appConfig.CamY)
			op.GeoM.Scale(scale, scale)

			screen.DrawImage(borderImage, op)
			//a.bg.DrawImage(borderImage, op)
			//vector.DrawFilledCircle(screen, float32(centerX), float32(centerY), float32(r), color.RGBA{0, 255, 0, 255}, false)
		}
		//currPosX := line.StartX
		//currPosY := line.StartY
		//op.GeoM.Reset()
		//op.GeoM.Translate(-160, -160)
		//op.GeoM.Scale(0.1, 0.1)
		//
		//op.GeoM.Translate(float64(currPosX)/scale, float64(currPosY)/scale)
		//
		//op.GeoM.Translate(-a.camX/scale, a.camY/scale)
		//op.GeoM.Scale(scale, scale)
	}

}

func (a *Application) Draw(screen *ebiten.Image) {
	//screen.Fill(color.RGBA{0xff, 0, 0, 0})
	//cx, cy := ScreenWidth/2, ScreenHeight/2
	target := screen
	scale := a.appConfig.CamScale
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(a.appConfig.CamX, -a.appConfig.CamY)
	op.GeoM.Scale(scale, scale)
	//op.GeoM.Translate(float64(cx), float64(cy))
	target.DrawImage(a.bg, op)

	if a.cursor.Pressed {
		op.GeoM.Reset()
		op.GeoM.Translate(-160, -160)
		op.GeoM.Scale(0.1, 0.1)

		op.GeoM.Translate(float64(a.cursor.X)/scale, float64(a.cursor.Y)/scale)
		//fmt.Println(scale)
		//op.GeoM.Translate(-a.camX/scale, a.camY/scale)
		//op.GeoM.Scale(scale, scale)

		//fmt.Println(a.cursor.startY, a.cursor.pressed)
		vector.StrokeLine(target, float32(a.cursor.StartX), float32(a.cursor.StartY), float32(a.cursor.EndX), float32(a.cursor.EndY), 3, color.RGBA{255, 0, 0, 64}, false)

		//a.bg.DrawImage(borderImage, op)
	} else {
		a.borders = append(a.borders, utils.NewLine(a.cursor.StartX, a.cursor.StartY, a.cursor.EndX, a.cursor.EndY))

	}

	a.DrawBorder(target)

	for _, d := range a.dogs {

		op.GeoM.Reset()

		op.GeoM.Scale(dogImgScale, dogImgScale)
		op.GeoM.Translate(float64(d.x), float64(d.y))
		op.GeoM.Translate(a.appConfig.CamX, -a.appConfig.CamY)

		op.GeoM.Scale(scale, scale)

		//op.GeoM.Translate(float64(cx), float64(cy))
		//screen.DrawImage(dogImageFight, op)
		d.Draw(target, op)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("ACTUAL FPS %f\n", ebiten.ActualFPS()))

	//screen.DrawImage(target, op)

}

func (a *Application) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return a.appConfig.WindowX, a.appConfig.WindowY
}
