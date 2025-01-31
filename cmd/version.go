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
	Short: "Print the version number of nerdlocke",
	Long:  `All software has versions. This is nerdlocke's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("nerdlocke v0.1.0")
	},
}
