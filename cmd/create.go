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
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [ContestName] [ContestNumber]",
	Short: "常設コンテスト(ABC,ARC,AGC)向けコマンド",
	Long: `wing create <コンテスト名> <コンテスト番号> でコンテストを作成することができます。作成場所は <コンテスト名>/<コンテスト番号>です。
このコマンドにはサンプル取得用のURLが含まれます。
	`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(common.TemplateShellFileName) == 0 {
			common.TemplateShellFileName = "debugger.sh.template"
		}
		if len(args) != 2 {
			return errors.New("引数の形式が不正です。　sample: wing create abc 170\n")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		contest := args[0]
		number := args[1]
		if err := os.MkdirAll(fmt.Sprintf("%+v/%+v", contest, number), 0777); err != nil {
			fmt.Println(err)
		}
		b := common.GetBinaryFromFile(common.TemplateCppFileName)

		common.WriteContestCMakeList(contest, number, 6)
		common.WriteCodeFiles(contest, number, 6, b)

		b, err := ioutil.ReadFile(common.TemplateShellFileName)
		if err != nil {
			panic(err)
		}
		cn := fmt.Sprintf("%+v%+v", contest, number)
		b = common.BinaryReplace(b, "###CONTEST_NAME###", cn)
		path := fmt.Sprintf("%+v/%+v/debugger.sh", contest, number)
		if err := common.WriteFile(path, b); err != nil {
			panic(err)
		}
	},
}

func init() {
	createCmd.Flags().StringVar(&common.TemplateShellFileName, "shell", "debugger.sh.template", "shellテンプレートファイルの名前を設定します。")
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
