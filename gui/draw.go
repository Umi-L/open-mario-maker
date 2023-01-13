package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func Draw(drawable ElementInterface, screen *ebiten.Image){
	drawable.Draw(screen)
}