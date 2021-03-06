/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/zedisdog/cola/cmd/cola/stubs"
	"github.com/zedisdog/cola/pather"
	"os"
	"os/exec"
	"strings"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [flags] [PROJECT_NAME]",
	Short: "init a new project.",
	Long:  `init a new project.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			color.Red("Error: required 1 params for project name.")
			return
		}
		path, _ := cmd.Flags().GetString("path")
		useCurrentPath, _ := cmd.Flags().GetBool("currentPath")
		if !useCurrentPath {
			path = fmt.Sprintf("%s/%s", strings.TrimRight(path, "/"), args[0])
			err := os.Mkdir(path, 0777)
			if err != nil {
				color.Red("Error: %s", err.Error())
				return
			}
		}
		installGoMigrate(cmd)
		initModule(path, args[0], cmd)
		err := renderTemp(path, args[0])
		if err != nil {
			color.Red("Error: %s", err.Error())
			panic(err)
		}
		tidy(path, cmd)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	newCmd.Flags().StringP("path", "p", "./", "project path.")
	newCmd.Flags().BoolP("currentPath", "c", false, "if use current path as project root path.")
}

func installGoMigrate(cmd *cobra.Command) {
	c := exec.Command(
		"go",
		"install",
		"github.com/golang-migrate/migrate/v4/cmd/migrate@latest",
	)
	c.Stdout = cmd.OutOrStdout()
	c.Stderr = cmd.OutOrStderr()
	err := c.Run()
	if err != nil {
		color.Red("Error: %s", err.Error())
		panic(err)
	}
	color.Green("go-migrate install successful.")
}

func initModule(path string, name string, cmd *cobra.Command) {
	c := exec.Command("go", "mod", "init", name)
	c.Dir = path
	c.Stdout = cmd.OutOrStdout()
	c.Stderr = cmd.OutOrStderr()
	err := c.Run()
	if err != nil {
		color.Red("Error: %s", err.Error())
		panic(err)
	}
	color.Green("go module init successful.")
}

func tidy(path string, cmd *cobra.Command) {
	c := exec.Command("go", "mod", "tidy")
	c.Dir = path
	c.Stderr = cmd.OutOrStdout()
	c.Stdout = cmd.OutOrStdout()
	err := c.Run()
	if err != nil {
		color.Red("Error: %s", err.Error())
		panic(err)
	}
	color.Green("go mod tidy successful.")
}

func renderTemp(path string, moduleName string) (err error) {
	p := pather.New(strings.TrimLeft(path, "/"))
	err = createDirectory(p)
	if err != nil {
		return
	}

	err = renderMigrate(p)
	if err != nil {
		return
	}

	err = renderGitIgnore(p)
	if err != nil {
		return
	}

	err = renderLog(p)
	if err != nil {
		return
	}

	err = renderMain(p, moduleName)
	if err != nil {
		return
	}

	err = renderAirConf(p)
	if err != nil {
		return
	}

	err = renderDB(p)
	if err != nil {
		return
	}

	err = renderConfig(p)
	if err != nil {
		return
	}

	err = renderConfigGo(p, moduleName)
	if err != nil {
		return
	}

	err = renderRoutes(p)
	if err != nil {
		return
	}

	err = renderTestController(p, moduleName)
	if err != nil {
		return
	}

	err = renderDockerCompose(p)
	if err != nil {
		return
	}

	err = renderStorage(p)
	if err != nil {
		return
	}

	return nil
}

func renderGitIgnore(path *pather.Pather) error {
	return renderFile(
		path.Gen(".gitignore"),
		stubs.GitIgnore,
	)
}

func renderMain(path *pather.Pather, moduleName string) error {
	return renderFile(
		path.Gen("cmd/app/main.go"),
		stubs.MainTemp,
		"{{moduleName}}", moduleName,
	)
}

func renderAirConf(path *pather.Pather) error {
	return renderFile(
		path.Gen(".air.conf"),
		stubs.AirConf,
	)
}

func renderLog(path *pather.Pather) error {
	return renderFile(
		path.Gen("internal/log/log.go"),
		stubs.LogTemp,
	)
}

func renderFile(path string, tmp string, oldnew ...string) error {
	f, err := os.OpenFile(
		path,
		os.O_CREATE|os.O_TRUNC|os.O_WRONLY,
		0777,
	)
	if err != nil {
		return err
	}
	defer f.Close()

	replacer := strings.NewReplacer(oldnew...)
	_, err = f.Write([]byte(replacer.Replace(tmp)))
	if err != nil {
		return err
	}
	return nil
}

func renderDB(path *pather.Pather) error {
	return renderFile(
		path.Gen("internal/database/db.go"),
		stubs.DbTemp,
	)
}

func renderConfig(path *pather.Pather) error {
	return renderFile(
		path.Gen("config.yaml"),
		stubs.ConfigTemp,
	)
}

func renderRoutes(path *pather.Pather) error {
	return renderFile(
		path.Gen("internal/controllers/routes.go"),
		stubs.RoutesTemp,
	)
}

func renderMigrate(path *pather.Pather) error {
	err := renderFile(
		path.Gen("internal/database/migrations/1615120646098759700_create_accounts_table.up.sql"),
		stubs.MigrateUp,
	)
	if err != nil {
		return err
	}
	return renderFile(
		path.Gen("internal/database/migrations/1615120646098759700_create_accounts_table.down.sql"),
		stubs.MigrateDown,
	)
}

func renderTestController(path *pather.Pather, moduleName string) error {
	return renderFile(
		path.Gen("internal/controllers/test_controller.go"),
		stubs.ControllerTemp,
		"{{moduleName}}", moduleName,
		"{{name}}", "test",
		"{{shortName}}", "t",
		"{{varName}}", "Test",
	)
}

func renderDockerCompose(path *pather.Pather) error {
	return renderFile(
		path.Gen("docker-compose.yml"),
		stubs.DockerComposeTemp,
	)
}

func renderConfigGo(path *pather.Pather, moduleName string) error {
	return renderFile(
		path.Gen("config.go"),
		stubs.ConfigGoTemp,
		"{{moduleName}}", moduleName,
	)
}

func renderStorage(path *pather.Pather) error {
	return renderFile(
		path.Gen("internal/storage/storage.go"),
		stubs.StorageTemp,
	)
}

func createDirectory(path *pather.Pather) (err error) {
	err = os.MkdirAll(path.Gen("cmd/app"), 0777)
	if err != nil {
		return
	}
	err = os.MkdirAll(path.Gen("internal/log"), 0777)
	if err != nil {
		return
	}
	err = os.MkdirAll(path.Gen("internal/controllers"), 0777)
	if err != nil {
		return
	}
	err = os.MkdirAll(path.Gen("internal/database/migrations"), 0777)
	if err != nil {
		return
	}
	err = os.MkdirAll(path.Gen("internal/models"), 0777)
	if err != nil {
		return
	}
	err = os.MkdirAll(path.Gen("internal/storage"), 0777)
	if err != nil {
		return
	}
	return nil
}
