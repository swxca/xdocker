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
	"log"
	"os"
	"xdocker/process"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		i, err := cmd.Flags().GetBool("it")
		if err!=nil{
			fmt.Printf("command run Err : %v" ,err)
			return
		}
		log.Printf("run command args is %v",args)
		run(i,args[1])
	},
}

func run(it bool, command string) {
	createProcessCmd := process.CreateProcess(it, command)
	err := createProcessCmd.Start()
	if err!=nil{
		panic(err)
	}
	fmt.Println("process is running")
	err = createProcessCmd.Wait()
	if err!=nil{
		panic(err)
	}
	os.Exit(-1)
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().BoolP("it","i",false,"交互命令窗口")
	//runCmd.MarkFlagRequired("it")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
