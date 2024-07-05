package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/deepakjois/podscript/cmd/configure"
	"github.com/deepakjois/podscript/cmd/deepgram"
	"github.com/deepakjois/podscript/cmd/ytt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "podscript",
	Short: "podscript generates podcast audio transcripts",
	Long: `A tool to generated transcripts for podcast audio files using LLM and
Speech-To-Text (STT) APIs.`,
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(configure.Command)
	rootCmd.AddCommand(ytt.Command)
	rootCmd.AddCommand(deepgram.Command)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SilenceUsage = true
}

func initConfig() {
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.SetConfigType("toml")
	viper.SetConfigFile(path.Join(homeDir, ".podscript.toml"))

	viper.SetEnvPrefix("PODSCRIPT")
	viper.AutomaticEnv()

	// Read in config file and ENV variables if set
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// Config file was found but another error was produced
			fmt.Printf("Error reading config file: %s\n", err)
		}
	}
}

func Execute() error {
	return rootCmd.Execute()
}
