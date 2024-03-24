package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var WIDTH int = 630
var HEIGHT int = 1024

type Game struct {
	Background *ebiten.Image
	Init       bool
}

func (g *Game) Update() error {
	if !g.Init {
		g.Init = true
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.Background, nil)
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIDTH, HEIGHT
}

func NewGame() *Game {
	g := &Game{}
	g.Init = false

	img, _, err := ebitenutil.NewImageFromFile("background.png")
	if err != nil {
		log.Fatal(err)
	}
	g.Background = img

	return g
}

func main() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)
	ebiten.SetWindowSize(WIDTH, HEIGHT)
	ebiten.SetWindowTitle("Sky Hopper: Ascension Infinie")

	g := NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
