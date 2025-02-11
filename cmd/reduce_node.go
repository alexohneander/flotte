package cmd

import (
	reducenode "github.com/alexohneander/flotte/internal/reduce-node"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(reducenodeCmd)
}

var reducenodeCmd = &cobra.Command{
	Use:   "reduce-node",
	Short: "Start a reduce node",
	Long:  `Long`,
	Run: func(cmd *cobra.Command, args []string) {
		reducenode.Start()
	},
}
