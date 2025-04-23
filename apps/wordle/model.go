package wordle

import (
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
	keyboard     [][]key
	grid         [][]key
	word         string
	wordLength   int
	numGuesses   int
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
	case tea.WindowSizeMsg:
		m.windowHeight = msg.Height
		m.windowWidth = msg.Width
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyBackspace:
			m.removeChar()
			return m, nil
		case tea.KeyEnter:
			if m.currentCol == len(m.grid[0]) && m.currentRow < len(m.grid) {
				m.validateCurrentRow()
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
	if m.currentCol >= m.wordLength || m.currentRow >= m.numGuesses {
		return
	}
	m.grid[m.currentRow][m.currentCol] = key{
		value: unicode.ToUpper(ch),
		state: unknownKey,
	}
	m.currentCol++
}

func (m *model) validateCurrentRow() {
	guess := m.grid[m.currentRow]
	validatedKeys, _ := validateWord(m.word, guess)
	m.grid[m.currentRow] = validatedKeys
	m.currentRow++
	m.currentCol = 0
	m.updateKeyboard(validatedKeys)
}

func (m *model) updateKeyboard(validatedKeys []key) {
	for _, k := range validatedKeys {
		for i := range m.keyboard {
			for j := range m.keyboard[i] {
				if m.keyboard[i][j].value == k.value {
					if m.keyboard[i][j].state != correctKey {
						m.keyboard[i][j].state = k.state
					}
				}
			}
		}
	}
}
