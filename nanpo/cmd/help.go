package cmd

import (
	"github.com/spf13/cobra"
)

// HelpCmd .
var HelpCmd = &cobra.Command{
	Use: "help",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		return
	},
}

func init() {
	AddCommand(HelpCmd)
}
