package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var MARGIN float32 = 5

type Player struct {
	game   *Game
	Img    *ebiten.Image
	x      float32
	y      float32
	scaleX float32
	scaleY float32
	width  float32
	height float32
	dx     float32
	dy     float32
}

func (p *Player) Move() {
	// Gravity
	p.dy += 1

	p.x += p.dx
	p.y += p.dy
}

func (p *Player) Jump() {
	p.dy -= 20
}

func (p *Player) Update() {

	for _, platform := range p.game.platforms {
		if p.CheckCollision(platform) {
			p.Jump()
		}
	}
	p.Move()
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(float64(p.scaleX), float64(p.scaleY))
	x := p.x - p.width/2
	y := p.y - p.height/2
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(p.Img, op)
}

func (p *Player) CheckCollision(platform *Platform) bool {
	if p.dy > 0 {
		if (p.x - p.width/2) > (platform.x + platform.width/2 - MARGIN) {
			return false
		} else if (p.x + p.width/2) < (platform.x - platform.width/2 + MARGIN) {
			return false
		} else if (p.y + p.height/2 + MARGIN) < (platform.y - platform.height/2) {
			return false
		} else if (p.y - p.height/2 - MARGIN) > (platform.y + platform.height/2) {
			return false
		}
		return true
	}
	return false
}

func NewPlayer(game *Game) *Player {
	p := &Player{}

	img, _, err := ebitenutil.NewImageFromFile("assets/character.png")
	if err != nil {
		log.Fatal(err)
	}

	p.game = game
	p.Img = img
	p.x = float32(WIDTH) / 2
	p.y = 70
	p.scaleX = 0.1
	p.scaleY = 0.1
	p.width = float32(p.Img.Bounds().Dx()) * p.scaleX
	p.height = float32(p.Img.Bounds().Dy()) * p.scaleY
	p.dx = 0
	p.dy = 0

	return p
}
