package editor

import (
	"log"

	"github.com/benjaminrae/gopad/pkg/gapbuffer"
)

type Editor struct {
	Mode          Mode
	CurrentBuffer gapbuffer.GapBuffer
}

func New() Editor {
	e := Editor{}

	e.SetNormalMode()
	e.initialiseEmptyBuffer()

	return e
}

func (e *Editor) SetNormalMode() {
	mode, err := GetMode("normal")

	if err != nil {
		log.Fatal(err)
	}

	e.Mode = mode
}

func (e *Editor) SetInsertMode() {
	mode, err := GetMode("insert")

	if err != nil {
		log.Fatal(err)
	}

	e.Mode = mode
}

func (e *Editor) SetCommandMode() {
	mode, err := GetMode("command")

	if err != nil {
		log.Fatal(err)
	}

	e.Mode = mode
}

func (e *Editor) SetVisualMode() {
	mode, err := GetMode("visual")

	if err != nil {
		log.Fatal(err)
	}

	e.Mode = mode
}

func (e *Editor) Insert(char rune) {
	e.CurrentBuffer.InsertRune(char)
}

func (e *Editor) initialiseEmptyBuffer() {
	e.CurrentBuffer = gapbuffer.New(80)
}
