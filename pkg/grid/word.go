package grid

var fourLetterWordList = []string{
	"NICE", "REST",
	"NOTE", "CORN",
	"MOVE", "LOVE",
	"HACK", "TERM",
	"PLOW", "GROW",
	"WEST", "EAST",
	"LEFT", "EXIT",
}

type Word struct {
	Text      string
	Positions []Position
	Found     bool
}
