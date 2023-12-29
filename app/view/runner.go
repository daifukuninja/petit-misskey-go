package view

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Model interface {
	tea.Model
	// SetProgram(p *tea.Program)
	ReceiverChannel() chan Model // TODO: Modelじゃなくてtea.Msg
}

func Run(model Model) {
	p := tea.NewProgram(model)
	ch := model.ReceiverChannel()

	go func() {
		for {
			<-ch
			p.Send(model)
		}
	}()

	if _, err := p.Run(); err != nil {
		fmt.Printf("tea runner error. %v", err) // TODO: ちゃんとエラー処理
		os.Exit(1)
	}
}
