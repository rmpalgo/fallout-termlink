package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rmpalgo/fallout-termlink/pkg/game"
)

// Update handles incoming events such as key presses and updates
// the model accordingly.
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.GameState.Current {
	case game.Main:
		switch msg := msg.(type) {
		// This handles key presses.
		case tea.KeyMsg:

			// What was pressed?
			switch msg.String() {

			// Pressing q or ctrl+c will exit the program.
			case "ctrl+c", "q":
				return m, tea.Quit
			case "up":
				m.moveUp()
			case "down":
				m.moveDown()
			case "left":
				m.moveLeft()
			case "right":
				m.moveRight()

			// Enter and space press is to enter the selection where the
			// cursor is currently positioned.
			case "enter":

			}
		}
	default:
		return m, tea.Quit
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Not returning a command.
	return m, nil
}

func (m *Model) moveUp() {
	if m.CursorPosition.Row > 0 {
		m.CursorPosition.MoveUp()
	}
}

func (m *Model) moveDown() {
	if m.CursorPosition.Row < len(m.Grid.Data)-1 {
		m.CursorPosition.MoveDown()
	}
}

func (m *Model) moveLeft() {
	if m.CursorPosition.Col > 0 {
		m.CursorPosition.Col--
	}
}

func (m *Model) moveRight() {
	if m.CursorPosition.Col < len(m.Grid.Data[0])-1 {
		m.CursorPosition.Col++
	}
}

func calculateLikeness(selectedWord, password string) int {
	count := 0
	minLen := len(password)
	if len(selectedWord) != minLen {
		return count
	}

	for i := 0; i < minLen; i++ {
		if selectedWord[i] == password[i] {
			count++
		}
	}

	return count
}
