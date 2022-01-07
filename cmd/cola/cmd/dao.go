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

	"github.com/spf13/cobra"
)

// daoCmd represents the dao command
var daoCmd = &cobra.Command{
	Use:   "dao",
	Short: "generate dao",
	Long:  `generate dao.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("dao name is required")
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
		daoPath := fmt.Sprintf("%s%cdao", path, os.PathSeparator)

		err = os.MkdirAll(daoPath, 0766)
		if err != nil {
			cmd.OutOrStderr().Write([]byte(fmt.Sprintf("err: %s\n", err.Error())))
			return
		}

		err = generate.File(tpl.FileTempOptions{
			PkgName: "dao",
			Imports: []tpl.ImportOptions{
				{
					Alias:  "Gorm",
					Import: "gorm.io/gorm",
				},
				{
					Import: "github.com/zedisdog/cola/database/drivers/gorm",
				},
			},
			Interfaces: []tpl.InterfaceTempOptions{
				{
					Name: name + "Dao",
					Funcs: []tpl.InterfaceMethod{
						{
							Name: "Transaction",
							Params: map[string]string{
								"f": "func(tx *Gorm.DB) error",
							},
							Returns: map[string]string{
								"err": "error",
							},
						},
						{
							Name: "WithTx",
							Params: map[string]string{
								"tx": "*Gorm.DB",
							},
							Returns: []string{
								name + "Dao",
							},
						},
					},
				},
			},
			Funcs: []tpl.FuncTempOptions{
				{
					Name:    "New" + name + "Dao",
					Returns: []string{name + "Dao"},
					Content: `return &` + name + `{
		db: gorm.NewDBHelper(),
	}`,
				},
			},
			Structs: []tpl.StructTempOptions{
				{
					Name: name,
					Fields: []tpl.Field{
						{
							Name: "db",
							Type: "*gorm.DBHelper",
						},
					},
					Methods: []tpl.FuncTempOptions{
						{
							Name: "Transaction",
							Params: map[string]string{
								"f": "func(tx *Gorm.DB) error",
							},
							Returns: map[string]string{
								"err": "error",
							},
							Content: fmt.Sprintf("return %s.db.Transaction(f)", strings.ToLower(string(name[0]))),
						},
						{
							Name: "WithTx",
							Params: map[string]string{
								"tx": "*Gorm.DB",
							},
							Returns: []string{
								name + "Dao",
							},
							Content: fmt.Sprintf(`%s.db.WithTx(tx)
	return &%s`, strings.ToLower(string(name[0])), strings.ToLower(string(name[0]))),
						},
					},
				},
			},
		}, fmt.Sprintf("%s%c%s.go", daoPath, os.PathSeparator, strcase.ToSnake(name)))
		if err != nil {
			cmd.OutOrStderr().Write([]byte(fmt.Sprintf("err: %s\n", err.Error())))
			return
		}

		cmd.OutOrStdout().Write([]byte("dao has be generated.\n"))
	},
}

func init() {
	makeCmd.AddCommand(daoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// daoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// daoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
