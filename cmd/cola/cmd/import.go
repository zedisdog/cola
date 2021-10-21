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
	"bytes"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import [path to import]",
	Short: "import a project as a template",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		skipDir := []string{".idea", "data", ".git", "storage", "tmp"}
		dest := fmt.Sprintf("stubs%stemplate", string(os.PathSeparator))
		src := strings.TrimRight(args[0], "\\/")

		info, err := os.Stat(src)
		if err != nil {
			panic(err)
		}
		if !info.IsDir() {
			panic(errors.New("source project dir is needed"))
		}

		os.RemoveAll(dest)
		os.Mkdir(dest, 0755)

		err = filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) (e error) {
			if d.IsDir() || err != nil {
				for _, skip := range skipDir {
					if d.Name() == skip {
						return filepath.SkipDir
					}
				}
				return
			}
			destPath, content, e := convertSourceToTemplate(dest, src, path)
			if e != nil {
				return
			}
			os.MkdirAll(filepath.Dir(destPath), 0755)
			os.WriteFile(destPath, content, 0744)
			return
		})
		if err != nil {
			panic(err)
		}
	},
}

func convertSourceToTemplate(destDir string, srcDir string, srcPath string) (destPath string, content []byte, err error) {
	destPath = makeImportPath(destDir, srcDir, srcPath)
	content, err = os.ReadFile(srcPath)
	if err != nil {
		return
	}
	ext := filepath.Ext(srcPath)
	if ext == ".go" || ext == ".mod" {
		content = bytes.ReplaceAll(content, []byte("template"), []byte("{{.}}"))
	}
	return
}

func makeImportPath(destDir string, srcDir string, srcPath string) (path string) {
	path = strings.Replace(srcPath, srcDir, destDir, 1)
	if filepath.Ext(path) == ".go" || filepath.Ext(path) == ".mod" { // convert source code file to stub file
		path = fmt.Sprintf("%s.stub", path)
	}
	return
}

func init() {
	devCmd.AddCommand(importCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// importCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// importCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
