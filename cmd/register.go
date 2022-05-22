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
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spinel/gophkeeper-client/pkg/api"
	"github.com/spinel/gophkeeper-client/pkg/models"
	"github.com/spinel/gophkeeper-client/pkg/services"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "register",
	Short: "This command will get the desired Gopher",
	Long:  `This get command will call GitHub respository in order to return the desired Gopher.`,
	Run: func(cmd *cobra.Command, args []string) {
		registerUser()
	},
}

type promptContent struct {
	errorMsg string
	label    string
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func promptGetInput(pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}

func registerUser() {
	emailPromptContent := promptContent{
		"Please provide an email.",
		"New user email",
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
	if err := client.Post("auth/register", userForm, apiResp, ""); err != nil {
		fmt.Printf("Register failed: %v\n", err)
		return
	}
	if apiResp.Error != "" {
		fmt.Printf("Register failed: %v\n", apiResp.Error)
		return
	}

	createdUser, err := services.Mgr.User.Create(context.Background(), userForm)
	if err != nil {
		fmt.Printf("Register failed: %v\n", err)
		return
	}

	fmt.Printf("New user %s was successfully created!\n", createdUser.Email)
}
