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
	"github.com/zedisdog/cola/cola/stubs"
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
		initModule(path, args[0], cmd)
		err := renderTemp(path, args[0])
		if err != nil {
			color.Red("Error: %s", err.Error())
			return
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

func initModule(path string, name string, cmd *cobra.Command) {
	c := exec.Command("go", "mod", "init", name)
	c.Dir = path
	c.Stdout = cmd.OutOrStdout()
	c.Stderr = cmd.OutOrStderr()
	err := c.Run()
	if err != nil {
		color.Red("Error: %s", err.Error())
		return
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
	}
	color.Green("go mod tidy successful.")
}

func renderTemp(path string, moduleName string) error {
	err := createDirectory(path)
	if err != nil {
		return err
	}

	err = renderLog(path)
	if err != nil {
		return err
	}

	err = renderMain(path, moduleName)
	if err != nil {
		return err
	}

	return nil
}

func renderMain(path string, moduleName string) error {
	f, err := os.OpenFile(
		fmt.Sprintf("%s/%s", strings.TrimLeft(path, "/"), "app/main.go"),
		os.O_TRUNC|os.O_CREATE,
		0777,
	)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write([]byte(strings.ReplaceAll(stubs.MainTemp, "{{moduleName}}", moduleName)))
	if err != nil {
		return err
	}
	return nil
}

func renderLog(path string) error {
	f, err := os.OpenFile(
		fmt.Sprintf("%s/%s", strings.TrimLeft(path, "/"), "internal/log/log.go"),
		os.O_CREATE|os.O_TRUNC,
		0777,
	)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write([]byte(stubs.LogTemp))
	if err != nil {
		return err
	}
	return nil
}

func createDirectory(path string) error {
	err := os.Mkdir(fmt.Sprintf("%s/%s", strings.TrimLeft(path, "/"), "internal"), 0777)
	if err != nil {
		return err
	}
	err = os.Mkdir(fmt.Sprintf("%s/%s", strings.TrimLeft(path, "/"), "app"), 0777)
	if err != nil {
		return err
	}
	err = os.Mkdir(fmt.Sprintf("%s/%s", strings.TrimLeft(path, "/"), "internal/log"), 0777)
	if err != nil {
		return err
	}
	err = os.Mkdir(fmt.Sprintf("%s/%s", strings.TrimLeft(path, "/"), "internal/controllers"), 0777)
	if err != nil {
		return err
	}
	return nil
}
