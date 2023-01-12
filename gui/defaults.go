package gui

type defaultsT struct{} // unexported type

var Defaults defaultsT

func (_ defaultsT)CalculateRect(e ElementInterface){

	c := e.GetContainer()

	//X
	if c.Transform.XPercent == 0{ //check for default value, also cant be 0% as that devides by 0
		c.Rect.X = c.Transform.X
	}else{
		c.Rect.X = float32(c.Parent.Rect.W) / c.Transform.XPercent
	}

	//Y
	if c.Transform.YPercent == 0{ //check for default value, also cant be 0% as that devides by 0
		c.Rect.Y = c.Transform.Y
	}else{
		c.Rect.Y = float32(c.Parent.Rect.H) / c.Transform.YPercent
	}

	//W
	if c.Transform.WPercent == 0{ //check for default value, also cant be 0% as that devides by 0
		c.Rect.W = c.Transform.W
	}else{
		c.Rect.W = float32(c.Parent.Rect.W) / c.Transform.WPercent
	}

	//H
	if c.Transform.HPercent == 0{ //check for default value, also cant be 0% as that devides by 0
		c.Rect.H = c.Transform.H
	}else{
		c.Rect.H = float32(c.Parent.Rect.H) / c.Transform.HPercent
	}

	for _, child := range c.children{
		child.CalculateRect()
	}
}