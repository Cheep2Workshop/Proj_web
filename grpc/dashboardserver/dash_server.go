package dashboardserver

import (
	"context"
	"log"
	"net"

	"github.com/Cheep2Workshop/proj-web/controller"
	pb "github.com/Cheep2Workshop/proj-web/grpc/dashboard"
	"github.com/Cheep2Workshop/proj-web/models"
	"github.com/Cheep2Workshop/proj-web/models/repo"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	port = ":50045"
)

var SuccessResult = &pb.Result{Status: true}

type DashServer struct {
	pb.DashboardServiceServer
	DBClient *repo.DbClient
	JwtMgr   *controller.JWTManager
}

func Run() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create jwt manager
	jwtMgr := &controller.JWTManager{
		Issuer: controller.Issuer,
		Secret: controller.Secret,
	}
	// create interceptor
	authInterceptor := AuthInterceptor{
		JwtMgr: jwtMgr,
		AccessibleMethods: map[string]bool{
			"SetUser":      true,
			"DeleteUser":   true,
			"GetLoginLogs": true,
			"GetUser":      true,
		},
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor.Unary()),
	)
	// run service
	runAuthServer(s, jwtMgr)
	runDashServer(s, jwtMgr)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func runDashServer(s *grpc.Server, jwtMgr *controller.JWTManager) {

	dashserv := &DashServer{
		DBClient: repo.Client,
		JwtMgr:   jwtMgr,
	}

	pb.RegisterDashboardServiceServer(s, dashserv)

}

// Signup will save the user if email not duplicate
func (s *DashServer) Signup(ctx context.Context, in *pb.SignupReq) (*pb.SignupRes, error) {
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
func (s *DashServer) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginRes, error) {
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
func (s *DashServer) Logout(ctx context.Context, in *pb.LogoutReq) (*pb.LogoutRes, error) {
	return &pb.LogoutRes{Result: SuccessResult}, nil
}

// Change the user data by email
func (s *DashServer) SetUser(ctx context.Context, in *pb.SetUserReq) (*pb.SetUserRes, error) {
	var err error
	// handle req
	req := repo.SetUserReq{
		Email:           in.Email,
		TargetEmail:     in.TargetEmail,
		ChangedName:     in.ChangedName,
		ChangedPassword: in.ChangedPassword,
	}
	err = repo.Client.SetUser(req)
	if err != nil {
		return &pb.SetUserRes{Result: &pb.Result{Status: false}}, err
	}
	return &pb.SetUserRes{Result: SuccessResult}, nil
}

// Delete user by email
func (s *DashServer) DeleteUser(ctx context.Context, in *pb.DeleteUserReq) (*pb.DeleteUserRes, error) {
	var err error
	// handle request
	req := repo.DeleteUserReq{
		Email:       in.Email,
		DeleteEmail: in.DeleteEmail,
	}
	err = repo.Client.DeleteUser(req)
	if err != nil {
		return &pb.DeleteUserRes{Result: &pb.Result{Status: false}}, err
	}
	return &pb.DeleteUserRes{Result: SuccessResult}, nil
}

// Get the login logs by email
func (s *DashServer) GetLoginLogs(ctx context.Context, in *pb.LoginLogReq) (*pb.LoginLogRes, error) {
	var err error
	// get logs
	results, err := repo.Client.GetLoginLogs(in.UserEmail)
	if err != nil {
		return &pb.LoginLogRes{Result: &pb.Result{Status: false}}, err
	}
	// response
	res := &pb.LoginLogRes{Result: SuccessResult}
	res.Logs = make([]*pb.LoginLog, len(results))
	for i, result := range results {
		res.Logs[i] = &pb.LoginLog{
			LogId: int32(result.ID),
			User: &pb.User{
				Name:     result.Name,
				Email:    result.Email,
				Admin:    result.Admin,
				CreateAt: timestamppb.New(result.CreatedAt),
			},
		}
	}

	return res, nil
}

func (s *DashServer) GetUser(ctx context.Context, in *pb.GetUserReq) (*pb.GetUserRes, error) {
	result, err := repo.Client.GetUserByEmail(in.Email)
	if err != nil {
		return &pb.GetUserRes{Result: &pb.Result{Status: false, Msg: err.Error()}}, nil
	}
	return &pb.GetUserRes{
		Result: SuccessResult,
		User: &pb.User{
			Name:     result.Name,
			Email:    result.Email,
			Admin:    result.Admin,
			CreateAt: timestamppb.New(result.CreatedAt),
		}}, nil
}
