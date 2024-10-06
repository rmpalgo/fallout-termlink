package model

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/rmpalgo/fallout-termlink/pkg/game"
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
		attemptsMsg(m.GameState.Attempts)
		sb.WriteRune('\n')

		// Grid rows and cells
		for i, row := range m.Grid.Data {
			// 24 cells
			cells := len(row)
			// middle index
			mid := (cells / 2) - 1
			// add addressA first then when middle index add addressB
			sb.WriteString(defaultStyle.Render(fmt.Sprintf("%s%d ", startingAddrA, i)))
			for j, cell := range row {
				cellStr := string(cell)

				cursor := m.CursorPosition
				isCursorPos := (cursor.PosX() == i) && (cursor.PosY() == j)

				if isCursorPos {
					sb.WriteString(renderCursor(cellStr))
				} else {
					sb.WriteString(renderDefault(cellStr))
				}

				if j == mid {
					addr := fmt.Sprintf(" %s%d ", startingAddrB, i)
					sb.WriteString(renderDefault(addr))
				}
			}
			sb.WriteRune('\n')
		}

		return sb.String()
	case game.Unlocked:
		return "For Overseer Eyes Only!\nClearance Granted.\n"
	case game.Locked:
		return "TERMINAL LOCKED\nPlease Contact an Administrator!"
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
