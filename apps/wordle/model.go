package wordle

import (
	"unicode"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type keyState int

const (
	emptyKey keyState = iota
	correctKey
	missingKey
	presentKey
)

type style struct {
	renderer     *lipgloss.Renderer
	missingStyle lipgloss.Style
	correctStyle lipgloss.Style
	presentStyle lipgloss.Style
	emptyStyle   lipgloss.Style
}

type key struct {
	value rune
	state keyState
}

type model struct {
	grid         [5][5]key
	gameOver     bool
	currentRow   int
	currentCol   int
	windowHeight int
	windowWidth  int
	*style
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyBackspace:
			m.removeChar()
			return m, nil
		case tea.KeyRunes:
			if len(msg.Runes) == 1 {
				m.acceptChar(msg.Runes[0])
				return m, nil
			}
		}
	}

	return m, nil
}

func (m *model) removeChar() {
	if m.currentCol <= 0 {
		return
	}
	m.currentCol--
	m.grid[m.currentRow][m.currentCol] = key{}
}

func (m *model) acceptChar(ch rune) {
	if m.currentCol == len(m.grid[0]) {
		return
	}
	m.grid[m.currentRow][m.currentCol] = key{
		value: unicode.ToUpper(ch),
		state: correctKey,
	}
	m.currentCol++
}
