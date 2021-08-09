package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/Cheep2Workshop/proj-web/grpc/dashboard"
	"github.com/Cheep2Workshop/proj-web/orm"
	"google.golang.org/grpc"
)

const (
	addr = "localhost:50045"
)

type Stub struct {
	Client  pb.DashboardServiceClient
	Context context.Context
}

var token string

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDashboardServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stub := &Stub{c, ctx}
	stub.Signup("Trash", "Trash@gmail.com", "654321", false)
	//stub.Login("Trash@gmail.com", "654321")
	//stub.ChangeUserData("Trash@gmail.com", "", "")
	// stub.ChangeUserData("Trash@gmail.com", "Trash", "")
	// stub.ChangeUserData("Trash@gmail.com", "", "999888")
}

func (stub *Stub) Signup(name string, email string, password string, admin bool) {
	res, err := stub.Client.Signup(
		stub.Context,
		&pb.SignupReq{
			Name:     name,
			Email:    email,
			Password: password,
			Admin:    admin})
	if err != nil {
		log.Fatalf("Connect failed: %v", err)
	}
	if res.Result.Status {
		fmt.Println("Signup succeed.")
		return
	}
	fmt.Println("Signup failed.")
}

func (stub *Stub) Login(email string, password string) {
	res, err := stub.Client.Login(stub.Context, &pb.LoginReq{Email: email, Password: password})
	if err != nil {
		log.Fatalf("Connect failed: %v", err)
	}
	if res.Result.Status {
		// cache token
		token = res.Token
		return
	}
	fmt.Println("Login failed.")
}

func (stub *Stub) SetUser(req orm.SetUserReq) {
	res, err := stub.Client.SetUser(
		stub.Context,
		&pb.SetUserReq{
			Token:           token,
			Email:           req.Email,
			TargetEmail:     req.TargetEmail,
			ChangedName:     req.ChangedName,
			ChangedPassword: req.ChangedPassword,
		})
	if err != nil {
		log.Fatalf("Connect failed: %v", err)
	}
	if res.Result.Status {
		fmt.Println("Change user data succeed.")
		return
	}
	fmt.Printf("Change user data failed: %s", res.Result.Msg)
}

func (stub *Stub) DeleteUser(req orm.DeleteUserReq) {
	res, err := stub.Client.DeleteUser(
		stub.Context,
		&pb.DeleteUserReq{
			Token:       token,
			Email:       req.Email,
			DeleteEmail: req.DeleteEmail,
		},
	)
	if err != nil {
		log.Fatalf("Connect failed: %v", err)
	}
	if res.Result.Status {
		fmt.Println("Delete user succeed.")
		return
	}
	fmt.Printf("Delete user failed: %s", res.Result.Msg)
}

func (stub *Stub) GetUser(email string) {
	res, err := stub.Client.GetUser(
		stub.Context,
		&pb.GetUserReq{
			Email: email,
		},
	)
	if err != nil {
		log.Fatalf("Connect failed: %v", err)
	}
	if res.Result.Status {
		fmt.Println("Get user succeed.")
		return
	}
	fmt.Printf("Get user failed: %s", res.Result.Msg)
}
