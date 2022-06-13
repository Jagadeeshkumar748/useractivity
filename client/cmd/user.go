package cmd

import (
	"context"
	"fmt"

	// "os/user"

	"github.com/Jagadeeshkumar748/useractivity/user"
	"github.com/spf13/cobra"
)

// var userclient user.UserActivityClient

var createUserCmd = &cobra.Command{

	Use:   "createuser",
	Short: "Create a new user",
	Long: `Create a new user on the server through gRPC
	       name, email and phone are required`,

	RunE: func(cmd *cobra.Command, args []string) error {

		name, err := cmd.Flags().GetString("name")
		email, err := cmd.Flags().GetString("email")
		phone, err := cmd.Flags().GetString("phone")
		if err != nil {
			return err
		}

		u := &user.User{

			Name:  name,
			Email: email,
			Phone: phone,
		}
		user, err := userclient.AddUser(
			context.TODO(),

			&user.UserRequest{
				User: u,
			},
		)
		if err != nil {
			return err
		}
		fmt.Printf("User created: %d\n", user.User)
		return nil
	},
}

func init() {
	createUserCmd.Flags().StringP("name", "n", "", "Add a name")
	createUserCmd.Flags().StringP("email", "e", "", "An email of the user")
	createUserCmd.Flags().StringP("phone", "p", "", "Add a phone")
	createUserCmd.MarkFlagRequired("name")
	createUserCmd.MarkFlagRequired("email")
	createUserCmd.MarkFlagRequired("phone")
	rootCmd.AddCommand(createUserCmd)
}
