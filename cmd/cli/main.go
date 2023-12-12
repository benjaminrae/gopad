package main

import (
	"fmt"
	"os"

	"github.com/benjaminrae/gopad/internal/editor"
	"github.com/benjaminrae/gopad/internal/tui"
	"github.com/gdamore/tcell/v2"
)

func main() {
	tui.InitScreen()

	editor := editor.New()

	tui.Tui.DrawTextLine(1, "Hello world")

	for {
		tui.Tui.CreateStatusBar(editor.Mode.ToString())
		tui.Tui.Screen.Show()
		tui.Tui.Screen.SetCursorStyle(tcell.CursorStyleBlinkingBlock)

		ev := tui.Tui.Screen.PollEvent()
		fmt.Println(ev)
		handleEvent(ev, tui.Tui, editor)
	}
}

func handleEvent(ev tcell.Event, t tui.TerminalUI, e editor.Editor) {
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

			fmt.Println(ev.Rune())
			handler := tui.NormalRuneHandlers[ev.Rune()]

			if handler != nil {
				handler(t, e)
			}
		}

		if e.Mode.ToString() == "insert" {
			handler := tui.InsertKeyHandlers[ev.Key()]

			if handler != nil {
				handler(t, e)
			}
		}
	}

}
