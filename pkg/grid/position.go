package grid

type Position struct {
	Row int
	Col int
}

func NewPosition() Position {
	return Position{Row: 0, Col: 0}
}

func (p *Position) PosX() int {
	return p.Row
}

func (p *Position) PosY() int {
	return p.Col
}
