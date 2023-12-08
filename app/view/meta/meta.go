package meta

import (
	"context"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/daifukuninja/petit-misskey-go/service/meta"
	"github.com/daifukuninja/petit-misskey-go/util"
)

type Model struct {
	viewport viewport.Model
	service  *meta.Service
	ctx      context.Context
	quitting bool
}

func NewModel(ctx context.Context, service *meta.Service) *Model {
	const width = 120

	vp := viewport.New(width, 7)
	vp.Style = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("62")).PaddingRight(2)

	j, err := service.Do(ctx)
	if err != nil {
		return nil
	}

	vp.SetContent(util.PrittyJson(j))

	return &Model{
		viewport: vp,
		service:  service,
		ctx:      ctx,
		quitting: false,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}
	default:
		return m, nil
	}
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (m Model) View() string {
	view := m.viewport.View()
	if m.quitting {
		view += "\n" // NOTE: 終了時に最後の行がつぶれないようにする
	}
	return view
}
