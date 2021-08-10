package dashboardserver

import (
	"context"
	"log"

	"github.com/Cheep2Workshop/proj-web/controller"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthInterceptor struct {
	JwtMgr            *controller.JWTManager
	AccessibleMethods map[string]bool
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {
		log.Println("Auth interceptor")
		err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) error {
	// check the method accessible
	_, ok := interceptor.AccessibleMethods[method]
	if !ok {
		return nil
	}
	// check metadata existed
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided.")
	}
	// check authorization in metadata
	values := md["authorization"]
	if len(values) <= 0 {
		return status.Errorf(codes.Unauthenticated, "authorization is not provided.")
	}
	// verify jwt
	token := values[0]
	_, err := interceptor.JwtMgr.VerifyJwt(token)
	if err != nil {
		return err
	}
	return nil
}
