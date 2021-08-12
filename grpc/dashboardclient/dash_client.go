package dashboardclient

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/Cheep2Workshop/proj-web/grpc/dashboard"
	"github.com/Cheep2Workshop/proj-web/models/repo"
	"google.golang.org/grpc"
)

const (
	addr = "localhost:50045"
)

type DashClient struct {
	Serv pb.DashboardServiceClient
}

var token string

func runClient() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	authClient := NewAuthClient(conn, "jack81722@gmail.com", "123456")
	authClient.Login()
}

func (client *DashClient) SetUser(req repo.SetUserReq) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.Serv.SetUser(
		ctx,
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

func (client *DashClient) DeleteUser(req repo.DeleteUserReq) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.Serv.DeleteUser(
		ctx,
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

func (client *DashClient) GetUser(email string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.Serv.GetUser(
		ctx,
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
