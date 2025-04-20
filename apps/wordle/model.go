package wordle

import (
	"slices"
	"unicode"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type keyState int

const (
	unknownKey keyState = iota
	correctKey
	missingKey
	presentKey
)

type style struct {
	renderer     *lipgloss.Renderer
	missingStyle lipgloss.Style
	correctStyle lipgloss.Style
	presentStyle lipgloss.Style
	unknownStyle lipgloss.Style
}

type key struct {
	value rune
	state keyState
}

type model struct {
	grid         [5][5]key
	word         string
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
		case tea.KeyEnter:
			if m.currentCol == len(m.grid[0]) {
				m.validateCurrentRow()
				m.currentRow++
				m.currentCol = 0
			}
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
		state: unknownKey,
	}
	m.currentCol++
}

func (m *model) validateCurrentRow() {
	wordRunes := []rune(m.word)
	for i, k := range m.grid[m.currentRow] {
		currKey := m.grid[m.currentRow][i]
		if i == slices.Index(wordRunes, k.value) {
			m.grid[m.currentRow][i].state = correctKey
			wordRunes[i] = -1
		} else if slices.Contains(wordRunes, currKey.value) {
			m.grid[m.currentRow][i].state = presentKey
		} else {
			m.grid[m.currentRow][i].state = missingKey
		}
	}
}
