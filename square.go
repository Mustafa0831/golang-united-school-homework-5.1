package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s *Square) End() Point {
	e := Point{
		s.start.x + int(s.a),
		s.start.y + int(s.a),
	}
	return e
}

func (p *Square) Area() uint {
	// implement me
	return uint(p.a) * uint(p.a)
}

func (p *Square) Perimeter() uint {
	// implement
	return uint(2*p.a) + uint(2*p.a)
}
