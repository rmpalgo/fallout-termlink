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
			case "up", "w":
				m.moveUp()
			case "down", "s":
				m.moveDown()
			case "left", "a":
				m.moveLeft()
			case "right", "d":
				m.moveRight()

			// Enter and space press is to enter the selection where the
			// cursor is currently positioned.
			case "enter":
				pos := m.CursorPosition

				// Check if the position is part of a word
				if word, exists := m.Grid.PositionToWord[pos]; exists && !word.Found {
					if word.Text == m.Grid.CorrectPassword {
						m.unlocked = true
						m.GameState.Current = game.Unlocked
						m.GameState.LikenessMsg = append(m.GameState.LikenessMsg, fmt.Sprintf("Found Password: %s", word.Text))
						return m, nil
					}

					total := calculateLikeness(word.Text, m.Grid.CorrectPassword)
					m.GameState.Likeness = total
					m.GameState.Attempts--
					m.GameState.LikenessMsg = append(m.GameState.LikenessMsg, fmt.Sprintf("Likeness: %d", total))
				}

				if m.GameState.Attempts == 0 {
					m.GameState.Current = game.Locked
				}

				return m, nil
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
		m.CursorPosition.Row--
	}
}

func (m *Model) moveDown() {
	if m.CursorPosition.Row < len(m.Grid.Data)-1 {
		m.CursorPosition.Row++
	}
}

func (m *Model) moveLeft() {
	if m.CursorPosition.Col > 0 {
		if word, exists := m.Grid.PositionToWord[m.CursorPosition]; exists {
			// When cursor is on a word, then skip to next single character.
			endPos := word.Positions[len(word.Positions)-1].Col
			m.CursorPosition.Col = endPos - len(word.Text)
			return
		}
		m.CursorPosition.Col--
	}
}

func (m *Model) moveRight() {
	if m.CursorPosition.Col < len(m.Grid.Data[0])-1 {
		if word, exists := m.Grid.PositionToWord[m.CursorPosition]; exists {
			// When cursor is on a word, then skip over to
			// next single character.
			startPos := word.Positions[0].Col
			m.CursorPosition.Col = startPos + len(word.Text)
			return
		}
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
