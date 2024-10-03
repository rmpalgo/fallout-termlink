package model

import (
	"fmt"
	"strings"

	"github.com/rmpalgo/fallout-termlink/pkg/game"
)

// View function that renders the UI based on the data model. Simply, this
// is rendering strings.
func (m *Model) View() string {
	switch m.GameState.Current {
	case game.Main:
		// Writing the header
		s := "Welcome to RobCo Industries Termlink.\nPassword Required.\n"
		s += fmt.Sprintf("Attempts left: %s\n\n", strings.Repeat("▓ ", m.GameState.Attempts))

		for i, choice := range m.GameState.Choices {
			// The cursor will point at the choice
			cursor := " " // no cursor
			if m.cursor == i {
				cursor = ">" // the cursor!
			}

			s += fmt.Sprintf("%s %s\n", cursor, choice)
		}
		s += "\n"
		if m.GameState.LikenessMsg != "" {
			s += fmt.Sprintf("> %s", m.GameState.LikenessMsg)
		}

		s += "\nPress q to quit.\n"

		// Send the UI for rendering
		return s
	case game.Unlocked:
		return "For Overseer Eyes Only!\nClearance Granted.\n"
	case game.Locked:
		return "TERMINAL LOCKED\nPlease Contact an Administrator!"
	}

	return ""
}
