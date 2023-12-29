package stream

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	ctx context.Context
}

func NewModel() *Model {
	ctx := context.Background()
	return &Model{
		ctx: ctx,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	return "not implemented."
}
