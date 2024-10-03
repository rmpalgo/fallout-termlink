package model

import (
	tea "github.com/charmbracelet/bubbletea"
)

const (
	regularModeWordCount = 4
	hardModeWordCount    = 5

	allowedAttempts = 4
)

// Model the state of the game. Working with
// the bubbletea example in their README.
type Model struct {
	choices         []string
	cursor          int
	selected        string
	likeness        int
	likenessMsg     string
	correctPassword string
	unlocked        bool
	wordCount       int
	attempts        int
}

func InitialModel() *Model {
	return &Model{
		choices:         []string{"BROW", "GROW", "NOTE"},
		selected:        "",
		correctPassword: "GROW",
		likeness:        0,
		likenessMsg:     "",
		wordCount:       regularModeWordCount,
		attempts:        allowedAttempts,
	}
}

// Init initial commands for the app to run.
func (m *Model) Init() tea.Cmd {
	// For now, just return nil, meaning no I/O.
	return nil
}
