package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Lameorc/tomatotea/internal/pomodoro"
	"github.com/Lameorc/tomatotea/internal/types"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	workTime     = 25 * time.Minute
	breakTime    = 5 * time.Minute
	bigBreakTime = 15 * time.Minute

	// NOTE: debug values
	// breakTime    = 2 * time.Second
	// workTime     = 5 * time.Second
	// bigBreakTime = 3 * time.Second
)

type model struct {
	pomodoro      types.Pomodoro
	time          timer.Model
	roundsElapsed int
}

func (m *model) onLongBreak() bool {
	return m.pomodoro.State() == types.LongBreak
}

func (m *model) timeoutReached() {
	m.pomodoro.Advance()

	var timeout time.Duration
	switch m.pomodoro.State() {
	case types.Working:
		timeout = workTime
	case types.ShortBreak:
		timeout = breakTime
	case types.LongBreak:
		timeout = bigBreakTime
	}

	m.time.Timeout = timeout
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
		m.timeoutReached()
		return m, nil

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
	s += m.pomodoro.String()

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
		pomodoro:      pomodoro.New(),
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
