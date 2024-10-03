package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rmpalgo/fallout-termlink/pkg/game"
)

// Model the state of the game. Working with
// the bubbletea example in their README.
type Model struct {
	GameState *game.State
	cursor    int
	unlocked  bool
}

func InitialModel() *Model {
	return &Model{
		GameState: game.NormalMode(),
	}
}

// Init initial commands for the app to run.
func (m *Model) Init() tea.Cmd {
	// For now, just return nil, meaning no I/O.
	return nil
}
