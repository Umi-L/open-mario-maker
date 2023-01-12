package utils

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	. "github.com/umi-l/open-mario-maker/types"
)

func GetDt() func() time.Duration {
	lastUpdate := time.Now()
	dt := time.Since(lastUpdate)
	return func() time.Duration {
		dt = time.Since(lastUpdate)
		lastUpdate = time.Now()
		return dt
	}
}

func DrawImageAtRect(screen *ebiten.Image, image *ebiten.Image, rect Rect, op *ebiten.DrawImageOptions){
	sw, sh := screen.Size()
	iw, ih := image.Size()


	op.GeoM.Scale(float64(sw / iw), float64(sh / ih))

	screen.DrawImage(image, op)
}