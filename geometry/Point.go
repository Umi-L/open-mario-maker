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

func (p Point) Mul(p2 Point) Point {
	p.X *= p2.X
	p.Y *= p2.Y

	return p
}

func (p Point) Div(p2 Point) Point {
	p.X /= p2.X
	p.Y /= p2.Y

	return p
}

// divide by float64
func (p Point) DivF(f float64) Point {
	p.X /= f
	p.Y /= f

	return p
}

// multiply by float64
func (p Point) MulF(f float64) Point {
	p.X *= f
	p.Y *= f

	return p
}
