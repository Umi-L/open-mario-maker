package gui

import (
	"github.com/umi-l/open-mario-maker/gui_update_params"
	"github.com/umi-l/open-mario-maker/types"
)

type defaultsT struct{} // unexported type

var Defaults defaultsT

func (_ defaultsT) CalculateRect(e ElementInterface) types.Rect {

	c := e.GetContainer()

	calcRect := types.Rect{}

	//X
	if c.Transform.XPercent == 0 { //check for default value, also cant be 0% as that devides by 0
		calcRect.X = c.Transform.X
	} else {
		calcRect.X = float32(c.Parent.Rect.W) * c.Transform.XPercent

	}

	//Y
	if c.Transform.YPercent == 0 { //check for default value, also cant be 0% as that devides by 0
		calcRect.Y = c.Transform.Y
	} else {
		calcRect.Y = float32(c.Parent.Rect.H) * c.Transform.YPercent
	}

	//W
	if c.Transform.WPercent == 0 { //check for default value, also cant be 0% as that devides by 0
		calcRect.W = c.Transform.W
	} else {
		calcRect.W = float32(c.Parent.Rect.W) * c.Transform.WPercent
	}

	//H
	if c.Transform.HPercent == 0 { //check for default value, also cant be 0% as that devides by 0
		calcRect.H = c.Transform.H
	} else {
		calcRect.H = float32(c.Parent.Rect.H) * c.Transform.HPercent
	}

	//Origin
	c.Transform.Origin = CalculateOriginFromRect(calcRect, c.Transform.Origin.OriginIndex)

	calcRect.X -= c.Transform.Origin.X
	calcRect.Y -= c.Transform.Origin.Y

	for _, child := range c.children {
		child.CalculateRect()
	}

	return calcRect
}

func (_ defaultsT) UpdateChildren(e ElementInterface, params gui_update_params.UpdateParams) {

	for _, child := range e.GetContainer().children {
		child.Update(params)
	}
}
