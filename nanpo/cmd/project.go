package cmd

import (
	"errors"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/racoondad/nanpo/nanpo/template/project"
	"github.com/spf13/cobra"
)

// NewProjectCmd .
var NewProjectCmd = &cobra.Command{
	Use:   "new [project_name]",
	Short: "New project based on nanpo",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if len(args) < 1 || args[0] == "" {
			return errors.New("[project_name] empty")
		}
		sysPath, err := filepath.Abs(args[0])
		if err != nil {
			return
		}

		mkdirAll(sysPath)

		projectName := args[0]
		pdata := map[string]string{
			"PackagePath": projectName,
			"PackageName": projectName,
		}
		m := project.FileContent()
		for filepath, content := range m {
			var pf *os.File
			pf, err = os.Create(sysPath + filepath)
			if err != nil {
				return err
			}
			tmpl, err := template.New(projectName).Parse(content)
			if err != nil {
				return errors.New("parse error: " + err.Error())
			}
			if err = tmpl.Execute(pf, pdata); err != nil {
				return err
			}
		}
		exec.Command("gofmt", "-w", sysPath).Output()
		return
	},
}

func init() {
	AddCommand(NewProjectCmd)
}

func mkdirAll(projectPath string) {
	projectPath += "/app"
	os.MkdirAll(projectPath+"/server", os.ModePerm)
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
