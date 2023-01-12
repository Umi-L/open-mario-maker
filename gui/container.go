package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	. "github.com/umi-l/open-mario-maker/types"
)

type Container struct {
	Parent *Container
	Rect Rect

	Transform Transform

	children []ElementInterface
}

func (c *Container) SetTransform(t Transform){
	c.Transform = t
	c.CalculateRect()
}

func (c *Container) AddChild(child ElementInterface){
	c.children = append(c.children, child)
	child.SetParent(*c)
	child.CalculateRect()
}

func (c Container) SetParent(parent Container){
	c.Parent = &parent
}

func (c Container) Draw(screen ebiten.Image){}

func (c Container) DrawTree(screen ebiten.Image){
	for _, child := range c.children{
		Draw(c, screen)
		child.DrawTree(screen)
	}
}

func (c Container) GetContainer() Container{
	return c
}


func (c Container) CalculateRect(){
	Defaults.CalculateRect(c)
}

func NewRelativeContainer(parent *Container) Container{
	newContainer := Container{
		Parent: parent,
	}

	parent.children = append(parent.children, newContainer)

	return newContainer
}

func NewRootContainer(screenW int, screenH int) Container{

	return Container{
		Rect: Rect{
			X: 0,
			Y: 0,
			W: float32(screenW),
			H: float32(screenH),
		},
	}
}