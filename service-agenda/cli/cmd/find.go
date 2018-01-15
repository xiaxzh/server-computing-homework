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

	service "github.com/freakkid/service-agenda/cli/service"
	tools "github.com/freakkid/service-agenda/cli/tools"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "User find",
	Long:  `Use this command to find user by id.`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		// validate
		ok, message := tools.ValidateId(id)
		tools.DealMessage(ok, message)
		var retJSON service.SingleUserInfo
		ok, message, retJSON = service.FindUser(id)
		tools.DealMessage(ok, message)
		fmt.Printf("%-5s%-15s%-25s%-25s\n", "Id", "Username", "Phone number", "E-mail")
		fmt.Printf("%-5d%-15s%-25s%-25s\n", retJSON.ID, retJSON.UserName, retJSON.Phone, retJSON.Email)
	},
}

func init() {
	userCmd.AddCommand(findCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	findCmd.Flags().StringP("id", "i", "", "user id")
}
