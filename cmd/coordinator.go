package cmd

import (
	"fmt"

	"github.com/alexohneander/flotte/internal/coordinator"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(coordinatorCmd)
}

var coordinatorCmd = &cobra.Command{
	Use:   "coordinator",
	Short: "Start Flotte as a coordinator",
	Long:  `Start a Coordinator to manage the queue and distribute jobs to workers`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting Coordinator...")
		coordinator.StartCoordinator()
	},
}
