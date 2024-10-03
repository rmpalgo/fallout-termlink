package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rmpalgo/fallout-termlink/pkg/game"
)

// Model the state of the application. Working with
// the bubbletea example in their README.
// Bubbe Tea programs comprises a model that describes the
// application state and three methods on that model.
//
// Init()   initial commands to run.
// Update() handles incoming commands and updates the model state accordingly.
// View()   rendering method for the TUI.
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
