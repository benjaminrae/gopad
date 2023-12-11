package main

import (
	"github.com/benjaminrae/gopad/internal/tui"
	"github.com/gdamore/tcell/v2"
)

func main() {
	tui.InitScreen()

	tui.Tui.DrawTextLine(1, "Hello world")

	for {
		tui.Tui.Screen.Show()
		tui.Tui.Screen.SetCursorStyle(tcell.CursorStyleBlinkingBlock)

		ev := tui.Tui.Screen.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			tui.Tui.Screen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyCtrlC {
				return
			}

		}
	}
}
