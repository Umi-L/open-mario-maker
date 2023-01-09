package geometry

type Transform struct {
	Position
	Rotation float64
}

func NewEmptyTransform() Transform {
	return Transform{Position{X: 0, Y: 0}, 0.0}
}
