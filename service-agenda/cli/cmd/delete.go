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
	"fmt"
	"os"
	"bufio"
	service "github.com/freakkid/service-agenda/cli/service"
	tools	"github.com/freakkid/service-agenda/cli/tools"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var udeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete user account",
	Long: `Use this command to delete your account, meetings included.`,
	Run: func(cmd *cobra.Command, args []string) {
		// hints to ensure and enter password to delete User
		var password string
		fmt.Print("Plase enter password: ")
		reader := bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()
		password = string(data)
		// validate
		ok, message := tools.ValidatePass(password)
		if !ok {
			fmt.Fprintf(os.Stderr, message)
			os.Exit(1)
		}
		// delete user and meetings it participate
		ok, err := service.DeleteUser(password)
		if !ok {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(0)
		} else {
			fmt.Fprintln(os.Stdout, "Success delete current user.")
		}
	},
}

func init() {
	userCmd.AddCommand(udeleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	udeleteCmd.Flags().StringP("username", "u", "", "Delete user")
}