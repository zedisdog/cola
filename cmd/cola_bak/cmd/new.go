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
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/zedisdog/cola/cmd/cola_bak/lib"

	"os/exec"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [PROJECT_NAME]",
	Short: "init a new project.",
	Long:  `init a new project.`,
	Run: func(cmd *cobra.Command, args []string) {
		moduleName := args[0]
		if len(args) < 1 {
			color.Red("Error: required 1 params for project name.")
			return
		}
		err := lib.Export(moduleName, moduleName)
		if err != nil {
			color.Red("Error: %s", err.Error())
			return
		}
		// install go-migrate cli tool
		//installGoMigrate(cmd)
		tidy(moduleName, cmd)
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
}

func installGoMigrate(cmd *cobra.Command) {
	if _, err := exec.LookPath("migrate"); err != nil {
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
	} else {
		color.Green("go-migrate already installed.")
	}
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
