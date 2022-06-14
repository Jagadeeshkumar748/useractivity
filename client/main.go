package main

import (
	"context"
	"log"
	"time"

	// "fmt"
	"github.com/Jagadeeshkumar748/useractivity/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addrress = "localhost:50051"
	id       = ""
	// name     = "Jagadeesh"
	email = "jagadeeshkumar748.32@gmail.com"
	// phone    = "9963106464"

	// details = &user.User{Userid: id, Name: name, Email: email, Phone: phone}

	activitydetails = &user.Activity{Id: id, ActivityType: 2, Duration: 10, Email: email}
)

func main() {

	conn, err := grpc.Dial(addrress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := user.NewUserActivityClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Adding user block
	// r, err := c.AddUser(ctx, &user.UserRequest{User: details})
	// if err != nil {
	// 	log.Fatal("error in adding user", err)
	// }

	// log.Printf("Adding user: %s", r.GetUser())

	//querying user block
	// f, err := c.QueryUser(ctx, &user.QueryUserRequest{Userid: "62a5ffdc7f9bea11dfbd184a"})
	// if err != nil {
	// 	log.Fatal("error in finding user", err)
	// }
	// fmt.Println("User details :", f.GetUser())

	//adding activity block

	a, err := c.AddActivity(ctx, &user.ActivityRequest{Activity: activitydetails})
	if err != nil {
		log.Fatal("error in adding user", err)
	}

	log.Printf("Adding user: %s", a.GetActivity())

}
