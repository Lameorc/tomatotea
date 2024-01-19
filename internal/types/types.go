package types

// Pomodoro defines a pomodoro timer which can cycle through various states of pomodoro technique
// See https://en.wikipedia.org/wiki/Pomodoro_Technique for details.
type Pomodoro interface {
	// State returns the current state of the pomodoro.
	State() State
	// Advance advances the pomodoro to the next state.
	Advance()
	// String returns the representation of the current pomodoro
	String() string
}

type State int

// String returns a textual representation of the current state
func (s State) String() string {
	switch s {
	case Working:
		return "working"
	case ShortBreak:
		return "break"
	case LongBreak:
		return "break (long)"
	}

	return "invalid state"
}

const (
	Working State = iota
	ShortBreak
	LongBreak
)
