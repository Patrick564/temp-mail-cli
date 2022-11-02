package main

import (
	"fmt"
	"os"

	"github.com/Patrick564/temp-mail-cli/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error at load env file: %v", err)
		os.Exit(1)
	}

	p := tea.NewProgram(ui.InitialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Error in program exec.: %v", err)
		os.Exit(1)
	}
}
