package main

import (
	"log"
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

	rows := FetchDockerContainers()
	tableModel := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)

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

	m := model{table: tableModel}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		log.Println("Error running program:", err)
		os.Exit(1)
	}
}
