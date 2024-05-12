package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// NewProjectCmd .
var NewProjectCmd = &cobra.Command{
	Use:   "new-project [project_name]",
	Short: "New a microservice project based on freedom",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		sysPath, err := filepath.Abs(args[0])
		if err != nil {
			return
		}
		mkdirAll(sysPath)
		return
	},
}

func init() {
	AddCommand(NewProjectCmd)
}

func mkdirAll(projectPath string) {
	os.MkdirAll(projectPath+"/server", os.ModePerm)
	os.MkdirAll(projectPath+"/server/conf", os.ModePerm)
	os.MkdirAll(projectPath+"/adapter", os.ModePerm)
	os.MkdirAll(projectPath+"/adapter/controller", os.ModePerm)
	os.MkdirAll(projectPath+"/adapter/repository", os.ModePerm)
	os.MkdirAll(projectPath+"/domain", os.ModePerm)
	os.MkdirAll(projectPath+"/domain/aggregate", os.ModePerm)
	os.MkdirAll(projectPath+"/domain/entity", os.ModePerm)
	os.MkdirAll(projectPath+"/domain/vo", os.ModePerm)
	os.MkdirAll(projectPath+"/domain/po", os.ModePerm)
	os.MkdirAll(projectPath+"/domain/event", os.ModePerm)
	os.MkdirAll(projectPath+"/infra", os.ModePerm)
}
