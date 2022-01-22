package userservice

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/endpoint"
	kitGRPC "github.com/go-kit/kit/transport/grpc"
	"github.com/jcromanu/final_project/pb"
)

var (
	errBadTypeError = errors.New("Bad type error ")
)

type userServiceServer struct {
	createUser *kitGRPC.Server
	pb.UnimplementedUserServiceServer
}

func makeCreateUserGRPCServer(ep endpoint.Endpoint, opts []kitGRPC.ServerOption, logger log.Logger) *kitGRPC.Server {
	return kitGRPC.NewServer(
		ep,
		makeDecodeGRPCCreateUserRequest(logger),
		makeEncodeGRPCCReateUserResonse(logger),
		opts...,
	)
}

func NewGRPCServer(ep Endpoints, opts []kitGRPC.ServerOption, log log.Logger) pb.UserServiceServer {
	return &userServiceServer{
		createUser: makeCreateUserGRPCServer(ep.CreateUser, opts, log),
	}
}

func (srv *userServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	_, resp, err := srv.createUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r, ok := resp.(*pb.CreateUserResponse)
	if !ok {
		return nil, errBadTypeError
	}
	return r, nil
}

func (srv *userServiceServer) mustEmbedUnimplementedUserServiceServer() {

}