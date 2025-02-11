package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of flotte",
	Long:  `All software has versions. This is flotte's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("flotte v0.9 -- HEAD")
	},
}
