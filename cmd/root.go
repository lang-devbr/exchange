package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "exchange",
	Short: "exchange root command",
	Long:  "exchange root command",
}

// Execute root command
func Execute() {
	defer func() {
		err := recover()
		if err != nil {
			log.Fatalf("unexpected error while executing command %v", err)
		}
	}()

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal("error while executing command")
	}
}
