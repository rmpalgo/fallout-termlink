package model

import (
	"fmt"

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

			// up arrow and k to move cursor up.
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}

			// down arrow and j to move cursor down.
			case "down", "j":
				if m.cursor < len(m.GameState.Choices)-1 {
					m.cursor++
				}

			// Enter and space press is to enter the selection where the
			// cursor is currently positioned.
			case "enter", " ":
				likeness := calculateLikeness(m.GameState.Choices[m.cursor], m.GameState.CorrectPassword)
				m.GameState.Likeness = likeness

				// checks if game mode word count equals to likeness count.
				m.unlocked = m.GameState.WordCount == likeness
				m.GameState.LikenessMsg = fmt.Sprintf("Likeness: %d", m.GameState.Likeness)

				// If not unlocked then check if there are more than one attempt left.
				// Otherwise lock the game.
				if m.unlocked {
					m.GameState.Current = game.Unlocked
				} else if m.GameState.Attempts > 1 {
					m.GameState.Attempts--
				} else {
					m.GameState.Current = game.Locked
				}
			}
		}
	default:
		return m, tea.Quit
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Not returning a command.
	return m, nil
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
