package meta

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/daifukuninja/petit-misskey-go/domain/view"
	"github.com/daifukuninja/petit-misskey-go/infrastructure/bubbles"
	"github.com/daifukuninja/petit-misskey-go/service/meta"
	"github.com/daifukuninja/petit-misskey-go/util"
	runner "github.com/daifukuninja/petit-misskey-go/view"
)

type Model struct {
	view       view.SimpleView
	service    *meta.Service
	ctx        context.Context
	quitting   bool
	teaProgram *tea.Program
}

func NewModel(service *meta.Service, viewFactory bubbles.SimpleViewFactory) *Model {
	ctx := context.Background()

	j, err := service.Do(ctx)
	if err != nil {
		return nil
	}

	view := viewFactory.View()
	view.SetContent(util.PrittyJson(j))

	return &Model{
		view:     view,
		service:  service,
		ctx:      ctx,
		quitting: false,
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
func (m *Model) View() string {
	view := m.view.View()
	if m.quitting {
		view += "\n" // NOTE: 終了時に最後の行がつぶれないようにする
	}
	return view
}

func (m *Model) SetProgram(p *tea.Program) {
	m.teaProgram = p
}

func (m *Model) ReceiverChannel() chan runner.Model {
	return nil
}
