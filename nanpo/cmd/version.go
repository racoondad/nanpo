package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	versionNum = "v1.0.0"
)

// VersionCmd .
var VersionCmd = &cobra.Command{
	Use:   "v",
	Short: "Output current version number",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		fmt.Println("nanpo " + versionNum)
		return
	},
}

func init() {
	rootCmd.AddCommand(VersionCmd)
}
