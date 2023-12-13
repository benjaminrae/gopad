package tui

import (
	"github.com/benjaminrae/gopad/internal/editor"
	"github.com/gdamore/tcell/v2"
)

type Handler func(t *TerminalUI, e *editor.Editor)

type KeyHandlers map[tcell.Key]Handler

type RuneHandlers map[rune]Handler

var NormalRuneHandlers = RuneHandlers{
	rune('i'): func(t *TerminalUI, e *editor.Editor) {
		e.SetInsertMode()
	},
	'v': func(t *TerminalUI, e *editor.Editor) {
		e.SetVisualMode()
	},
	':': func(t *TerminalUI, e *editor.Editor) {
		e.SetCommandMode()
	},
}

var NormalKeyHandlers = KeyHandlers{}

var InsertKeyHandlers = KeyHandlers{
	tcell.KeyEscape: func(t *TerminalUI, e *editor.Editor) {
		e.SetNormalMode()
	},
}

var InsertRuneHandlers = RuneHandlers{}

var CommandKeyHandlers = KeyHandlers{
	tcell.KeyEscape: func(t *TerminalUI, e *editor.Editor) {
		e.SetNormalMode()
	},
}

var CommandRuneHandlers = RuneHandlers{}

var VisualKeyHandlers = KeyHandlers{
	tcell.KeyEscape: func(t *TerminalUI, e *editor.Editor) {
		e.SetNormalMode()
	},
}

var VisualRuneHandlers = RuneHandlers{}
