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
	tools "github.com/freakkid/service-agenda/cli/tools"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "User login",
	Long:  `Use this command to sign in to the system.`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		// wait for password
		var password string
		fmt.Printf("Please enter the password: ")
		reader := bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()
		password = string(data)

		// validate
		ok, message := tools.ValidateUsername(username)
		tools.DealMessage(ok, message)
		ok, message = tools.ValidatePass(password)
		tools.DealMessage(ok, message)
		// send login request
		ok, message = service.GetUserKey(username, password)
		tools.DealMessage(ok, message)
		fmt.Println(message)
	},
}

func init() {
	userCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	loginCmd.Flags().StringP("username", "u", "", "Login username")
}
