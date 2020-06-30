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
	"fmt"
	"github.com/reud/wing/common"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var (
	toggle bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wing",
	Short: "定期コンテストのクリエイトを行えます。",
	Long:  `定期コンテストのクリエイトを行います。第一引数はコンテスト名のlowercaseで第二引数は開催番号です。`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// not root
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().IntVarP(&common.Number, "number", "n", 6, "問題数を設定します。")

	// フラグの値を変数にセットする場合
	// 第1引数: 変数のポインタ
	// 第2引数: フラグ名
	// 第3引数: デフォルト値
	// 第4引数: 説明
	rootCmd.PersistentFlags().StringVar(&common.TemplateCppFileName, "cpp", "cpp.template", "cppテンプレートファイルの名前を設定します。")
	rootCmd.PersistentFlags().StringVar(&common.TemplateShellFileName, "shell", "", "shellテンプレートファイルの名前を設定します。 (default virtual.sh.template (if virtual command)or debugger.sh.template)")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".wing" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".wing")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
