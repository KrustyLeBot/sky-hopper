package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var WIDTH int = 550
var HEIGHT int = 900

type Game struct {
	scaleX     float64
	scaleY     float64
	loss_img   *ebiten.Image
	background *ebiten.Image
	init       bool
	player     Player
	platforms  []*Platform
	score      int
	font       font.Face
}

func UpdateTitle(g *Game) {
	ebiten.SetWindowTitle("Sky Hopper: Ascension Infinie")
}

func (g *Game) Update() error {
	if !g.init {
		g.init = true
	}
	UpdateTitle(g)
	g.player.Update()

	if g.player.IsAtTopOfScreen() {
		for _, platform := range g.platforms {
			if platform.Move(g.player.dy) {
				g.score += 1
			}
		}
	}

	if g.player.dead {
		g.CheckRestart()
	}

	return nil
}

func (g *Game) CheckRestart() {
	if ebiten.IsKeyPressed(ebiten.KeyF1) {
		g.player = *NewPlayer(g)
		g.CreateDefaultPlatforms()
		g.score = 0
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawBackground(screen)
	g.DrawScore(screen)
	g.player.Draw(screen)

	if g.player.dead {
		g.DrawLoss(screen)
	} else {
		for _, platform := range g.platforms {
			platform.Draw(screen)
		}
	}
}

func (g *Game) DrawBackground(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(g.scaleX), float64(g.scaleY))
	screen.DrawImage(g.background, op)
}

func (g *Game) DrawScore(screen *ebiten.Image) {
	text.Draw(screen, fmt.Sprint("SCORE: ", g.score), g.font, WIDTH/3, 38, color.White)
}

func (g *Game) DrawLoss(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(g.scaleX), float64(g.scaleY))
	screen.DrawImage(g.loss_img, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIDTH, HEIGHT
}

func (g *Game) CreateDefaultPlatforms() {
	g.platforms = nil
	g.platforms = append(g.platforms, NewPlatform(float32(WIDTH)/2, float32(HEIGHT)-30))
	g.platforms = append(g.platforms, NewPlatform(float32(WIDTH)/3*2, float32(HEIGHT)/5))
	g.platforms = append(g.platforms, NewPlatform(float32(WIDTH)/4, float32(HEIGHT)/5*2))
	g.platforms = append(g.platforms, NewPlatform(float32(WIDTH)/3, float32(HEIGHT)/5*3))
	g.platforms = append(g.platforms, NewPlatform(float32(WIDTH)/4*3, float32(HEIGHT)/5*4))
}

func NewGame() *Game {
	g := &Game{}
	g.init = false

	img, _, err := ebitenutil.NewImageFromFile("assets/background.png")
	if err != nil {
		log.Fatal(err)
	}
	g.background = img

	g.scaleX = float64(WIDTH) / float64(img.Bounds().Dx())
	g.scaleY = float64(HEIGHT) / float64(img.Bounds().Dy())

	img, _, err = ebitenutil.NewImageFromFile("assets/loss.png")
	if err != nil {
		log.Fatal(err)
	}
	g.loss_img = img

	g.player = *NewPlayer(g)
	g.CreateDefaultPlatforms()

	g.score = 0

	// Load Font
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    36,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
	g.font = mplusNormalFont

	return g
}

func main() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)
	ebiten.SetWindowSize(WIDTH, HEIGHT)

	g := NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
