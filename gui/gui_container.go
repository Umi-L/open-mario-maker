package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	. "github.com/umi-l/open-mario-maker/types"
)

type Container struct {
	Parent *Container
	rect Rect

	Transform Transform

	children []*Container
}

func (c *Container) SetTransform(t Transform){
	c.Transform = t
	c.calculateRect()
}

func (c *Container) AddChild(child *Container){
	c.children = append(c.children, child)
	child.Parent = c
	child.calculateRect()
}

func (c *Container) calculateRect(){


	//X
	if c.Transform.XPercent == 0{ //check for default value, also cant be 0% as that devides by 0
		c.rect.X = c.Transform.X
	}else{
		c.rect.X = float32(c.Parent.rect.W) / c.Transform.XPercent
	}

	//Y
	if c.Transform.YPercent == 0{ //check for default value, also cant be 0% as that devides by 0
		c.rect.Y = c.Transform.Y
	}else{
		c.rect.Y = float32(c.Parent.rect.H) / c.Transform.YPercent
	}

	//W
	if c.Transform.WPercent == 0{ //check for default value, also cant be 0% as that devides by 0
		c.rect.W = c.Transform.W
	}else{
		c.rect.W = float32(c.Parent.rect.W) / c.Transform.WPercent
	}

	//H
	if c.Transform.HPercent == 0{ //check for default value, also cant be 0% as that devides by 0
		c.rect.H = c.Transform.H
	}else{
		c.rect.H = float32(c.Parent.rect.H) / c.Transform.HPercent
	}

	for _, child := range c.children{
		child.calculateRect()
	}
}

func NewRelativeContainer(parent *Container) Container{
	newContainer := Container{
		Parent: parent,
	}

	parent.children = append(parent.children, &newContainer)

	return newContainer
}

func NewRootContainer(screen ebiten.Image) Container{

	screenW, screenH := screen.Size()

	return Container{
		rect: Rect{
			X: 0,
			Y: 0,
			W: float32(screenW),
			H: float32(screenH),
		},
	}
}