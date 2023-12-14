package tui

import (
	"fmt"

	"github.com/benjaminrae/gopad/internal/editor"
	"github.com/gdamore/tcell/v2"
)

type Handler func(t *TerminalUI, e *editor.Editor)

type KeyHandlers map[tcell.Key]Handler

type RuneHandlers map[rune]Handler

func InsertRuneHandler(t *TerminalUI, e *editor.Editor, ev *tcell.EventKey) {
	fmt.Println(ev.Rune())
	e.Insert(ev.Rune())
	t.Screen.ShowCursor(e.CurrentBuffer.GapStart, 1)
}

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
	tcell.KeyBackspace: func(t *TerminalUI, e *editor.Editor) {
		e.CurrentBuffer.DeleteLeft()
	},
	tcell.KeyInsert: func(t *TerminalUI, e *editor.Editor) {
		e.CurrentBuffer.DeleteRight()
	},
	tcell.KeyLeft: func(t *TerminalUI, e *editor.Editor) {
		e.CurrentBuffer.StepLeft()
	},
	tcell.KeyRight: func(t *TerminalUI, e *editor.Editor) {
		e.CurrentBuffer.StepRight()
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
	},
}

var CommandRuneHandlers = RuneHandlers{}

var VisualKeyHandlers = KeyHandlers{
	tcell.KeyEscape: func(t *TerminalUI, e *editor.Editor) {
		e.SetNormalMode()
	},
}

var VisualRuneHandlers = RuneHandlers{}
