package model

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/rmpalgo/fallout-termlink/pkg/game"
	"github.com/rmpalgo/fallout-termlink/pkg/grid"
)

var (
	defaultStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FF00"))

	cursorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#00FF00"}).
			Background(lipgloss.AdaptiveColor{Light: "EE6FF8", Dark: "#005F00"})
)

const (
	startingAddrA = "0xFB2A0"
	startingAddrB = "0xFB2B0"
)

// View function that renders the UI based on the data model. Simply, this
// is rendering strings.
func (m *Model) View() string {
	switch m.GameState.Current {
	case game.Main:
		var sb strings.Builder
		// Welcome message and attempts
		sb.WriteString(welcomeMsg())
		sb.WriteRune('\n')
		sb.WriteString(passwordReqMsg())
		sb.WriteRune('\n')
		sb.WriteString(renderDefault(fmt.Sprintf("Correct Password: %s", m.Grid.CorrectPassword)))
		sb.WriteRune('\n')
		sb.WriteString(attemptsMsg(m.GameState.Attempts))
		sb.WriteRune('\n')
		sb.WriteRune('\n')

		// Determine if the cursor is over a word or a special sequence
		var wordToHighlight *grid.Word
		if word, exists := m.Grid.PositionToWord[m.CursorPosition]; exists && !word.Found {
			wordToHighlight = word
		}

		// Determine positions to highlight
		highlightPositions := make(map[grid.Position]bool)
		if wordToHighlight != nil {
			for _, pos := range wordToHighlight.Positions {
				highlightPositions[pos] = true
			}
		} else {
			highlightPositions[m.CursorPosition] = true
		}

		// Render the grid
		for i, row := range m.Grid.Data {
			cells := len(row)
			mid := (cells / 2) - 1 // Middle index for AddressB

			// Render AddressA with row number
			sb.WriteString(defaultStyle.Render(fmt.Sprintf("%s%02d ", startingAddrA, i)))

			for j, cell := range row {
				pos := grid.Position{Row: i, Col: j}
				cellStr := string(cell)

				// Apply styles based on the position
				if highlightPositions[pos] {
					if wordToHighlight != nil {
						// Apply wordStyle to entire word
						cellStr = renderCursor(cellStr)
					} else {
						// Apply cursorStyle to single character
						cellStr = renderCursor(cellStr)
					}
				} else {
					cellStr = renderDefault(cellStr)
				}

				sb.WriteString(cellStr)

				// Insert AddressB at the middle of the row
				if j == mid {
					addr := fmt.Sprintf(" %s%02d ", startingAddrB, i)
					sb.WriteString(renderDefault(addr))
				}
			}

			if i != len(m.Grid.Data)-1 {
				sb.WriteRune('\n')
			} else {
				sb.WriteString(renderDefault(fmt.Sprintf(" > %s", m.GameState.LikenessMsg)))
				sb.WriteRune('\n')
			}
		}

		return sb.String()
	case game.Unlocked:
		return renderDefault("For Overseer Eyes Only!\nClearance Granted.\n")
	case game.Locked:
		return renderDefault("TERMINAL LOCKED\nPlease Contact an Administrator!")
	}

	return ""
}

// default text
func renderDefault(s string) string {
	return defaultStyle.Render(s)
}

// cursor highlight
func renderCursor(s string) string {
	return cursorStyle.Render(s)
}

func welcomeMsg() string {
	title := "Welcome to RobCo Industries Termlink."
	return renderDefault(title)
}

func passwordReqMsg() string {
	msg := "Password Required"
	return renderDefault(msg)
}

func attemptsMsg(count int) string {
	attempts := strings.Repeat("▓ ", count)
	msg := fmt.Sprintf("Attempts left: %s", attempts)
	return renderDefault(msg)
}
