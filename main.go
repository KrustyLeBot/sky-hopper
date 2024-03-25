package main

import (
	"fmt"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var WIDTH int = 550
var HEIGHT int = 900

type Game struct {
	background *ebiten.Image
	init       bool
	player     Player
	platforms  []*Platform
}

func UpdateTitle(g *Game) {
	ebiten.SetWindowTitle(fmt.Sprint("Sky Hopper: Ascension Infinie      TPS: ", int(ebiten.ActualTPS())))
}

func (g *Game) Update() error {
	if !g.init {
		g.init = true
	}
	UpdateTitle(g)
	g.player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawBackground(g, screen)
	g.player.Draw(screen)

	for _, platform := range g.platforms {
		platform.Draw(screen)
	}
}

func DrawBackground(g *Game, screen *ebiten.Image) {
	screen.DrawImage(g.background, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIDTH, HEIGHT
}

func NewGame() *Game {
	g := &Game{}
	g.init = false

	img, _, err := ebitenutil.NewImageFromFile("assets/background.png")
	if err != nil {
		log.Fatal(err)
	}
	g.background = img
	g.player = *NewPlayer(g)

	g.platforms = append(g.platforms, NewPlatform())

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
