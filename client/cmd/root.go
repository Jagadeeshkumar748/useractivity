package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/Jagadeeshkumar748/useractivity/user"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var userclient user.UserActivityClient

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "a gRPC client to communicate with the UserActivity server",
	Long: `a gRPC client to communicate with the UserActivity server.
	You can use this client to create and read user activities.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		fmt.Println("+++++++++++++++++++++++++++++++++++++")
		os.Exit(1)
	}
}

// func initConfig() {
// }

func init() {
	cobra.OnInitialize()
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	// conn, err := grpc.Dial(os.Getenv("USER_PORT"), grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	userclient = user.NewUserActivityClient(conn)

}
