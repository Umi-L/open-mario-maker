package tiled

import "github.com/umi-l/open-mario-maker/loader"

type Map struct {
	Sheet loader.Sheet
	TileLayers  []TileLayer
}

type TileLayer struct {
	TileData []int
	
	width int
	height int

	id int
	name string
	visible bool

	x int
	y int
}