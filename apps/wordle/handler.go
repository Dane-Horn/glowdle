package wordle

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish/bubbletea"
)

func Handler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, _ := s.Pty()
	renderer := bubbletea.MakeRenderer(s)

	blockStyle := renderer.NewStyle().Foreground(lipgloss.Color("15")).Border(lipgloss.NormalBorder(), true)
	style := style{
		renderer:     renderer,
		correctStyle: blockStyle.Foreground(lipgloss.Color("2")).BorderForeground(lipgloss.Color("2")),
		presentStyle: blockStyle.Foreground(lipgloss.Color("3")).BorderForeground(lipgloss.Color("3")),
		missingStyle: blockStyle.Foreground(lipgloss.Color("8")).BorderForeground(lipgloss.Color("8")),
		unknownStyle: blockStyle,
	}

	m := model{
		windowWidth:  pty.Window.Width,
		windowHeight: pty.Window.Height,
		style:        &style,
		word:         "SASSY",
		wordLength:   5,
		numGuesses:   10,
		grid:         initialGrid(5, 10),
		keyboard:     initialKeyboard(),
	}
	return m, []tea.ProgramOption{tea.WithAltScreen()}
}

func initialGrid(wordLength int, guesses int) [][]key {
	grid := make([][]key, guesses)
	for i := range guesses {
		grid[i] = make([]key, wordLength)
	}
	return grid
}

func initialKeyboard() [][]key {
	return [][]key{
		{{value: 'Q'}, {value: 'W'}, {value: 'E'}, {value: 'R'}, {value: 'T'}, {value: 'Y'}, {value: 'U'}, {value: 'I'}, {value: 'O'}, {value: 'P'}},
		{{value: 'A'}, {value: 'S'}, {value: 'D'}, {value: 'F'}, {value: 'G'}, {value: 'H'}, {value: 'J'}, {value: 'K'}, {value: 'L'}},
		{{value: 'Z'}, {value: 'X'}, {value: 'C'}, {value: 'V'}, {value: 'B'}, {value: 'N'}, {value: 'M'}},
	}
}
