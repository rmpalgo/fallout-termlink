package grid

type Position struct {
	Row int
	Col int
}

func NewPosition() *Position {
	return &Position{Row: 0, Col: 0}
}

func (p *Position) PosX() int {
	return p.Row
}

func (p *Position) PosY() int {
	return p.Col
}

func (p *Position) MoveUp() {
	p.Row--
}

func (p *Position) MoveDown() {
	p.Row++
}

func (p *Position) MoveRight() {
	p.Col--
}

func (p *Position) MoveLeft() {
	p.Col++
}
