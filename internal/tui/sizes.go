package tui

type Element struct {
	Height int
	Width  int
}

var StatusBar = Element{
	Height: 2,
}

var LineNumbers = Element{
	Width: 0,
}
