package model

import (
	"fmt"
	"strings"
)

// View function that renders the UI based on the data model. Simply, this
// is rendering strings.
func (m *Model) View() string {
	if m.unlocked {
		return "For Overseer Eyes Only!\nClearance Granted.\n"
	}
	if m.attempts == 0 {
		return "TERMINAL LOCKED\nPlease Contact an Administrator!"
	}
	// Writing the header
	s := "Welcome to RobCo Industries Termlink.\nPassword Required.\n"
	s += fmt.Sprintf("Attempts left: %s\n\n", strings.Repeat("▓ ", m.attempts))

	for i, choice := range m.choices {
		// The cursor will point at the choice
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // the cursor!
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	s += "\n"
	if m.likenessMsg != "" {
		s += fmt.Sprintf("> %s", m.likenessMsg)
	}

	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
