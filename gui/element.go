package gui

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/umi-l/open-mario-maker/gui_update_params"
	"github.com/umi-l/open-mario-maker/utils"
)

type Element struct { //elements are just Containers with drawables
	Container

	Image *ebiten.Image

	initialized bool

	name string

	OnClick func()
}

func (e *Element) Init() {
	e.Transform = MakeTransformWithImage(e.Image, OriginTopLeft)
	e.CalculateRect()

	e.initialized = true
}

func (e Element) checkInitialized() {
	if !e.initialized {
		log.Fatal("Element is not initialized; this may cause unexpected behaviour and as such is an error. Call element.Init() to fix this.")
	}
}

func (e Element) Draw(screen *ebiten.Image) {
	e.checkInitialized()

	if !e.Visible {
		return
	}
	utils.DrawImageAtRect(screen, e.Image, e.Rect, &ebiten.DrawImageOptions{})
}

func (e Element) DrawTree(screen *ebiten.Image) {

	if !e.Visible {
		return
	}

	for _, child := range e.children {
		Draw(&e, screen)
		child.DrawTree(screen)
	}
}

func (e *Element) Update(params gui_update_params.UpdateParams) {
	Defaults.UpdateChildren(e, params)
}

func (e *Element) CalculateRect() {
	e.Rect = Defaults.CalculateRect(e)
}

func (e Element) GetContainer() Container {
	return e.Container
}

func (e *Element) SetParent(parent *Container) {
	e.Parent = parent
}

func MakeElement(image *ebiten.Image) Element {
	elm := Element{Image: image}
	elm.Init()

	elm.Visible = true

	return elm
}

type ElementInterface interface {
	Draw(screen *ebiten.Image)
	DrawTree(screen *ebiten.Image)
	Update(gui_update_params.UpdateParams)
	CalculateRect()
	SetParent(parent *Container)
	GetContainer() Container
}
