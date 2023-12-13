package gapbuffer_test

import (
	"testing"

	"github.com/benjaminrae/gopad/pkg/gapbuffer"
)

func TestNew(t *testing.T) {
	var size = 10

	result := gapbuffer.New(size)

	if result.Size != size {
		t.Errorf("result.Size = %d, want %d", result.Size, size)
	}

	if result.Buffer == nil {
		t.Errorf("result.Buffer = %d, want %d", result.Buffer, make([]rune, size))
	}

	if result.ToString() != "" {
		t.Errorf("result.ToString = %s, want \"\"", result.ToString())
	}

	if result.GapSize != size {
		t.Errorf("result.GapSize = %v, want %v", result.GapSize, size)
	}

	if result.GapStart != 0 {
		t.Errorf("result.GapStart = %v, want %v", result.GapStart, 0)
	}
}

func TestInsertRune(t *testing.T) {
	result := gapbuffer.New(10)

	result.InsertRune('h')

	if result.ToString() != "h" {
		t.Errorf("result.ToString() = %s, want %s", result.ToString(), "h")
	}

	if result.GapStart != 1 {
		t.Errorf("result.GapStart = %v, want 1", result.GapStart)
	}

	if result.GapSize != 9 {
		t.Errorf("result.GapSize = %v, want 9", result.GapSize)
	}

}

func TestInsertString(t *testing.T) {
	var size = 10
	result := gapbuffer.New(size)
	var input = "Hello"

	result.InsertString(input)

	if result.ToString() != input {
		t.Errorf("result.ToString() = %s, want %s", result.ToString(), input)
	}

	if result.GapStart != len(input) {
		t.Errorf("result.GapStart = %v, want %v", result.GapStart, len(input))
	}

	if result.GapSize != size-len(input) {
		t.Errorf("result.GapSize = %v, want %v", result.GapSize, size-len(input))
	}
}

func TestGrow(t *testing.T) {
	var size = 10
	var expectedSize = size * 2
	gb := gapbuffer.New(size)

	gb.Grow()

	resultSize := gb.Size

	if resultSize != expectedSize {
		t.Errorf("result.Size = %v, want %v", resultSize, expectedSize)
	}
}

func TestStep(t *testing.T) {
	var size = 10
	gb := gapbuffer.New(size)
	var input = "Hello"
	gb.InsertString(input)
	var initialGapStart = gb.GapStart
	var expectedGapStart = initialGapStart - 1

	gb.StepLeft()

	var resultGapStart = gb.GapStart

	if resultGapStart != expectedGapStart {
		t.Errorf("result.GapStart = %v, want %v", resultGapStart, expectedGapStart)
	}

	gb.StepRight()

	resultGapStart = gb.GapStart

	if resultGapStart != initialGapStart {
		t.Errorf("result.GapStart = %v, want %v", resultGapStart, expectedGapStart)
	}

}

func TestDeleteLeft(t *testing.T) {
	var size = 10
	gb := gapbuffer.New(size)
	var input = "Hello"
	gb.InsertString(input)
	var expectedString = "Hell"

	gb.DeleteLeft()

	var result = gb.ToString()

	if result != expectedString {
		t.Errorf("result.ToString = %s, want %s", result, expectedString)
	}
}

func TestDeleteRight(t *testing.T) {
	var size = 10
	gb := gapbuffer.New(size)
	var input = "Hello"
	gb.InsertString(input)
	var expectedString = "Hell"

	gb.StepLeft()
	gb.DeleteRight()

	var result = gb.ToString()

	if result != expectedString {
		t.Errorf("result.ToString = %s, want %s", result, expectedString)
	}
}
