package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/Jagadeeshkumar748/useractivity/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var db *mongo.Client
var usersdb *mongo.Collection
var activitydb *mongo.Collection
var mongoCtx context.Context

type UserActivityServer struct {
	user.UnimplementedUserActivityServer
}

type UserItem struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Email string             `bson:"email"`
	Phone string             `bson:"phone"`
}

// type ActivityItem struct {
// 	ID                    primitive.ObjectID `bson:"_id,omitempty"`
// 	         `bson:"activitytype"`
// 	Duration              int64              `bson:"duration"`
// 	Email                 string             `bson:"emaill"`
// 	Timestamp             int64              `bson:"timestamp"`
// }

type Sleep struct {
	*user.Activity
}

type Play struct {
	*user.Activity
}

type Eat struct {
	*user.Activity
}
type Read struct {
	*user.Activity
}

// type ActivityMethods interface {
// 	isDone() bool
// 	isValid() error
// }

// func Validate(a *user.Activity) error {
// 	if a.User.Userid == "0" {
// 		return fmt.Errorf("invalid user id %S", a.User.Userid)
// 	}
// 	switch a.ActivityType {
// 	case user.Activity_Sleep:
// 		return Sleep{a}.isValid()
// 	case user.Activity_Play:
// 		return Play{a}.isValid()
// 	case user.Activity_Eat:
// 		return Eat{a}.isValid()
// 	case user.Activity_Read:
// 		return Read{a}.isValid()
// 	}

// 	return nil
// }

// func getLabel(a *user.Activity) string {

// 	l := false

// 	switch a.ActivityType {
// 	case user.Activity_Sleep:
// 		l = Sleep{a}.isDone()
// 	case user.Activity_Play:
// 		l = Play{a}.isDone()
// 	case user.Activity_Eat:
// 		l = Eat{a}.isDone()
// 	case user.Activity_Read:
// 		l = Read{a}.isDone()
// 	}

// 	if l {
// 		return "COMPLETED"
// 	}
// 	return "PENDING"
// }

// func (p Play) isDone() bool {
// 	return diff(p.Timestamp, p.Duration)
// }
// func (s Sleep) isDone() bool {
// 	return diff(s.Timestamp, s.Duration)
// }
// func (e Eat) isDone() bool {
// 	return diff(e.Timestamp, e.Duration)
// }
// func (r Read) isDone() bool {
// 	return diff(r.Timestamp, r.Duration)
// }

// func diff(t, dura int64) bool {
// 	d := time.Now().Unix() - t

// 	if d*int64(time.Second) < dura {
// 		return false
// 	}

// 	return true
// }

// func (s Sleep) isValid() error {

// 	d := time.Duration(s.Duration)

// 	fmt.Println(d, 6*time.Hour)

// 	if d < time.Hour*6 || d > time.Hour*8 {
// 		return fmt.Errorf("sleep duration is invalid: %s", d)
// 	}

// 	return nil
// }

// func (p Play) isValid() error {

// 	d := time.Duration(p.Duration)

// 	if d < time.Minute*10 || d > time.Hour*2 {
// 		return fmt.Errorf("play duration is invalid: %s", d)
// 	}

// 	return nil
// }

// func (e Eat) isValid() error {

// 	d := time.Duration(e.Duration)

// 	if d < time.Minute*5 || d > time.Hour*1 {
// 		return fmt.Errorf("eat duration is invalid: %s", d)
// 	}

// 	return nil
// }

// func (r Read) isValid() error {

// 	d := time.Duration(r.Duration)

// 	if d < time.Minute*15 || d > time.Hour*2 {
// 		return fmt.Errorf("read duration is invalid: %s", d)
// 	}

// 	return nil
// }

func (s *UserActivityServer) AddUser(ctx context.Context, req *user.UserRequest) (*user.UserResponse, error) {
	users := req.GetUser()
	fmt.Println("testing")

	data := UserItem{
		Name:  users.GetName(),
		Email: users.GetEmail(),
		Phone: users.GetPhone(),
	}

	result, err := usersdb.InsertOne(mongoCtx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	oid := result.InsertedID.(primitive.ObjectID)
	_ = oid
	users.Userid = oid.Hex()

	return &user.UserResponse{User: users}, nil

}

func (s *UserActivityServer) AddActivity(ctx context.Context, req *user.ActivityRequest) (*user.ActivityResponse, error) {
	fmt.Println("testing")
	activity := req.GetActivity()
	fmt.Println("Email:")
	email := activity.Email
	fmt.Println(email)
	check := usersdb.FindOne(ctx, bson.M{"email": email})
	fmt.Println("checking COMPLETED")
	data := UserItem{}
	if err := check.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("could not find user with the email %s: %v", req.Activity.GetEmail(), err))
	}
	activity.Timestamp = time.Now().Unix()
	fmt.Println(activity.Timestamp)
	// data1 := ActivityItem{
	// 	Email:                 activity.GetEmail(),
	// 	Activity_ActivityType: activity.GetActivityType(),
	// 	Timestamp:             activity.GetTimestamp(),
	// 	Duration:              activity.GetDuration(),
	// }

	result, err := activitydb.InsertOne(mongoCtx, activity)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	oid := result.InsertedID.(primitive.ObjectID)
	_ = oid
	activity.Id = oid.Hex()

	return &user.ActivityResponse{Activity: activity}, nil

}

func (s *UserActivityServer) QueryUser(ctx context.Context, req *user.QueryUserRequest) (*user.QueryUserResponse, error) {
	oid, err := primitive.ObjectIDFromHex(req.GetUserid())

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	result := usersdb.FindOne(ctx, bson.M{"_id": oid})

	data := UserItem{}

	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("could not find user with Object Id %s: %v", req.GetUserid(), err))
	}

	response := &user.QueryUserResponse{
		User: &user.User{
			Userid: oid.Hex(),
			Name:   data.Name,
			Email:  data.Email,
			Phone:  data.Phone,
		},
	}
	return response, nil

}

// func (s *UserActivityServer) QueryActivity(ctx context.Context, req *user.QueryActivityRequest) (*user.QueryActivityResponse, error) {
// 	oid, err := primitive.ObjectIDFromHex(req.GetActivityid())
// 	if err != nil {
// 		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
// 	}
// 	result := activitydb.FindOne(ctx, bson.M{"_id": oid})
// 	// result.Label = getLabel(result)

// 	data := ActivityItem{}

// 	if err := result.Decode(&data); err != nil {
// 		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("could not find user with Object Id %s: %v", req.GetActivityid(), err))
// 	}

// 	response := &user.QueryActivityResponse{
// 		Activity: &user.Activity{
// 			Id: oid.Hex(),
// 		},
// 	}
// 	return response, nil

// }

// func (s *UserActivityServer) AddActivity(ctx context.Context, req *user.ActivityRequest) (*user.ActivityResponse, error) {

// 	activity := req.GetActivity()

// 	if err := Validate(activity); err != nil {
// 		return nil, err
// 	}
// 	activity.Timestamp = time.Now().Unix()

// 	result, err := activitydb.InsertOne(mongoCtx, activity)
// 	if err != nil {
// 		return nil, status.Errorf(
// 			codes.Internal,
// 			fmt.Sprintf("Internal error: %v", err),
// 		)
// 	}

// 	oid := result.InsertedID.(primitive.ObjectID)
// 	_ = oid
// 	activity.Id = oid.Hex()
// 	activity.Label = getLabel(activity)

// 	return &user.ActivityResponse{Activity: activity}, nil

// }

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	srv := &UserActivityServer{}

	user.RegisterUserActivityServer(s, srv)
	mongoCtx = context.Background()
	db, err = mongo.Connect(mongoCtx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping(mongoCtx, nil)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v\n", err)
	} else {
		fmt.Println("Connected to MongoDB")
	}
	usersdb = db.Database("useractivity").Collection("user")
	activitydb = db.Database("useractivity").Collection("activity")
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
