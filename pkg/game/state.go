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
	Likeness    int
	LikenessMsg []string
	Attempts    int
	Current     gameState
}

func NormalMode() *State {
	return &State{
		Likeness:    0,
		LikenessMsg: make([]string, 0),
		Attempts:    allowedAttempts,
		Current:     Main,
	}
}
