package termlink

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rmpalgo/fallout-termlink/pkg/model"
)

func Run() error {
	m := model.InitialModel()

	program := tea.NewProgram(m)

	if _, err := program.Run(); err != nil {
		return fmt.Errorf("failed to run tea application: %w", err)
	}

	return nil
}
