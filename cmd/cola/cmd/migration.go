/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/zedisdog/cola/cmd/cola/internal/generate"
	"github.com/zedisdog/cola/cmd/cola/tpl"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// migrationCmd represents the migration command
var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "generate a migration",
	Long:  `generate a migration.`,
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
		migrationPath := fmt.Sprintf("%s%c%s", path, os.PathSeparator, "migations")

		err = os.MkdirAll(migrationPath, 0766)
		if err != nil {
			cmd.OutOrStderr().Write([]byte(fmt.Sprintf("err: %s\n", err.Error())))
			return
		}

		up, down, err := tpl.GenMigration(name)
		if err != nil {
			cmd.OutOrStderr().Write([]byte(fmt.Sprintf("err: %s\n", err.Error())))
			return
		}
		t := time.Now().Format("20060102150405")
		err = generate.FileBytes(up, fmt.Sprintf(
			"%s%c%s_create_%s_table.%s.sql",
			migrationPath,
			os.PathSeparator,
			t,
			name,
			"up",
		))
		if err != nil {
			cmd.OutOrStderr().Write([]byte(fmt.Sprintf("err: %s\n", err.Error())))
			return
		}

		err = generate.FileBytes(down, fmt.Sprintf(
			"%s%c%s_create_%s_table.%s.sql",
			migrationPath,
			os.PathSeparator,
			t,
			name,
			"down",
		))
		if err != nil {
			cmd.OutOrStderr().Write([]byte(fmt.Sprintf("err: %s\n", err.Error())))
			return
		}

		cmd.OutOrStdout().Write([]byte("migrations has be generated.\n"))
	},
}

func init() {
	makeCmd.AddCommand(migrationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
