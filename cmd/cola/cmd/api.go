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
	"github.com/iancoleman/strcase"
	"github.com/spf13/viper"
	"github.com/zedisdog/cola/cmd/cola/stubs"
	"github.com/zedisdog/cola/pather"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// apiCmd represents the controller command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "generate api",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			color.Red("required a name for api")
			os.Exit(1)
		}
		packageName, _ := cmd.Flags().GetString("packageName")
		path, _ := cmd.Flags().GetString("path")

		p := pather.NewProjectPath()
		fileName := fmt.Sprintf("%s/%s.go", path, strcase.ToSnake(args[0]))

		err := os.MkdirAll(p.Dir(fileName), 0777)
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}

		f, err := os.OpenFile(p.Gen(fileName), os.O_CREATE|os.O_TRUNC|os.O_WRONLY|os.O_EXCL, 0777)
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}
		defer f.Close()
		replacer := strings.NewReplacer(
			"{{name}}", strcase.ToLowerCamel(args[0]),
			"{{moduleName}}", viper.GetString("moduleName"),
			"{{shortName}}", string([]rune(strcase.ToLowerCamel(args[0]))[0]),
			"{{varName}}", strcase.ToCamel(args[0]),
			"{{packageName}}", packageName,
		)
		f.WriteString(replacer.Replace(stubs.ControllerTemp))
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// controllerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// controllerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	apiCmd.Flags().StringP("path", "p", "internal/app/api", "Specify directory path to create in")
	apiCmd.Flags().StringP("packageName", "P", "api", "Specify package name")
}