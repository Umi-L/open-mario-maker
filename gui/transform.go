package gui

type Transform struct {
	//absolute pixel values
	X float32
	Y float32
	W float32
	H float32

	//Relative % values
	XPercent float32
	YPercent float32
	WPercent float32
	HPercent float32
}