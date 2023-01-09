package animation

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/umi-l/open-mario-maker/loader"
	"image"
)

type Animation struct {
	parent       *loader.Sheet
	index        int
	frames       int
	sprites      []image.Rectangle
	speed        float32
	timer        float32
	currentFrame int
}

func (anim *Animation) Draw(screen *ebiten.Image, x float64, y float64, rotation float64) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(x, y)
	op.GeoM.Rotate(rotation)

	screen.DrawImage(anim.parent.Texture.SubImage(anim.sprites[anim.currentFrame]).(*ebiten.Image), op)
}

func (anim *Animation) updateTimer(dt float32) {
	anim.timer += dt

	if anim.timer > anim.speed {
		anim.currentFrame += 1
		anim.timer = 0.0
	}

	if anim.currentFrame == anim.frames {
		anim.currentFrame = 0
	}
}
