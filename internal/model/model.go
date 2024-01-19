package model

import (
	"fmt"
	"os"
	"time"

	"github.com/Lameorc/tomatotea/internal/config"
	"github.com/Lameorc/tomatotea/internal/pomodoro"
	"github.com/Lameorc/tomatotea/internal/types"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

type pomodoroModel struct {
	pomodoro      types.Pomodoro
	time          timer.Model
	roundsElapsed int
	cfg           *config.DurationConfig
}

func (m *pomodoroModel) timeoutReached() {
	m.pomodoro.Advance()

	var timeout time.Duration
	switch m.pomodoro.State() {
	case types.Working:
		timeout = m.cfg.Work
	case types.ShortBreak:
		timeout = m.cfg.Break
	case types.LongBreak:
		timeout = m.cfg.Rest
	}

	m.time.Timeout = timeout
}

// Update implements tea.Model.
func (m *pomodoroModel) Update(tm tea.Msg) (tea.Model, tea.Cmd) {
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
func (m *pomodoroModel) View() string {
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
func (m *pomodoroModel) Init() tea.Cmd {
	return m.time.Init()
}

var _ tea.Model = (*pomodoroModel)(nil)

func newModel(c *config.DurationConfig, t timer.Model) *pomodoroModel {
	return &pomodoroModel{
		pomodoro:      pomodoro.New(),
		time:          t,
		roundsElapsed: 0,
		cfg:           c,
	}
}

// Run runs the bubbletea program defined by models in this package
func Run(c *config.Config) {
	p := tea.NewProgram(newModel(&c.Durations, timer.New(c.Durations.Work)))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Program run failed: %v", err)
		os.Exit(1)
	}

}
