package wordle

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish/bubbletea"
)

func Handler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	renderer := bubbletea.MakeRenderer(s)

	blockStyle := renderer.NewStyle().Foreground(lipgloss.Color("15")).Border(lipgloss.NormalBorder(), true)
	style := style{
		renderer:     renderer,
		correctStyle: blockStyle.Foreground(lipgloss.Color("2")).BorderForeground(lipgloss.Color("2")),
		presentStyle: blockStyle.Foreground(lipgloss.Color("3")).BorderForeground(lipgloss.Color("3")),
		missingStyle: blockStyle.Foreground(lipgloss.Color("8")).BorderForeground(lipgloss.Color("8")),
		emptyStyle:   blockStyle,
	}

	m := model{
		style: &style,
	}
	return m, []tea.ProgramOption{tea.WithAltScreen()}
}
