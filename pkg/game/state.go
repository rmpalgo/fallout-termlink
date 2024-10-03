package game

const (
	regularModeWordCount = 4
	hardModeWordCount    = 5

	allowedAttempts = 4
)

type gameState int

const (
	Main = iota
	Unlocked
	Locked
)

type State struct {
	Choices         []string
	Likeness        int
	LikenessMsg     string
	CorrectPassword string
	Attempts        int
	Selected        string
	WordCount       int
	Current         gameState
}

func NormalMode() *State {
	return &State{
		Choices:         []string{"BROW", "GROW", "NOTE"},
		Selected:        "",
		CorrectPassword: "GROW",
		Likeness:        0,
		LikenessMsg:     "",
		Attempts:        allowedAttempts,
		WordCount:       regularModeWordCount,
	}
}
