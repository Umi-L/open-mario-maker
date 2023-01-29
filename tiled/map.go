package tiled

import (
	"github.com/umi-l/waloader"
)

type Map struct {
	Sheet      waloader.Sheet
	TileLayers []TileLayer
}

type TileLayer struct {
	TileData []int

	width  int
	height int

	id      int
	name    string
	visible bool

	x int
	y int
}
