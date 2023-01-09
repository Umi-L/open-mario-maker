package physics

type Velocity struct {
	X float64
	Y float64
}

func NewEmptyVelocity() Velocity {
	return Velocity{0.0, 0.0}
}
