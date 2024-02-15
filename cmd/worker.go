package cmd

import (
	"fmt"

	"github.com/alexohneander/flotte/internal/worker"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(workerCmd)
}

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Start Flotte as a worker",
	Long:  `Start a Worker to process jobs from the queue`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting Worker...")
		worker.StartWorker()
	},
}
