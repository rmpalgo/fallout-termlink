package grid

import "math/rand"

const (
	gridRows int = 8
	gridCols int = 24
)

type Grid struct {
	Data [][]rune
}

func NormalMode() *Grid {
	g := &Grid{}
	g.initializeGrid()
	return g
}

func (g *Grid) initializeGrid() {
	rows := gridRows
	cols := gridCols
	g.Data = make([][]rune, rows)
	for i := 0; i < rows; i++ {
		g.Data[i] = make([]rune, cols)
		for j := 0; j < cols; j++ {
			g.Data[i][j] = randomSpecialChar()
		}
	}
}

func randomSpecialChar() rune {
	chars := []rune{
		'{',
		'}',
		'[',
		']',
		'<',
		'>',
		'$',
		'#',
		'&',
		'+',
		'@',
		'!',
		'*',
		'^',
		'(',
		')',
		'/',
		'|',
		'\\',
		'%',
		'+',
		'=',
	}
	r := rune(rand.Intn(len(chars)))
	return chars[r]
}
