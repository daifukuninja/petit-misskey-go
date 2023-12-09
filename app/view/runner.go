package view

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func Run(model tea.Model) {
	if _, err := tea.NewProgram(model).Run(); err != nil {
		fmt.Printf("tea runner error. %v", err) // TODO: ちゃんとエラー処理
		os.Exit(1)
	}
}
