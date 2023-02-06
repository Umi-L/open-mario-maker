package geometry

type Point struct {
	X float64
	Y float64
}

func (p Point) Add(p2 Point) Point {
	p.X += p2.X
	p.Y += p2.Y
	
	return p
}

func (p Point) Sub(p2 Point) Point {
	p.X -= p2.X
	p.Y -= p2.Y

	return p
}
