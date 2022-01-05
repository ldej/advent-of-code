package tools

type Point struct {
	X, Y int
}

func (p *Point) Add(v Point) {
	p.X += v.X
	p.Y += v.Y
}
