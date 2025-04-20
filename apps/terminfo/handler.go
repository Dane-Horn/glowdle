package terminfo

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish/bubbletea"
)

func Handler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, _ := s.Pty()
	renderer := bubbletea.MakeRenderer(s)
	txtStyle := renderer.NewStyle().Foreground(lipgloss.Color("10"))
	quitStyle := renderer.NewStyle().Foreground(lipgloss.Color("8"))
	bg := "light"
	if renderer.HasDarkBackground() {
		bg = "dark"
	}

	m := model{
		term:    pty.Term,
		profile: renderer.ColorProfile().Name(),
		width:   pty.Window.Width,
		height:  pty.Window.Height,
		bg:      bg,
		style: &style{
			renderer:  renderer,
			txtStyle:  txtStyle,
			quitStyle: quitStyle,
		},
	}
	return m, []tea.ProgramOption{tea.WithAltScreen()}
}
