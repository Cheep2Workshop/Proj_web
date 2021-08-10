package dashboardclient

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/Cheep2Workshop/proj-web/grpc/dashboard"
	"google.golang.org/grpc"
)

type AuthClient struct {
	Serv     pb.AuthServiceClient
	Email    string
	Password string
}

func NewAuthClient(c *grpc.ClientConn, email string, password string) AuthClient {
	service := pb.NewAuthServiceClient(c)
	return AuthClient{service, email, password}
}

func (client *AuthClient) Signup(name string, email string, password string, admin bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := client.Serv.Signup(
		ctx,
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
		client.Email = email
		client.Password = password
		return
	}
	fmt.Println("Signup failed.")
}

func (client *AuthClient) Login() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := client.Serv.Login(
		ctx,
		&pb.LoginReq{
			Email:    client.Email,
			Password: client.Password})
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
