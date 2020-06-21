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
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

func writeCMakeList(contest string,number string,cnt int) {
	//os.O_RDWRを渡しているので、同時に読み込みも可能
	file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		//エラー処理
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Fprintln(file, fmt.Sprintf("# Contest %+v %+v",contest,number)) //書き込み
	for i:=0 ; i < cnt ; i++ {
		fmt.Fprintln(file, fmt.Sprintf("add_executable(%+v-%+v_%+v %+v/%+v/%+v.cpp)",contest,number,string(rune('a'+i)),contest,number,string(rune('a'+i)))) //書き込み
	}
}

func writeCodeFiles(contest string,number string,cnt int,template []byte) {
	for i:=0 ; i < cnt ; i++ {
		fn := fmt.Sprintf("%+v/%+v/%+v.cpp",contest,number,string(rune('a'+i)))
		if err := ioutil.WriteFile(fn,template,0666) ; err != nil {
			fmt.Printf("%+v",err)
		}
	}
}


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wing",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		contest := args[0]
		number := args[1]
		if err := os.MkdirAll(fmt.Sprintf("%+v/%+v",contest,number), 0777); err != nil {
			fmt.Println(err)
		}
		b, err := ioutil.ReadFile("cpp.template")
		if err != nil {
			panic(err)
		}
		writeCMakeList(contest,number,6)
		writeCodeFiles(contest,number,6,b)
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

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.wing.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
