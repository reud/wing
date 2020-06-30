/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"errors"
	"fmt"
	"github.com/reud/wing/common"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

// otherCmd represents the other command
var otherCmd = &cobra.Command{
	Use:   "other [contestName]",
	Short: "企業コンとか作る時に使うやつ",
	Long:  `企業コンとか作る時に使うやつ`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(common.TemplateShellFileName) == 0 {
			common.TemplateShellFileName = "debugger.sh.template"
		}
		if len(args) != 1 {
			panic(errors.New("bad args"))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		contest := args[0]
		if err := os.MkdirAll(fmt.Sprintf("others/%+v", contest), 0777); err != nil {
			fmt.Println(err)
		}
		b, err := ioutil.ReadFile(common.TemplateCppFileName)
		if err != nil {
			panic(err)
		}
		common.WriteContestCMakeList("others", contest, 6)
		common.WriteCodeFiles("others", contest, 6, b)

		b, err = ioutil.ReadFile(common.TemplateShellFileName)
		if err != nil {
			panic(err)
		}
		b = common.BinaryReplace(b, "###CONTEST_NAME###", contest)
		path := fmt.Sprintf("%+v/%+v/debugger.sh", "others", contest)
		if err := common.WriteFile(path, b); err != nil {
			panic(err)
		}
	},
}

func init() {
	otherCmd.Flags().StringVar(&common.TemplateShellFileName, "shell", "", "shellテンプレートファイルの名前を設定します。")
	rootCmd.AddCommand(otherCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// otherCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// otherCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
