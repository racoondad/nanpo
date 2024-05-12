package cmd

import (
	"os"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "自动代码工具",
	Short: "auto coding",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		color.Error.Println(err.Error())
		os.Exit(1)
	}
}

func init() {
}

// AddCommand add sub command to root
func AddCommand(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}

// Commands returns all sub command
func Commands() []*cobra.Command {
	return rootCmd.Commands()
}
