package animation

import (
	"github.com/umi-l/open-mario-maker/loader"
	"image"
)

func Load(parent *loader.Sheet, index int, frames int, speed float32) Animation {
	var sprites []image.Rectangle

	for i := 0; i < frames; i++ {
		sprites = append(sprites, image.Rectangle{
			Min: image.Point{
				X: parent.CellWidth * i,
				Y: parent.CellHeight * index,
			},
			Max: image.Point{
				X: (parent.CellWidth * i) + parent.CellWidth,
				Y: (parent.CellHeight * index) + parent.CellHeight,
			},
		})
	}

	return Animation{
		parent:       parent,
		index:        index,
		frames:       frames,
		sprites:      sprites,
		speed:        speed,
		timer:        0.0,
		currentFrame: 0,
	}
}
