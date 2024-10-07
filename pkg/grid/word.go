package grid

type Word struct {
	Text      string
	Positions []Position
	Found     bool
}

var fourLetterWordList = []string{
	"NICE", "REST",
	"NOTE", "CORN",
	"MOVE", "LOVE",
	"CLAW", "THAW",
	"HACK", "TERM",
	"PLOW", "GROW",
	"MANA", "PRAY",
	"GRAY", "TRAY",
	"WEST", "EAST",
	"LEFT", "EXIT",
	"BAND", "BEND",
	"BIND", "BOND",
	"BARD", "BUND",
	"BAWD", "CARD",
	"CORD", "CURD",
	"COLD", "CALD",
	"FATE", "FILT",
	"FORT", "FAST",
	"FEAT", "GATE",
	"GILT", "GOUT",
	"GRIT", "GUST",
	"BAKE", "CAKE",
	"FAKE", "LAKE",
	"MAKE", "RAKE",
	"COND",
}
