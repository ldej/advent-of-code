package tools

type Point struct {
	X, Y int
}

func (p *Point) Add(v Point) {
	p.X += v.X
	p.Y += v.Y
}

func (p *Point) ManhattanDistance(a Point) int {
	x := 0
	if p.X > a.X {
		x = p.X - a.X
	} else {
		x = a.X - p.X
	}

	y := 0
	if p.Y > a.Y {
		y = p.Y - a.Y
	} else {
		y = a.Y - p.Y
	}
	return x + y
}
