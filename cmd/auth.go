/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spinel/gophkeeper-client/pkg/api"
	"github.com/spinel/gophkeeper-client/pkg/models"
	"github.com/spinel/gophkeeper-client/pkg/services"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "This command will auth the user",
	Long:  `This command will call gateway to auth a user.`,
	Run: func(cmd *cobra.Command, args []string) {
		authUser()
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}

func getAuthorisedUser() *models.User {
	user, err := services.Mgr.User.GetAuthorised(context.Background())
	if err != nil {
		fmt.Printf("Get auth user failed: %v\n", err)
		return nil
	}

	return user
}

func authUser() {
	emailPromptContent := promptContent{
		"Please provide an email.",
		"User email",
	}
	email := promptGetInput(emailPromptContent)

	pwdPromptContent := promptContent{
		"Please provide a password.",
		"User password",
	}
	pwd := promptGetInput(pwdPromptContent)

	userForm := models.UserForm{
		Email:    email,
		Password: pwd,
	}

	apiResp := &api.APIResponse{}
	client := api.NewClient(services.Mgr.Config.GatewayURL)
	if err := client.Post("auth/login", userForm, apiResp, ""); err != nil {
		fmt.Printf("Auth failed: %v\n", err)
		return
	}
	if apiResp.Error != "" {
		fmt.Printf("Auth failed: %v\n", apiResp.Error)
		return
	}
	userForm.Token = apiResp.Token

	createdUser, err := services.Mgr.User.SetToken(context.Background(), userForm)
	if err != nil {
		fmt.Printf("Auth failed: %v\n", err)
		return
	}

	fmt.Printf("User %s was successfully authorised!\n", createdUser.Email)
}
