package config

import (
	"time"

	"github.com/spf13/viper"
)

const (
	WorkDuration  = "work-duration"
	BreakDuration = "break-duration"
	RestDuration  = "rest-duration"
)

type DurationConfig struct {
	Work  time.Duration
	Break time.Duration
	Rest  time.Duration
}

type Config struct {
	Durations DurationConfig
}

func FromViper() *Config {
	return &Config{
		Durations: DurationConfig{
			Work:  viper.GetDuration(WorkDuration),
			Break: viper.GetDuration(BreakDuration),
			Rest:  viper.GetDuration(RestDuration),
		},
	}
}
