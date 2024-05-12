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
