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
	"github.com/zedisdog/cola/cmd/cola_bak/migrate"
	"os"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create migration",
	Long:  `create migration`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			color.Red("command required a file name for file")
			os.Exit(1)
		}
		ext, _ := cmd.Flags().GetString("ext")
		format, _ := cmd.Flags().GetString("format")
		err := migrate.Create(args[0], ext, format)
		if err != nil {
			color.Red(err.Error())
		}
	},
}

func init() {
	migrateCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createCmd.Flags().StringP("ext", "e", "sql", "ext for sql file")
	createCmd.Flags().StringP("format", "f", "unixNano", "prefix format")
}
