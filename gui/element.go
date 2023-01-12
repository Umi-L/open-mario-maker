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
	utils.DrawImageAtRect(&screen, &e.Image, e.Rect, &ebiten.DrawImageOptions{})

	log.Print("guidrawcall")
}
func (e Element) DrawTree(screen ebiten.Image){
	for _, child := range e.children{
		Draw(e, screen)
		child.DrawTree(screen)
	}
}
func (e Element) CalculateRect(){
	Defaults.CalculateRect(e)
}
func (e Element) GetContainer() Container{
	return e.Container
}

func (e Element) SetParent(parent Container) {
	e.Parent = &parent
}

type ElementInterface interface {
	Draw(screen ebiten.Image)
	DrawTree(screen ebiten.Image)
	CalculateRect()
	SetParent(parent Container)
	GetContainer() Container
}