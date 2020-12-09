// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// testcmdCmd represents the testcmd command
var testcmdCmd = &cobra.Command{
	Use:   "testcmd",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("testcmd called-->PersistentPreRun")
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("testcmd called-->PreRun")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("testcmd called-->Run")
		if len(args) > 1 {
			fmt.Println("arg1-->" + args[0])
			fmt.Println("arg2-->" + args[1])
		}
		fmt.Printf("Stu Name is %s, Stu age is %d\n", stu_name, stu_age)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("testcmd called-->PostRun")
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("testcmd called-->PersistentPostRunE")
	},
}

var stu_name string
var stu_age int

func init() {
	rootCmd.AddCommand(testcmdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testcmdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testcmdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	testcmdCmd.Flags().StringVarP(&stu_name, "name", "n", "def_str", "an stu nameeeee")
	testcmdCmd.Flags().IntVarP(&stu_age, "age", "a", 5201314, "an stu ageeee")
}
