package dashboardserver

import (
	"context"

	"github.com/Cheep2Workshop/proj-web/controller"
	pb "github.com/Cheep2Workshop/proj-web/grpc/dashboard"
	"github.com/Cheep2Workshop/proj-web/models"
	"github.com/Cheep2Workshop/proj-web/models/repo"
	"google.golang.org/grpc"
)

type AuthServer struct {
	pb.AuthServiceServer
	DBClient *repo.DbClient
	JwtMgr   *controller.JWTManager
}

func runAuthServer(s *grpc.Server, jwtMgr *controller.JWTManager) {
	authServ := &AuthServer{
		DBClient: repo.Client,
		JwtMgr:   jwtMgr,
	}
	pb.RegisterAuthServiceServer(s, authServ)
}

// Signup will save the user if email not duplicate
func (s *AuthServer) Signup(ctx context.Context, in *pb.SignupReq) (*pb.SignupRes, error) {
	user := models.User{
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
		Admin:    in.Admin,
	}
	result, err := repo.Client.Signup(user)
	return &pb.SignupRes{Result: &pb.Result{Status: result}}, err
}

// Login will check the email, password existed and save login log
func (s *AuthServer) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginRes, error) {
	req := repo.LoginReq{
		Email:    in.Email,
		Password: in.Password,
	}
	// step1: login
	user, err := repo.Client.Login(req)
	if err != nil {
		return &pb.LoginRes{Result: &pb.Result{Status: false}}, err
	}
	// step2: generate jwt
	token, err := s.JwtMgr.GenerateJwt(*user)
	if err != nil {
		return &pb.LoginRes{Result: &pb.Result{Status: false}}, err
	}
	// step3: save login log
	err = repo.Client.SaveLoginLog(in.Email)
	if err != nil {
		return &pb.LoginRes{Result: &pb.Result{Status: false}}, err
	}
	return &pb.LoginRes{Result: &pb.Result{Status: true}, Token: token}, err
}

// Logout
func (s *AuthServer) Logout(ctx context.Context, in *pb.LogoutReq) (*pb.LogoutRes, error) {
	return &pb.LogoutRes{Result: SuccessResult}, nil
}
