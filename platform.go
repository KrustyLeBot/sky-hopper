package main

import (
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Platform struct {
	Img    *ebiten.Image
	x      float32
	y      float32
	scaleX float32
	scaleY float32
	width  float32
	height float32
}

func (platform *Platform) Move(dy float32) bool {
	platform.y -= dy

	// Link top and bottom of the screen
	if platform.y > float32(HEIGHT) {
		platform.x = rand.Float32() * (float32(WIDTH))
		platform.y = rand.Float32() * (platform.y - float32(HEIGHT))
		return true
	}

	return false
}

func (platform *Platform) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(platform.scaleX), float64(platform.scaleY))
	x := platform.x - platform.width/2
	y := platform.y - platform.height/2
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(platform.Img, op)
}

func NewPlatform(x float32, y float32) *Platform {
	platform := &Platform{}

	platform.x = x
	platform.y = y

	img, _, err := ebitenutil.NewImageFromFile("assets/platform.png")
	if err != nil {
		log.Fatal(err)
	}
	platform.Img = img

	// We force platform width to be 70
	platform.scaleX = 80 / float32(platform.Img.Bounds().Dx())
	platform.scaleY = platform.scaleX
	platform.width = float32(platform.Img.Bounds().Dx()) * platform.scaleX
	platform.height = float32(platform.Img.Bounds().Dy()) * platform.scaleY

	return platform
}
