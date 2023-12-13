package editor

import "errors"

type Mode struct {
	name string
}

var (
	Unknown = Mode{}
	Normal  = Mode{"normal"}
	Visual  = Mode{"visual"}
	Command = Mode{"command"}
	Insert  = Mode{"insert"}
)

func GetMode(s string) (Mode, error) {
	switch s {
	case Normal.name:
		return Normal, nil
	case Visual.name:
		return Visual, nil
	case Command.name:
		return Command, nil
	case Insert.name:
		return Insert, nil
	}

	return Unknown, errors.New("Unknown mode: " + s)
}

func (m *Mode) ToString() string {
	return m.name
}
