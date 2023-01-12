package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Element struct { //elements are just Containers with drawables
	Container

	image ebiten.Image
}

func (e Element) draw(screen ebiten.Image){
	
}

type ElementInterface interface {
	draw()
}