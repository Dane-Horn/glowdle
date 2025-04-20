package wordle

import "github.com/charmbracelet/lipgloss"

func (m model) View() string {
	return m.viewGrid()
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
