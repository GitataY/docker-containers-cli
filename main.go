package main

import (
	"log"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	ContainerIDWidth = 20
	ImageWidth       = 25
	CommandWidth     = 25
	CreatedWidth     = 20
	StatusWidth      = 15
	PortsWidth       = 30
	NamesWidth       = 20
	TableHeight      = 7
)

func main() {
	columns := []table.Column{
		{Title: "Container ID", Width: ContainerIDWidth},
		{Title: "Image", Width: ImageWidth},
		{Title: "COMMAND", Width: CommandWidth},
		{Title: "CREATED", Width: CreatedWidth},
		{Title: "STATUS", Width: StatusWidth},
		{Title: "PORTS", Width: PortsWidth},
		{Title: "NAMES", Width: NamesWidth},
	}

	// Fetch container data from models.go
	rows := FetchDockerContainers()

	// Create table model
	tableModel := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(TableHeight),
	)

	// Set styles
	styles := table.DefaultStyles()
	styles.Header = styles.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	styles.Selected = styles.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	tableModel.SetStyles(styles)

	// Initialize the TUI model
	m := model{table: tableModel}

	// Run the TUI program
	if _, err := tea.NewProgram(m).Run(); err != nil {
		log.Println("Error running program:", err)
		os.Exit(1)
	}
}
