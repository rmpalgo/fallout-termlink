package model

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

// Model the state of the game. Working with
// the bubbletea example in their README.
type Model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func InitialModel() *Model {
	return &Model{
		choices:  []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},
		selected: make(map[int]struct{}),
	}
}

// Init initial commands for the app to run.
func (m *Model) Init() tea.Cmd {
	// For now, just return nil, meaning no I/O.
	return nil
}

// Update handles incoming events such as key presses and updates
// the model accordingly.
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// Enter and space press is to enter the selection where the
		// cursor is currently positioned.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Not returning a command.
	return m, nil
}

// View function that renders the UI based on the data model. Simply, this
// is rendering strings.
func (m *Model) View() string {
	// Writing the header
	s := "What should we buy at the market?\n\n"

	for i, choice := range m.choices {

		// The cursor will point at the choice
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // the cursor!
		}

		// Has the choice been selected?
		// this is when the user presses the key `enter` or `space`
		checked := " " // no then leave blank
		if _, ok := m.selected[i]; ok {
			checked = "x" // yes, it is selected!
		}

		// Render the row for the to-do list example.
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
