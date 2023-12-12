package tui

func (t *TerminalUI) CreateStatusBar(status string) {
	_, height := t.Screen.Size()

	t.DrawTextLine(height-1, status)
}
