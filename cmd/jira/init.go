package jira

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	jiraConfig "github.com/ankitpokhrel/jira-cli/internal/config"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Init initializes jira config",
	Long:    "Init initializes jira configuration required for the tool to work properly.",
	Aliases: []string{"initialize", "configure", "config", "setup"},
	Run:     initialize,
}

func initialize(*cobra.Command, []string) {
	c := jiraConfig.NewJiraCLIConfig()

	if err := c.Generate(); err != nil {
		if err == jiraConfig.ErrSkip {
			fmt.Printf("\n\033[0;32m✓\033[0m Skipping config generation. Current config: %s\n", viper.ConfigFileUsed())
		} else {
			fmt.Printf("\n\033[0;31m✗\033[0m Unable to generate configuration: %s\n", viper.ConfigFileUsed())
		}

		os.Exit(1)
	}

	fmt.Printf("\n\033[0;32m✓\033[0m Configuration generated: %s\n", viper.ConfigFileUsed())
}

func init() {
	rootCmd.AddCommand(initCmd)
}
