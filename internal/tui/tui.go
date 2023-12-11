package tui

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

type TerminalUI struct {
	Screen tcell.Screen
}

var Tui = TerminalUI{}

func InitScreen() {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	s, err := tcell.NewScreen()

	if err != nil {
		log.Panic(err)
	}

	if err := s.Init(); err != nil {
		log.Panic(err)
	}

	s.Clear()
	s.SetStyle(tcell.StyleDefault)
	s.SetCursorStyle(tcell.CursorStyleBlinkingUnderline)
	s.ShowCursor(1, 1)

	Tui.Screen = s
}

func (t *TerminalUI) DrawTextLine(row int, text string) {
	for i, rune := range []rune(text) {
		t.Screen.SetContent(1+i, row, rune, nil, tcell.StyleDefault)
	}
}
