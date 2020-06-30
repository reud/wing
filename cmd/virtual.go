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
	"log"
	"os"

	"github.com/spf13/cobra"
)

var ()

func writeVirtualContestCMakeList(contest string, cnt int) {
	//os.O_RDWRを渡しているので、同時に読み込みも可能
	file, err := os.OpenFile("CMakeLists.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		//エラー処理
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Fprintln(file, fmt.Sprintf("# Virtual Contest %+v", contest)) //書き込み
	for i := 0; i < cnt; i++ {
		fmt.Fprintln(file, fmt.Sprintf("add_executable(%+v-%+v_%+v %+v/%+v/%+v.cpp)", "virtual", contest, i+1, "virtual", contest, i+1)) //書き込み
	}
}

func writeVirtualCodeFiles(contest string) {
	b := common.GetBinaryFromFile(common.TemplateCppFileName)
	for i := 0; i < common.Number; i++ {
		fn := fmt.Sprintf("%+v/%+v/%+v.cpp", "virtual", contest, i+1)
		if err := ioutil.WriteFile(fn, b, 0666); err != nil {
			fmt.Printf("%+v", err)
		}
	}
}

func setUpVirtualContestFiles(contestName string) {

	// writeCMakeList
	writeVirtualContestCMakeList(contestName, common.Number)
	// writeCodefiles
	writeVirtualCodeFiles(contestName)
	writeShellScript(contestName)
}

// for case で作れる
func writeShellScript(contestName string) {
	// 入れ込み
	var ret string = ""
	for i := 0; i < common.Number; i++ {
		ret = ret + fmt.Sprintf(`  "%+v" ) URL=https://atcoder.jp/contests/abcxxx/tasks/abcxxx_a ;;`+"\n", i+1)
	}
	b := common.GetBinaryFromFile(common.TemplateShellFileName)
	b = common.BinaryReplace(b, "###INJECT_CASES###", ret)
	err := common.WriteFile(fmt.Sprintf("virtual/%+v/debugger.sh", contestName), b)
	if err != nil {
		panic(err)
	}
}

// virtualCmd represents the virtual command
var virtualCmd = &cobra.Command{
	Use:   "virtual [ContestName]",
	Short: "バチャを行う時に呼ぶコマンド",
	Long:  `バチャを行う時に呼ぶコマンドとなります。`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(common.TemplateShellFileName) == 0 {
			common.TemplateShellFileName = "virtual.sh.template"
		}
		if len(args) != 1 {
			return errors.New("引数の数が不正です。")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		contestName := args[0]
		err := common.CreateDirectory("virtual", contestName)
		if err != nil {
			panic(err)
		}
		// pick template from template file
		// write CMakeLists
		setUpVirtualContestFiles(contestName)
	},
}

func init() {
	rootCmd.AddCommand(virtualCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// virtualCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// virtualCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
