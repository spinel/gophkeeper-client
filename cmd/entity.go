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
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spinel/gophkeeper-client/pkg/api"
	"github.com/spinel/gophkeeper-client/pkg/models"
	"github.com/spinel/gophkeeper-client/pkg/services"
)

// entityCmd represents the entity command
var entityCmd = &cobra.Command{
	Use:   "entity create | get",
	Short: "This command will auth the user",
	Long:  `This command will call gateway to auth a user.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("no operation specified")
			return
		}
		entity(args[0])
	},
}

func init() {
	rootCmd.AddCommand(entityCmd)
}

func entityCreateAccount(e models.Entity, authUser *models.User) {
	entityAccountUsernamePromptContent := promptContent{
		"Enter entity account username.",
		"Entity account username",
	}
	entityAccountUsername := promptGetInput(entityAccountUsernamePromptContent)

	entityAccountPasswordPromptContent := promptContent{
		"Enter entity account password.",
		"Entity account password",
	}
	entityAccountPassword := promptGetInput(entityAccountPasswordPromptContent)

	e.EntityPod.Account.Login = entityAccountUsername
	e.EntityPod.Account.Password = entityAccountPassword

	apiResp := &api.APIResponse{}
	client := api.NewClient(services.Mgr.Config.GatewayURL)
	if err := client.Post("entity/create", e, apiResp, authUser.Token); err != nil {
		fmt.Printf("Entity create failed: %v\n", err)
		return
	}
	if apiResp.Error != "" {
		fmt.Printf("Entity create failed: %v\n", apiResp.Error)
		return
	}
}

func entityCreate() {
	authUser := getAuthorisedUser()
	if authUser == nil {
		fmt.Println("You must be authorised ")
		return
	}

	entityTypePromptContent := promptContent{
		"Enter entity type.",
		"Entity type",
	}
	entityType := promptGetInput(entityTypePromptContent)

	entityIdentifierPromptContent := promptContent{
		"Enter entity identifier.",
		"Entity identifier",
	}
	entityIdentifier := promptGetInput(entityIdentifierPromptContent)

	entity := models.Entity{
		Identifier: entityIdentifier,
		TypeID:     1,
	}

	switch entityType {
	case models.EntityTypeAccount:
		entityCreateAccount(entity, authUser)
	default:
		fmt.Printf("Wrong entity type %s\n", entityType)
		return
	}
}

func entityGet() {

}

func entity(operation string) {
	switch operation {
	case "create":
		entityCreate()
	case "get":
		entityGet()
	default:
		fmt.Printf("Wrong entity operation %s\n", operation)
		return
	}
}
