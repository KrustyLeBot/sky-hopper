package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Platform struct {
	x      float32
	y      float32
	width  float32
	height float32
}

func (platform *Platform) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, platform.x-(platform.width/2), platform.y-(platform.height/2), platform.width, platform.height, color.White, false)
}

func NewPlatform() *Platform {
	platform := &Platform{}

	platform.x = float32(WIDTH) / 2
	platform.y = float32(HEIGHT) - 30
	platform.width = 70
	platform.height = 10

	return platform
}
