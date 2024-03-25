package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

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
	dead   bool
}

func (p *Player) IsAtTopOfScreen() bool {
	if p.y < (float32(WIDTH)/2) && p.dy < 0 {
		return true
	}
	return false
}

func (p *Player) Move() {
	// Gravity
	p.dy += 1
	p.y += p.dy
	p.x += p.dx

	// Link left and right of the screen
	if p.x < 0 {
		p.x = float32(WIDTH) + p.x
	} else if p.x > float32(WIDTH) {
		p.x = p.x - float32(WIDTH)
	}

	if p.y > float32(HEIGHT) {
		p.dead = true
	}
}

var DX_INCR float32 = 4
var DX_MAX float32 = 8

func (p *Player) CheckKeyPressed() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		if p.dx > -DX_MAX {
			p.dx -= DX_INCR
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		if p.dx < DX_MAX {
			p.dx += DX_INCR
		}
	} else {
		p.dx = 0
	}
}

func (p *Player) Jump() {
	p.dy = -20
}

func (p *Player) Update() {
	if !p.dead {
		p.CheckKeyPressed()
		p.CheckCollisions()
		p.Move()
	}

}

func (p *Player) Draw(screen *ebiten.Image) {
	if !p.dead {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(float64(p.scaleX), float64(p.scaleY))
		x := p.x - p.width/2
		y := p.y - p.height/2
		op.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(p.Img, op)
	}
}

// Collision box is 5px larger than the real platform
var MARGIN_X float32 = 5
var MARGIN_Y float32 = 5

func (p *Player) CheckCollisions() {
	for _, platform := range p.game.platforms {
		if p.CheckCollision(platform) {
			p.Jump()
		}
	}
}

func (p *Player) CheckCollision(platform *Platform) bool {
	if p.dy > 0 {
		if (p.x - p.width/2) > (platform.x + platform.width/2 - MARGIN_X) {
			return false
		} else if (p.x + p.width/2) < (platform.x - platform.width/2 + MARGIN_X) {
			return false
		} else if (p.y + p.height/2) < (platform.y - platform.height/2 - MARGIN_Y) {
			return false
		} else if (p.y - p.height/2) > (platform.y + platform.height/2 + MARGIN_Y) {
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
	p.y = float32(HEIGHT) / 2
	p.scaleX = 0.1
	p.scaleY = 0.1
	p.width = float32(p.Img.Bounds().Dx()) * p.scaleX
	p.height = float32(p.Img.Bounds().Dy()) * p.scaleY
	p.dx = 0
	p.dy = 0
	p.dead = false

	return p
}
