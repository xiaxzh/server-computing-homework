// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"bufio"
	"fmt"
	"os"
	service "github.com/freakkid/service-agenda/cli/service"
	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update password",
	Long: `Use this command to update password`,
	Run: func(cmd *cobra.Command, args []string) {
		var old, password, cpassword string
		fmt.Print("Plase enter old password: ")
		reader := bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()
		old = string(data)

		fmt.Print("Plase enter new password: ")
		reader = bufio.NewReader(os.Stdin)
		data, _, _ = reader.ReadLine()
		password = string(data)

		fmt.Print("Plase enter new password again: ")
		reader = bufio.NewReader(os.Stdin)
		data, _, _ = reader.ReadLine()
		cpassword = string(data)
		ok, err := service.UpdatePassword(old, password, cpassword)
		if ok == false {
			fmt.Fprintln(os.Stderr, err)
		} else {
			fmt.Fprintln(os.Stdout, "Update password success!")
		}
		os.Exit(0)
	},
}

func init() {
	userCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logoutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
