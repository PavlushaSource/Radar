package game

import (
	"bytes"
	"fmt"
	"github.com/PavlushaSource/Radar/resources"
	"github.com/PavlushaSource/Radar/view/config"
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

type Game struct {
	dogs  []*Dog
	pause bool

	bg *ebiten.Image

	borders    []Line
	camX, camY float64
	camScale   float64
	camScaleTo float64

	appConfig *config.ApplicationConfig
	cursor    Cursor
}

const (
	dogCount    = 100
	borderScale = 0.1
)

func NewGame() *Game {
	dogs := make([]*Dog, 0)

	for i := 0; i < dogCount; i++ {
		dogs = append(dogs, NewDog(dogImageFight))
	}

	bg := ebiten.NewImage(1920, 1080)
	bg.Fill(color.RGBA{R: 51, G: 232, B: 78, A: 0xff})

	return &Game{
		dogs:       dogs,
		camScale:   1,
		camScaleTo: 1,
		bg:         bg,
		borders:    make([]Line, 0),
	}
}

func (g *Game) Update() error {

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
	g.camScaleTo += scrollY * (g.camScaleTo / 7)

	// Clamp target zoom level.
	if g.camScaleTo < 0.01 {
		g.camScaleTo = 0.01
	} else if g.camScaleTo > 100 {
		g.camScaleTo = 100
	}

	// Smooth zoom transition.
	div := 10.0
	if g.camScaleTo > g.camScale {
		g.camScale += (g.camScaleTo - g.camScale) / div
	} else if g.camScaleTo < g.camScale {
		g.camScale -= (g.camScale - g.camScaleTo) / div
	}

	// Pan camera via keyboard.
	pan := 7.0 / g.camScale
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		g.camX -= pan
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		g.camX += pan
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		g.camY -= pan
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		g.camY += pan
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.camScaleTo = 1
		g.camScale = 1
		g.camX = 0
		g.camY = 0
	}

	var mx, my int = ebiten.CursorPosition()
	//mouseX
	mx -= int(g.camX * g.camScale)
	my += int(g.camY * g.camScale)
	//if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		fmt.Println(g.cursor.pressed)
		if !g.cursor.pressed {
			g.cursor.pressed = true
			g.cursor.startX, g.cursor.startY = mx, my
		}
		//bx, by := int(float64(borderImage.Bounds().Dx())*borderScale*g.camScale/1.75), int(float64(borderImage.Bounds().Dy())*borderScale*g.camScale/1.75)
		//if g.cursor.x-(bx) <= mx && g.cursor.x+(bx) >= mx && g.cursor.y-(by) <= my && g.cursor.y+by >= my {
		//} else {
		//	//fmt.Println("MOUSE POS", mx, my)
		//	g.cursor = Cursor{
		//		x:       mx,
		//		y:       my,
		//		pressed: true,
		//	}
		//}
		g.cursor.endX, g.cursor.endY = mx, my
	} else {
		g.cursor.pressed = false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.pause = !g.pause
	}

	if g.pause {
		return nil
	}
	for _, d := range g.dogs {
		d.Update()
	}
	return nil
}

func (g *Game) DrawBorder(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	scale := g.camScale

	for _, line := range g.borders {

		dx := float64(line.EndX - line.StartX)
		dy := float64(line.EndY - line.StartY)

		length := math.Hypot(dx, dy)

		dirX, dirY := dx/length, dy/length

		r := 16.0
		step := r * 1.5

		for i := 0.0; i < length; i += step {
			op.GeoM.Reset()

			op.GeoM.Translate(-160, -160)
			op.GeoM.Scale(0.1, 0.1)

			centerX := float64(line.StartX) + dirX*i
			centerY := float64(line.StartY) + dirY*i

			op.GeoM.Translate(centerX, centerY)

			op.GeoM.Translate(g.camX, -g.camY)
			op.GeoM.Scale(scale, scale)

			screen.DrawImage(borderImage, op)
			//g.bg.DrawImage(borderImage, op)
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
		//op.GeoM.Translate(-g.camX/scale, g.camY/scale)
		//op.GeoM.Scale(scale, scale)
	}

}

func (g *Game) Draw(screen *ebiten.Image) {
	//screen.Fill(color.RGBA{0xff, 0, 0, 0})
	//cx, cy := ScreenWidth/2, ScreenHeight/2
	target := screen
	scale := g.camScale
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(g.camX, -g.camY)
	op.GeoM.Scale(scale, scale)
	//op.GeoM.Translate(float64(cx), float64(cy))
	target.DrawImage(g.bg, op)

	if g.cursor.pressed {
		op.GeoM.Reset()
		op.GeoM.Translate(-160, -160)
		op.GeoM.Scale(0.1, 0.1)

		op.GeoM.Translate(float64(g.cursor.x)/scale, float64(g.cursor.y)/scale)
		//fmt.Println(scale)
		//op.GeoM.Translate(-g.camX/scale, g.camY/scale)
		//op.GeoM.Scale(scale, scale)

		//fmt.Println(g.cursor.startY, g.cursor.pressed)
		vector.StrokeLine(target, float32(g.cursor.startX), float32(g.cursor.startY), float32(g.cursor.endX), float32(g.cursor.endY), 3, color.RGBA{255, 0, 0, 64}, false)

		//g.bg.DrawImage(borderImage, op)
	} else {
		g.borders = append(g.borders, NewLine(g.cursor.startX, g.cursor.startY, g.cursor.endX, g.cursor.endY))

	}

	g.DrawBorder(target)

	for _, d := range g.dogs {

		op.GeoM.Reset()

		op.GeoM.Scale(d.imageScaleX, d.imageScaleY)
		op.GeoM.Translate(float64(d.x), float64(d.y))
		op.GeoM.Translate(g.camX, -g.camY)

		op.GeoM.Scale(scale, scale)

		//op.GeoM.Translate(float64(cx), float64(cy))
		//screen.DrawImage(dogImageFight, op)
		d.Draw(target, op)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("ACTUAL FPS %f\n", ebiten.ActualFPS()))

	//screen.DrawImage(target, op)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1920, 1080
}
