package model

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	regularModeWordCount = 4
	hardModeWordCount    = 4
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
}

func InitialModel() *Model {
	return &Model{
		choices:         []string{"BROW", "GROW", "NOTE"},
		selected:        "",
		correctPassword: "GROW",
		likeness:        0,
		likenessMsg:     "",
		wordCount:       regularModeWordCount,
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
			likeness := calculateLikeness(m.choices[m.cursor], m.correctPassword)
			m.likeness = likeness
			m.unlocked = m.wordCount == likeness
			m.likenessMsg = fmt.Sprintf("Likeness: %d", m.likeness)
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Not returning a command.
	return m, nil
}

// View function that renders the UI based on the data model. Simply, this
// is rendering strings.
func (m *Model) View() string {
	if m.unlocked {
		return "For Overseer Eyes Only!\nClearance Granted\n"
	}
	// Writing the header
	s := "Welcome to RobCo Industries Termlink.\nPassword Required.\n\n"

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
