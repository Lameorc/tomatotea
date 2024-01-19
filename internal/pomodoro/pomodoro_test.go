package pomodoro

import (
	"testing"

	"github.com/Lameorc/tomatotea/internal/types"
	"github.com/stretchr/testify/require"
)

func TestPomodoro(t *testing.T) {
	i := New()
	p := i.(*pomodoro)

	shortBreaks := 0
	for i := interval(0); i < maxIntervals-1; i++ {
		require.Equal(t, types.Working, p.State())
		require.Equal(t, i, p.currentInterval)

		p.Advance()
		require.Equal(t, types.ShortBreak, p.State())
		require.Equal(t, i, p.currentInterval)
		shortBreaks++

		p.Advance()
	}
	require.Equal(t, 3, shortBreaks)

	p.Advance()
	require.Equal(t, types.LongBreak, p.State())
	require.Equal(t, maxIntervals-1, p.currentInterval)

	p.Advance()
	require.Equal(t, types.Working, p.State())
	require.Equal(t, interval(0), p.currentInterval)
}
