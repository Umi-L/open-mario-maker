package gui

import "github.com/hajimehoshi/ebiten/v2"

type Transform struct {
	//absolute pixel values
	X float32
	Y float32
	W float32
	H float32

	//Relative % values
	XPercent float32
	YPercent float32
	WPercent float32
	HPercent float32

	Origin Origin
}

func MakeTransformWithImage(image *ebiten.Image, origin OriginIndex) Transform {
	return Transform{
		W: float32(image.Bounds().Max.X),
		H: float32(image.Bounds().Max.Y),

		Origin: Origin{OriginIndex: origin},
	}
}
