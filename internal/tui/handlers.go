package tui

import (
	"github.com/benjaminrae/gopad/internal/editor"
	"github.com/gdamore/tcell/v2"
)

type Handler func(t *TerminalUI, e *editor.Editor)

type KeyHandlers map[tcell.Key]Handler

type RuneHandlers map[rune]Handler

func InsertRuneHandler(t *TerminalUI, e *editor.Editor, ev *tcell.EventKey) {
	e.Insert(ev.Rune())
	t.Screen.ShowCursor(e.CurrentBuffer.GapStart, 1)
}

var NormalRuneHandlers = RuneHandlers{
	rune('i'): func(t *TerminalUI, e *editor.Editor) {
		e.SetInsertMode()
		t.Screen.SetCursorStyle(tcell.CursorStyleBlinkingBar)
	},
	'v': func(t *TerminalUI, e *editor.Editor) {
		e.SetVisualMode()
		t.Screen.SetCursorStyle(tcell.CursorStyleBlinkingBlock)
	},
	':': func(t *TerminalUI, e *editor.Editor) {
		e.SetCommandMode()
	},
}

var NormalKeyHandlers = KeyHandlers{}

var InsertKeyHandlers = KeyHandlers{
	tcell.KeyEscape: func(t *TerminalUI, e *editor.Editor) {
		e.SetNormalMode()
		t.Screen.SetCursorStyle(tcell.CursorStyleBlinkingBlock)
	},
	tcell.KeyBackspace: func(t *TerminalUI, e *editor.Editor) {
		e.CurrentBuffer.DeleteLeft()
		t.Screen.ShowCursor(e.CurrentBuffer.GapStart, 1)
	},
	tcell.KeyBackspace2: func(t *TerminalUI, e *editor.Editor) {
		e.CurrentBuffer.DeleteLeft()
		t.Screen.ShowCursor(e.CurrentBuffer.GapStart, 1)
	},
	tcell.KeyDelete: func(t *TerminalUI, e *editor.Editor) {
		e.CurrentBuffer.DeleteRight()
		t.Screen.ShowCursor(e.CurrentBuffer.GapStart, 1)
	},
	tcell.KeyLeft: func(t *TerminalUI, e *editor.Editor) {
		e.CurrentBuffer.StepLeft()
		t.Screen.ShowCursor(e.CurrentBuffer.GapStart, 1)
	},
	tcell.KeyRight: func(t *TerminalUI, e *editor.Editor) {
		e.CurrentBuffer.StepRight()
		t.Screen.ShowCursor(e.CurrentBuffer.GapStart, 1)
	},
}

var InsertRuneHandlers = KeyHandlers{
	tcell.RuneLArrow: func(t *TerminalUI, e *editor.Editor) {
		e.CurrentBuffer.StepLeft()
	},
	tcell.RuneRArrow: func(t *TerminalUI, e *editor.Editor) {
		e.CurrentBuffer.StepRight()
	},
}

var CommandKeyHandlers = KeyHandlers{
	tcell.KeyEscape: func(t *TerminalUI, e *editor.Editor) {
		e.SetNormalMode()
		t.Screen.SetCursorStyle(tcell.CursorStyleBlinkingBlock)
	},
}

var CommandRuneHandlers = RuneHandlers{}

var VisualKeyHandlers = KeyHandlers{
	tcell.KeyEscape: func(t *TerminalUI, e *editor.Editor) {
		e.SetNormalMode()
		t.Screen.SetCursorStyle(tcell.CursorStyleBlinkingBlock)
	},
}

var VisualRuneHandlers = RuneHandlers{}
