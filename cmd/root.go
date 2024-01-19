/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/Lameorc/tomatotea/internal/config"
	"github.com/Lameorc/tomatotea/internal/model"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tomatotea",
	Short: "A pomodoro timer",
	Long: `TomatoTea is a CLI application for running a pomodoro timer

The duration of the various periods can be configured with the relevant flags.
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) { model.Run(config.FromViper()) },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

const (
	workDefault  = 25 * time.Minute
	breakDefault = 5 * time.Minute
	restDefault  = 15 * time.Minute
)

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tomatotea.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().DurationP("work-duration", "w", workDefault, "duration of the work period")
	rootCmd.Flags().DurationP("break-duration", "b", breakDefault, "duration of the standard break between work sessions")
	rootCmd.Flags().DurationP("rest-duration", "r", restDefault, "duration of the period between a set of intervals")
	if err := viper.BindPFlags(rootCmd.Flags()); err != nil {
		panic("failed to bind pflags to viper")
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".tomatotea" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".tomatotea")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
