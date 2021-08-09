package dashboardserver

import (
	"context"

	"google.golang.org/grpc"
)

func Unary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return nil, nil
}
