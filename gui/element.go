package gui

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/umi-l/open-mario-maker/utils"
)

type Element struct { //elements are just Containers with drawables
	Container

	Image ebiten.Image
}

func (e Element) Draw(screen ebiten.Image){
	utils.DrawImageAtRect(&screen, &e.Image, e.rect, &ebiten.DrawImageOptions{})

	log.Print("guidrawcall")
}

type ElementInterface interface {
	Draw(screen ebiten.Image)
}