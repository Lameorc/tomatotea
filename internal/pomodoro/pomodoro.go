package pomodoro

import (
	"fmt"

	"github.com/Lameorc/tomatotea/internal/types"
)

type interval int

const (
	maxIntervals interval = 4
)

type pomodoro struct {
	state           types.State
	currentInterval interval
}

func New() types.Pomodoro {
	return &pomodoro{
		state:           types.Working,
		currentInterval: 0,
	}
}

// Advance implements types.Pomodoro.
func (p *pomodoro) Advance() {
	switch p.state {
	case types.Working:
		if p.currentInterval == maxIntervals-1 {
			p.state = types.LongBreak
		} else {
			p.state = types.ShortBreak
		}
	case types.ShortBreak:
		p.currentInterval++
		p.state = types.Working
	case types.LongBreak:
		p.currentInterval = 0
		p.state = types.Working
	}
}

// State implements types.Pomodoro.
func (p *pomodoro) State() types.State {
	return p.state
}

// String implements types.Pomodoro.
func (p *pomodoro) String() string {
	return fmt.Sprintf("%s (%d/%d)", p.State().String(), p.currentInterval+1, maxIntervals)
}

var _ types.Pomodoro = (*pomodoro)(nil)
