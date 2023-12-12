package editor

import "log"

type Editor struct {
	Mode Mode
}

func New() Editor {
	e := Editor{}

	e.SetNormalMode()

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
