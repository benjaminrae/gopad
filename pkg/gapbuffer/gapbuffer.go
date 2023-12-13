package gapbuffer

import (
	"strings"
	"unicode"
)

type GapBuffer struct {
	Buffer   []rune
	Size     int
	GapStart int
	GapSize  int
}

var nullRune = '\x00'

func New(size int) GapBuffer {
	gapBuffer := GapBuffer{
		Buffer:   make([]rune, size),
		Size:     size,
		GapStart: 0,
		GapSize:  size,
	}

	return gapBuffer
}

func (g *GapBuffer) Grow() {
	newSize := g.Size * 2
	newBuffer := make([]rune, newSize)

	for i := 0; i < g.GapStart; i++ {
		newBuffer[i] = g.Buffer[i]
	}

	for i := g.GapStart + g.GapSize; i < g.Size; i++ {
		newBuffer[i] = g.Buffer[i]
	}

	g.Buffer = newBuffer
	g.Size = newSize
}

func (g *GapBuffer) StepLeft() {
	if g.GapStart == 0 {
		return
	}

	gapEnd := g.GapStart + g.GapSize

	g.Buffer[gapEnd-1] = g.Buffer[g.GapStart-1]

	g.GapStart--

	g.Buffer[g.GapStart] = rune(nullRune)
}

func (g *GapBuffer) StepRight() {
	if g.GapStart+g.GapSize == g.Size {
		return
	}

	gapEnd := g.GapStart + g.GapSize

	g.Buffer[g.GapStart] = g.Buffer[gapEnd]

	g.GapStart++
	g.Buffer[gapEnd] = rune(nullRune)
}

func (g *GapBuffer) MoveLeft(distance int) {
	for i := 0; i < distance; i++ {
		g.StepLeft()
	}
}

func (g *GapBuffer) MoveRight(distance int) {
	for i := 0; i < distance; i++ {
		g.StepRight()
	}
}

func (g *GapBuffer) InsertRune(rune rune) {
	if g.GapSize == 0 {
		g.Grow()
	}

	g.Buffer[g.GapStart] = rune
	g.GapStart++
	g.GapSize--
}

func (g *GapBuffer) InsertString(string string) {
	for _, rune := range string {
		g.InsertRune(rune)
	}
}

func (g *GapBuffer) DeleteLeft() {
	if g.GapStart == 0 {
		return
	}

	g.Buffer[g.GapStart-1] = rune(nullRune)
	g.GapSize++
	g.GapStart--
}

func (g *GapBuffer) DeleteRight() {
	gapEnd := g.GapStart + g.GapSize

	if gapEnd == g.Size {
		return
	}

	g.Buffer[gapEnd] = rune(nullRune)
	g.GapSize++
}

func (g *GapBuffer) ToString() string {
	s := strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, string(g.Buffer))

	return s
}
