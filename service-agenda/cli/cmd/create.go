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

// createCmd represents the create command
var ucreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create user account",
	Long:  `Use this command to create a new user account.`,
	Run: func(cmd *cobra.Command, args []string) {
		// get createUser information
		createUsername, _ := cmd.Flags().GetString("username")
		createEmail, _ := cmd.Flags().GetString("email")
		createPhone, _ := cmd.Flags().GetString("phone")
		// get the password
		var createPassword string
		var prePassword string
		times := 1
		reader := bufio.NewReader(os.Stdin)

		for {
			if times == 1 {
				fmt.Print("Please enter the password you want to create: ")
				data, _, _ := reader.ReadLine()
				createPassword = string(data)
			} else {
				fmt.Print("Please enter the password again: ")
				data, _, _ := reader.ReadLine()
				createPassword = string(data)
				if createPassword == prePassword {
					break
				} else {
					fmt.Println("The two passwords entered are not consistent. \nPlease restart setting password.")
				}
			}
			times *= -1
			prePassword = createPassword
		}
		// validate
		ok, message := tools.ValidateUsername(createUsername)
		tools.DealMessage(ok, message)
		ok, message = tools.ValidatePhone(createPhone)
		tools.DealMessage(ok, message)
		ok, message = tools.ValidateEmail(createEmail)
		tools.DealMessage(ok, message)
		ok, message = tools.ValidatePass(createPassword)
		tools.DealMessage(ok, message)
		// send create request
		ok, message = service.CreateUser(createUsername, createPassword, createPhone, createEmail)
		tools.DealMessage(ok, message)
		fmt.Println(message)

	},
}

func init() {
	userCmd.AddCommand(ucreateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// xiaxzh's part:
	ucreateCmd.Flags().StringP("username", "u", "", "Create Username")
	ucreateCmd.Flags().StringP("email", "e", "", "Create Email")
	ucreateCmd.Flags().StringP("phone", "p", "", "Create Phone")
}
