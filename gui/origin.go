package gui

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/umi-l/open-mario-maker/types"
)

type OriginIndex int32

type Origin struct {
	X float32
	Y float32

	OriginIndex OriginIndex
}

const (
	OriginTopLeft      OriginIndex = 0
	OriginTopRight                 = 1
	OriginTopCenter                = 2
	OriginBottomLeft               = 3
	OriginBottomRight              = 4
	OriginBottomCenter             = 5
	OriginCenterLeft               = 6
	OriginCenterRight              = 7
	OriginCenter                   = 8
)

func CalculateOriginFromImage(image *ebiten.Image, origin OriginIndex) Origin {
	switch origin {
	case OriginTopLeft:
		return Origin{0, 0, origin}
	case OriginTopRight:
		return Origin{float32(image.Bounds().Max.X), 0, origin}
	case OriginTopCenter:
		return Origin{float32(image.Bounds().Max.X) / 2, 0, origin}
	case OriginBottomLeft:
		return Origin{0, float32(image.Bounds().Max.Y), origin}
	case OriginBottomRight:
		return Origin{float32(image.Bounds().Max.X), float32(image.Bounds().Max.Y), origin}
	case OriginBottomCenter:
		return Origin{float32(image.Bounds().Max.X) / 2, float32(image.Bounds().Max.Y), origin}
	case OriginCenterLeft:
		return Origin{0, float32(image.Bounds().Max.Y) / 2, origin}
	case OriginCenterRight:
		return Origin{float32(image.Bounds().Max.X), float32(image.Bounds().Max.Y) / 2, origin}
	case OriginCenter:
		return Origin{float32(image.Bounds().Max.X) / 2, float32(image.Bounds().Max.Y) / 2, origin}
	default:
		fmt.Println("Invalid OriginIndex")
		return Origin{0, 0, origin}
	}
}

func CalculateOriginFromRect(rect types.Rect, origin OriginIndex) Origin {
	switch origin {
	case OriginTopLeft:
		return Origin{0, 0, origin}
	case OriginTopRight:
		return Origin{rect.W, 0, origin}
	case OriginTopCenter:
		return Origin{rect.W / 2, 0, origin}
	case OriginBottomLeft:
		return Origin{0, rect.H, origin}
	case OriginBottomRight:
		return Origin{rect.W, rect.H, origin}
	case OriginBottomCenter:
		return Origin{rect.W / 2, rect.H, origin}
	case OriginCenterLeft:
		return Origin{0, rect.H / 2, origin}
	case OriginCenterRight:
		return Origin{rect.W, rect.H / 2, origin}
	case OriginCenter:
		return Origin{rect.W / 2, rect.H / 2, origin}
	default:
		fmt.Println("Invalid OriginIndex")
		return Origin{0, 0, origin}
	}
}
