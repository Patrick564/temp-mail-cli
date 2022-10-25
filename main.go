package main

import (
	"fmt"
	"os"

	"github.com/Patrick564/temp-mail-cli/pkg/cmd/root"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(root.InitialModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("Error in program exec.: %v", err)
		os.Exit(1)
	}
}
