/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
	"github.com/zedisdog/cola/cmd/cola/internal/generate"
	"github.com/zedisdog/cola/cmd/cola/tpl"
	"os"
	"strings"
)

// moduleCmd represents the module command
var moduleCmd = &cobra.Command{
	Use:   "module [flags] <moduleName>",
	Short: "generate a module",
	Long:  `generate a module.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("module name is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("path")
		if err != nil || path == "" {
			path = "."
		} else {
			path = strings.TrimRight(path, string(os.PathSeparator))
		}
		name := strcase.ToSnake(args[0])
		modulePath := fmt.Sprintf("%s%c%s", path, os.PathSeparator, name)

		err = os.MkdirAll(modulePath, 0766)
		if err != nil {
			cmd.OutOrStderr().Write([]byte(fmt.Sprintf("err: %s\n", err.Error())))
			return
		}

		err = generate.File(tpl.FileTempOptions{
			PkgName: name,
			Interfaces: []tpl.InterfaceTempOptions{
				{
					Name: "Service",
				},
			},
		}, fmt.Sprintf("%s%ccontract.go", modulePath, os.PathSeparator))
		if err != nil {
			cmd.OutOrStderr().Write([]byte(fmt.Sprintf("err: %s\n", err.Error())))
			return
		}

		cmd.OutOrStdout().Write([]byte("module has be generated.\n"))
	},
}

func init() {
	makeCmd.AddCommand(moduleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// moduleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// moduleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
