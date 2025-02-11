package cmd

import (
	generalnode "github.com/alexohneander/flotte/pkg/general-node"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generalnodeCmd)
}

var generalnodeCmd = &cobra.Command{
	Use:   "general-node",
	Short: "Start a general node",
	Long:  `Long`,
	Run: func(cmd *cobra.Command, args []string) {
		generalnode.Start()
	},
}
