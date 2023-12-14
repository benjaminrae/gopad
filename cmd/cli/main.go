package main

import (
	"os"

	"github.com/benjaminrae/gopad/internal/editor"
	"github.com/benjaminrae/gopad/internal/tui"
	"github.com/gdamore/tcell/v2"
)

func main() {
	tui.InitScreen()

	editor := editor.New()
	tui.Tui.Screen.SetCursorStyle(tcell.CursorStyleBlinkingBlock)
	tui.Tui.Screen.ShowCursor(editor.CurrentBuffer.GapStart, 1)

	for {
		tui.Tui.Screen.Clear()
		tui.Tui.DrawTextLine(1, editor.CurrentBuffer.ToString())
		tui.Tui.CreateStatusBar(editor.Mode.ToString())
		tui.Tui.Screen.Show()

		ev := tui.Tui.Screen.PollEvent()
		handleEvent(ev, &tui.Tui, &editor)
	}
}

func handleEvent(ev tcell.Event, t *tui.TerminalUI, e *editor.Editor) {
	switch ev := ev.(type) {
	case *tcell.EventResize:
		t.Screen.Sync()
	case *tcell.EventKey:

		if ev.Key() == tcell.KeyCtrlC {
			t.Screen.Fini()
			os.Exit(0)
			return
		}

		if e.Mode.ToString() == "normal" {
			handler := tui.NormalRuneHandlers[ev.Rune()]
			if handler != nil {
				handler(t, e)
				return
			}
			handler = tui.NormalKeyHandlers[ev.Key()]
			if handler != nil {
				handler(t, e)
				return
			}

		}

		if e.Mode.ToString() == "insert" {
			handler := tui.InsertKeyHandlers[ev.Key()]
			if handler != nil {
				handler(t, e)
				return
			}

			if ev.Key() == tcell.KeyRune {
				tui.InsertRuneHandler(t, e, ev)
				return
			}
		}

		if e.Mode.ToString() == "command" {
			handler := tui.CommandKeyHandlers[ev.Key()]
			if handler != nil {
				handler(t, e)
				return
			}

			handler = tui.CommandRuneHandlers[ev.Rune()]
			if handler != nil {
				handler(t, e)
				return
			}
		}

		if e.Mode.ToString() == "visual" {
			handler := tui.VisualKeyHandlers[ev.Key()]
			if handler != nil {
				handler(t, e)
				return
			}

			handler = tui.VisualRuneHandlers[ev.Rune()]
			if handler != nil {
				handler(t, e)
				return
			}
		}
	}

}
