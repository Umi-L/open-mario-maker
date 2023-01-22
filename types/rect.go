package types

type Rect struct {
	X float32
	Y float32

	W float32
	H float32
}

// if rect contains point
func (r Rect) Contains(x, y float32) bool {
	return x >= r.X && x <= r.X+r.W && y >= r.Y && y <= r.Y+r.H
}
