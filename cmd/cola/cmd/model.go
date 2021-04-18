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
	"github.com/zedisdog/cola/cmd/cola/stubs"
	"github.com/zedisdog/cola/pather"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// modelCmd represents the model command
var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			color.Red("required a name for model")
			os.Exit(1)
		}
		path := pather.NewProjectPath()
		// 创建目录
		err := os.MkdirAll(path.Gen("internal/models"), 0777)
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}
		modelPath := fmt.Sprintf("%s/%s.go", path.Gen("internal/models"), strcase.ToSnake(args[0]))
		f, err := os.OpenFile(modelPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY|os.O_EXCL, 0777)
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}
		defer f.Close()
		replacer := strings.NewReplacer(
			"{{name}}", strcase.ToCamel(args[0]),
		)
		f.WriteString(replacer.Replace(stubs.ModelTemp))
	},
}

func init() {
	rootCmd.AddCommand(modelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// modelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// modelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
