package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "flotte",
	Short: "Flotte is a MapReduce System for running distributed computations on large datasets.",
	Run: func(cmd *cobra.Command, args []string) {
		// Check if the user has provided a subcommand
		if len(args) == 0 {
			cmd.Help()
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
