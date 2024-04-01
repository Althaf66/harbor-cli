package list

import (
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).Padding(0, 1)

type model struct {
	table table.Model
}

var columns = []table.Column{
	{Title: "Project Name", Width: 12},
	{Title: "Access Level", Width: 12},
	// {Title: "Role", Width: 12},
	// {Title: "Type", Width: 12},
	{Title: "Repo Count", Width: 12},
	{Title: "Creation Time", Width: 30},
}

func (m model) Init() tea.Cmd {
	return tea.Quit
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func ListProjects(projects []*models.Project) {
	var rows []table.Row
	for _, project := range projects {
		accessLevel := "public"
		if project.Metadata.Public != "true" {
			accessLevel = "private"
		}
		rows = append(rows, table.Row{
			project.Name, // Project Name
			accessLevel,  // Access Level
			strconv.FormatInt(project.RepoCount, 10),
			project.CreationTime.String(), // Creation Time
		})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(len(rows)),
	)

	// Set the styles for the table
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		Bold(false)

	s.Selected = s.Selected.
		Foreground(s.Cell.GetForeground()).
		Background(s.Cell.GetBackground()).
		Bold(false)
	t.SetStyles(s)

	m := model{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
