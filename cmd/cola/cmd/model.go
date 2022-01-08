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

// modelCmd represents the model command
var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "generate a model",
	Long:  `generate a model.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("model name is required")
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
		name := strcase.ToCamel(args[0])
		modelPath := fmt.Sprintf("%s%c%s", path, os.PathSeparator, "models")

		err = os.MkdirAll(modelPath, 0766)
		if err != nil {
			cmd.OutOrStderr().Write([]byte(fmt.Sprintf("err: %s\n", err.Error())))
			return
		}

		err = generate.File(tpl.FileTempOptions{
			PkgName: "models",
			Imports: []tpl.ImportOptions{
				{
					Import: "github.com/zedisdog/cola/database/drivers/gorm",
				},
			},
			Structs: []tpl.StructTempOptions{
				{
					Name: name,
					Fields: []tpl.Field{
						{
							Name: "gorm.CommonField",
						},
					},
				},
			},
		}, fmt.Sprintf("%s%c%s.go", modelPath, os.PathSeparator, strcase.ToSnake(name)))
		if err != nil {
			cmd.OutOrStderr().Write([]byte(fmt.Sprintf("err: %s\n", err.Error())))
			return
		}

		cmd.OutOrStdout().Write([]byte("model has be generated.\n"))
	},
}

func init() {
	makeCmd.AddCommand(modelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// modelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// modelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
