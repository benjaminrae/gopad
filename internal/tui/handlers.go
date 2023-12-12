package tui

import (
	"github.com/benjaminrae/gopad/internal/editor"
	"github.com/gdamore/tcell/v2"
)

type RuneHandler func(t TerminalUI, e editor.Editor)

type KeyHandler func(t TerminalUI, e editor.Editor)

type KeyHandlers map[tcell.Key]KeyHandler

type RuneHandlers map[rune]RuneHandler

var NormalRuneHandlers = RuneHandlers{
	rune('i'): func(t TerminalUI, e editor.Editor) {
		e.SetInsertMode()
	},
}

var InsertKeyHandlers = KeyHandlers{
	tcell.KeyEscape: func(t TerminalUI, e editor.Editor) {
		e.SetNormalMode()
	},
}

var InsertModeRuneHandlers = RuneHandlers{}
