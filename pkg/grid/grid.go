package grid

import "math/rand"

const (
	gridRows int = 8
	gridCols int = 24
)

type Grid struct {
	PositionToWord map[Position]*Word
	Words          []Word
	Data           [][]rune
	// Number of columns
	Width int
	// Number of rows
	Height  int
	Regions []Region
	// Use memoization to track region with placed words to avoid
	// multiple words placed in the same region.
	UsedRegions map[int]map[Region]bool
}

type Region struct {
	StartCol int
	EndCol   int
}

func NormalMode() *Grid {
	g := &Grid{
		PositionToWord: make(map[Position]*Word),
		// Width is the grid cols (e.g., 24)
		Width: gridCols,
		// Height is the grid rows (e.g., 8)
		Height: gridRows,
		Regions: []Region{
			// Left address area (columns 0-11)
			{StartCol: 0, EndCol: 11},
			// Right address area (columns 12-23)
			{StartCol: 12, EndCol: 23},
		},
		UsedRegions: make(map[int]map[Region]bool),
	}
	g.initializeGrid()
	g.placeWords()
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

func (g *Grid) placeWords() {
	rand.Shuffle(len(fourLetterWordList), func(i, j int) {
		fourLetterWordList[i], fourLetterWordList[j] = fourLetterWordList[j], fourLetterWordList[i]
	})

	g.Words = make([]Word, 0, len(fourLetterWordList))

	for _, wordText := range fourLetterWordList {
		positions := g.findSpaceForWord(len(wordText))
		if positions != nil {
			word := Word{Text: wordText, Positions: positions}
			g.Words = append(g.Words, word)

			for idx, pos := range positions {
				g.Data[pos.PosX()][pos.PosY()] = rune(wordText[idx])
				g.PositionToWord[pos] = &g.Words[len(g.Words)-1]
			}
		}
	}
}

func (g *Grid) findSpaceForWord(length int) []Position {
	maxAttempts := 100

	for attempt := 0; attempt < maxAttempts; attempt++ {
		// Randomly select a region
		regionIndex := rand.Intn(len(g.Regions))
		region := g.Regions[regionIndex]
		// Randomly select a row
		row := rand.Intn(g.Height)

		if _, exists := g.UsedRegions[row]; !exists {
			g.UsedRegions[row] = make(map[Region]bool)
		}

		// If the region is already used in this row, skip to next attempt
		if g.UsedRegions[row][region] {
			continue
		}

		// Ensure the word fits within the selected region
		regionWidth := region.EndCol - region.StartCol
		if regionWidth < length {
			// Word is too long for this region, skip to next attempt
			continue
		}

		// Calculate the maximum starting column within the region
		maxStartCol := region.EndCol - length
		if maxStartCol < region.StartCol {
			// Not enough space to place the word, skip
			continue
		}

		// Pick a random starting column within the allowed range
		col := rand.Intn(maxStartCol-region.StartCol+1) + region.StartCol

		// Collect positions and check for conflicts
		positions := make([]Position, length)
		conflict := false
		for i := 0; i < length; i++ {
			pos := Position{Row: row, Col: col + i}
			if _, exists := g.PositionToWord[pos]; exists {
				conflict = true
				break
			}
			positions[i] = pos
		}

		if !conflict {
			// Track region has word placed
			g.UsedRegions[row][region] = true
			return positions
		}
	}

	// Failed to find space after maxAttempts
	return nil
}
