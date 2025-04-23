package wordle

import (
	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	gridView := m.viewGrid()
	keyboardView := m.viewKeyboard()
	view := lipgloss.JoinVertical(lipgloss.Center, gridView, keyboardView)
	return lipgloss.Place(m.windowWidth, m.windowHeight, lipgloss.Center, lipgloss.Center, view)
}

func (m *model) viewGrid() string {
	rowViews := make([]string, len(m.grid))
	for i, row := range m.grid {
		rowViews[i] = m.viewRow(row[:])
	}
	return lipgloss.JoinVertical(lipgloss.Left, rowViews...)
}

func (m *model) viewRow(row []key) string {
	keyViews := make([]string, len(row))
	for i, k := range row {
		keyViews[i] = m.viewKey(k)
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, keyViews...)
}

func (m *model) viewKey(key key) string {
	style := m.unknownStyle

	switch key.state {
	case correctKey:
		style = m.correctStyle
	case missingKey:
		style = m.missingStyle
	case presentKey:
		style = m.presentStyle
	}
	if key.value == 0 {
		return style.Render(" ")
	}
	return style.Render(string(key.value))
}

func (m *model) viewKeyboard() string {
	topRow := m.viewRow(m.keyboard[0])
	middleRow := m.viewRow(m.keyboard[1])
	bottomRow := m.viewRow(m.keyboard[2])
	keyboard := lipgloss.JoinVertical(
		lipgloss.Center,
		lipgloss.NewStyle().Padding(0, 2).Render(topRow),
		lipgloss.NewStyle().Padding(0, 4).Render(middleRow),
		lipgloss.NewStyle().Padding(0, 0).Render(bottomRow),
	)
	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(m.unknownStyle.GetForeground()).
		Padding(0, 1).
		Render(keyboard)
}
