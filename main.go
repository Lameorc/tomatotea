package main

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	workTime     = 25 * time.Minute
	breakTime    = 5 * time.Minute
	bigBreakTime = 15 * time.Minute
	// NOTE: debug values
	// breakTime = 2 * time.Second
	// workTime = 5 * time.Second
	// bigBreakTime           = 3 * time.Second
	intervalsUntilBigBreak = 4
)

type workState int

const (
	working workState = iota
	resting
)

type model struct {
	stat          workState
	time          timer.Model
	roundsElapsed int
	lastTimeout   time.Duration
}

func (m *model) onLongBreak() bool {
	// TODO: better handling of this, it will cause issues if the normal resting tim is equal to long break
	return m.lastTimeout == bigBreakTime
}

func (m *model) timeoutReached() tea.Cmd {
	var timeout time.Duration
	switch m.stat {
	case working:
		if m.roundsElapsed == intervalsUntilBigBreak-1 {
			timeout = bigBreakTime
			m.roundsElapsed = 0
		} else {
			timeout = breakTime
		}
		m.stat = resting
	case resting:
		timeout = workTime
		m.stat = working
		if !m.onLongBreak() {
			m.roundsElapsed++
		}
	default:
		panic(fmt.Sprintf("unexpected model state: %d", m.stat))
	}

	m.lastTimeout = timeout
	m.time = timer.New(timeout)
	return m.time.Init()
}

// Update implements tea.Model.
func (m *model) Update(tm tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := tm.(type) {
	case timer.TickMsg:
		var cmd tea.Cmd
		m.time, cmd = m.time.Update(msg)
		return m, cmd
	case timer.StartStopMsg:
		var cmd tea.Cmd
		m.time, cmd = m.time.Update(msg)
		return m, cmd

	case timer.TimeoutMsg:
		cmd := m.timeoutReached()
		return m, cmd

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

// View implements tea.Model.
func (m *model) View() string {
	// header
	s := "🍅: "
	switch m.stat {
	case working:
		s += "working"
	case resting:
		s += "resting"
		if m.onLongBreak() {
			s += " (long)"
		}
	}
	if !m.onLongBreak() {
		s += fmt.Sprintf(" (%d/%d)", m.roundsElapsed+1, intervalsUntilBigBreak)
	}

	s += "\n"
	s += m.time.View()

	// help
	s += "\npress ctrl+c or 'q' to quit"

	return s
}

// Init implements tea.Model
func (m *model) Init() tea.Cmd {
	return m.time.Init()
}

var _ tea.Model = (*model)(nil)

func newModel(t timer.Model) *model {
	return &model{
		stat:          working,
		time:          t,
		roundsElapsed: 0,
	}
}

func main() {
	p := tea.NewProgram(newModel(timer.New(workTime)))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Program run failed: %v", err)
		os.Exit(1)
	}

}
