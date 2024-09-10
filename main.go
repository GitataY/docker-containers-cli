package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	columns := []table.Column{
		{Title: "Container ID", Width: 20},
		{Title: "Image", Width: 25},
		{Title: "COMMAND", Width: 25},
		{Title: "CREATED", Width: 20},
		{Title: "STATUS", Width: 15},
		{Title: "PORTS", Width: 30},
		{Title: "NAMES", Width: 20},
	}

	// Fetch container data from models.go
	rows := FetchDockerContainers()

	// Create table
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	// Initialize the TUI model
	m := model{t}

	// Run the TUI program
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
