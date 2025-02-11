package cmd

import (
	mapnode "github.com/alexohneander/flotte/internal/map-node"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(mapnodeCmd)
}

var mapnodeCmd = &cobra.Command{
	Use:   "map-node",
	Short: "Start a map node",
	Long:  `Long`,
	Run: func(cmd *cobra.Command, args []string) {
		mapnode.Start()
	},
}
