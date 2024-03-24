package main

import (
	"fmt"
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

func UpdateTitle(g *Game) {
	ebiten.SetWindowTitle(fmt.Sprint("Sky Hopper: Ascension Infinie      TPS: ", int(ebiten.ActualTPS())))
}

func (g *Game) Update() error {
	if !g.Init {
		g.Init = true
	}
	UpdateTitle(g)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawBackground(g, screen)
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func DrawBackground(g *Game, screen *ebiten.Image) {
	screen.DrawImage(g.Background, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIDTH, HEIGHT
}

func NewGame() *Game {
	g := &Game{}
	g.Init = false

	img, _, err := ebitenutil.NewImageFromFile("Assets/background.png")
	if err != nil {
		log.Fatal(err)
	}
	g.Background = img

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
