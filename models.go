package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters" // Added import for filters
	"github.com/docker/docker/client"
)

const CommandTruncateLength = 20

func getBaseStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240"))
}

type model struct {
	table table.Model
}

// type terminalSize struct {
// 	width  int
// 	height int
// }

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.table.SetHeight(msg.Height - 2) // Leave some space for margins
		m.table.SetWidth(msg.Width)

	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	m.table, cmd = m.table.Update(msg)

	return m, cmd
}

func (m model) View() string {
	return getBaseStyle().Render(m.table.View()) + "\n"
}

func FetchDockerContainers() []table.Row {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.42"))
	if err != nil {
		panic(err)
	}

	// Initialize Filters using filters.NewArgs()
	containers, err := cli.ContainerList(context.Background(), container.ListOptions{
		All:     true,              // List all containers (both running and stopped)
		Limit:   0,                 // No limit to the number of containers returned
		Size:    false,             // Don't return size info (set to true if needed)
		Filters: filters.NewArgs(), // Properly initialized Filters
	})
	if err != nil {
		panic(err)
	}

	// Convert Docker container data into table rows
	rows := make([]table.Row, 0, len(containers))
	for _, ctr := range containers {
		names := strings.Join(ctr.Names, ", ")
		command := fmt.Sprintf("%.20s", ctr.Command)
		created := time.Unix(ctr.Created, 0).Format("2006-01-02 15:04:05")

		var portInfo []string
		for _, port := range ctr.Ports {
			portInfo = append(portInfo, fmt.Sprintf("%d->%d/%s", port.PrivatePort, port.PublicPort, port.Type))
		}
		ports := strings.Join(portInfo, ", ")

		// Append a row containing container details
		rows = append(rows, table.Row{
			ctr.ID[:12],
			ctr.Image,
			command,
			created,
			ctr.Status,
			ports,
			names,
		})
	}

	return rows
}
